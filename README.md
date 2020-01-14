1、项目说明

本文将介绍如何实现一个基于websocket系统。

使用golang实现websocket通讯，使用gin框架、nginx负载、可以水平部署、程序内部相互通讯。 后期准备支持grpc通讯协议

2、webSocket

2.1 webSocket的兼容性

![服务端处理一个请求]

大多数场景我们需要主动通知用户，如:聊天系统、用户完成任务主动告诉用户、一些运营活动需要通知到在线的用户
可以获取用户在线状态
在没有长连接的时候通过客户端主动轮询获取数据
可以通过一种方式实现，多种不同平台(H5/Android/IOS)去使用
2.2 webSocket建立过程

2.1.1

客户端先发起升级协议的请求
客户端发起升级协议的请求，采用标准的HTTP报文格式，在报文中添加头部信息 Connection: Upgrade表明连接需要升级 Upgrade: websocket需要升级到 websocket协议 Sec-WebSocket-Key: xxxxx 这个是base64 encode 的值，是浏览器随机生成的，与服务器响应的 Sec-WebSocket-Accept对应

升级协议完成以后，客户端和服务器就可以相互发送数据
2.1.2 启动端口监听（如何实现基于webSocket的长连接系统）

websocket需要监听端口，所以需要在golang 成功的 main 函数中用协程的方式去启动程序
main.go 实现启动
go route.WsRoute()
// 启动程序（使用go实现webSocket服务端）

func WsRoute(){
	http.HandleFunc("/ws", func(w http.ResponseWriter, req *http.Request) {
		con, err :=  upgrader.Upgrade(w,req,nil)
		defer func() {
			con.Close()
		}()
		if err != nil{
			log.Println("升级为websocket失败",err.Error())
			return
		}
		ser := wsrobot.NewSessionM()
		go ser.Start()
		client :=  wsrobot.NewSession()
		client.Addr = con.RemoteAddr().String()
		client.Conn = con
		client.TimeS = time.Now().Unix()
		ser.Register <- client
		ser.SetSession(client.Addr,client)
		client.Sessions = ser
		defer func() {
			ser.Unregister <- client
		}()
		//心跳检测
		go ser.HeartBeat(20)
		//session := ser.SessionMaster.GetSessionById(client.Addr)
		//if session == nil{
		//	return
		//}
		go client.Write()
		client.Read()

	})
	http.ListenAndServe(config.WebSocketPort, nil)
}
2.1.3 升级协议

客户端是通过http请求发送到服务端，我们需要对http协议进行升级为websocket协议
对http请求协议进行升级 golang 库[gorilla/websocket]已经做得很好了，我们直接使用就可以了
在实际使用的时候，建议每个连接使用两个协程处理客户端请求数据和向客户端发送数据，虽然开启协程会占用一些内存，但是读取分离，减少收发数据堵塞的可能
var upgrader = websocket.Upgrader{
 	ReadBufferSize:1024,
 	WriteBufferSize:1024,
 	CheckOrigin: func(r *http.Request) bool {
 		return true
 	},
 }
2.1.3 客户端连接的管理

当前程序有多少用户连接，还需要对用户广播的需要，这里我们就需要一个管理者(SessionM结构体)，处理这些事件:
记录全部的连接、登录用户的可以通过 client.Addr查到用户连接
使用map存储，就涉及到多协程并发读写的问题，所以使用sync.map
定义三个channel ，分别处理客户端建立连接、断开连接、全员广播事件
// 连接管理
type SessionM struct {
	sessions sync.Map
	Users sync.Map
	Register    chan *Session       // 连接连接处理
	Unregister  chan *Session   //断开连接
	Broadcast   chan *Msg   //全员广播
}

// 初始化
func NewSessionM()*SessionM{
	return &SessionM{
		Register:make(chan *Session,1000),
		Unregister:make(chan *Session,1000),
		Broadcast:make(chan *Msg,1000),
	}
}

2.1.4 注册客户端的socket的写的异步处理程序

防止发生程序崩溃，所以需要捕获异常
为了显示异常崩溃位置这里使用string(debug.Stack())打印调用堆栈信息
如果写入数据失败了，可能连接有问题，就关闭连接
client.go
// 初始化  定义一个channel  写队列inchan
type Session struct {
	Addr          string          // 客户端地址
	Conn *websocket.Conn
	inchan chan *Msg   //读队列
	PlatId         int32   // 登录的平台Id app/web/ios
	UserId        string          // 用户Id，用户登录以后才有
	LoginTime     int64          // 登录时间 登录以后才有
	Sessions    *SessionM
	TimeS int64    //记录用户连接时间 连接超时 关闭连接
}
- 添加AppId,PlatId，设计的时候为了做成通用性，设计AppId用来表示用户在哪个平台登录的(web、app、ios等)，方便后续扩展
// 向客户端写数据
func(this *Session)Write(){
	ticker := time.NewTicker(config.PingPeriod)
	defer func() {
		if r:=recover();r != nil{
			fmt.Println("停止写",string(debug.Stack()),r)
		}
	}()

	defer func() {
		this.Sessions.Unregister <- this
	}()

	defer func() {
		ticker.Stop()
	}()
	for{
		select {
		case msg,ok :=<- this.inchan:
			if !ok{
				fmt.Println("发送数据错误",this.Addr,ok)
				return
			}
			fmt.Println(string(msg.data),"---------------------")
			this.Conn.WriteMessage(msg.messType,msg.data)
		case <-ticker.C:
			// 出现超时情况
			this.Conn.SetWriteDeadline(time.Now().Add(config.WriteWait))
			if err := this.Conn.WriteMessage(websocket.PingMessage,nil);err != nil{
				return
			}
		}
	}



}

2.1.5 注册客户端的socket的读的异步处理程序

循环读取客户端发送的数据并处理
如果读取数据失败了，关闭channel
client.go
// 读取客户端数据
func(this *Session)Read(){
	defer func() {
		if r := recover();r != nil{
			fmt.Println("停止写",string(debug.Stack()),r)
		}
	}()

	defer func() {
		fmt.Println("关闭inchan管道",this)
	}()


	//this.Conn.SetReadLimit(config.MaxMessageSize)
	//this.Conn.SetReadDeadline(time.Now().Add(config.PongWait))
	for{
		messType,data,err :=this.Conn.ReadMessage()
		if err != nil{
			websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			log.Println("消息读取出现错误",this.Addr,err.Error())
			this.Sessions.Unregister <- this
			//this.Sessions.DelSessionById(this.Addr)
			return
		}
		fmt.Println("读取客户端数据：",string(data))
		msg := &Msg{
			messType,
			data,
		}
		this.dealMsg(msg)
	}
}

2.1.6 接收客户端数据

约定发送和接收请求数据格式，为了js处理方便，采用了json的数据格式发送和接收数据

登录发送数据示例:

      {
		    "Module": "robot",
            "Action": "Login",
            "Data":{
                "PlatId":"1",
                "UserId":"abc123456"
            },

        }
1.心跳检测:robot/app/route 用来发送登录请求和连接保活(长时间没有数据发送的长连接容易被浏览器、移动中间商、nginx、服务端程序断开) 2.防止内存溢出和Goroutine不回收 定时任务清除超时连接 go ser.HeartBeat(20) //遍历所有的sess 上次接受消息的时间 如果超过num 就删除sess
func(this *SessionM)HeartBeat(num int64){
		for{
			time.Sleep(time.Second)
			this.sessions.Range(func(key, val interface{}) bool {
				sess,ok := val.(*Session)
				if !ok{
					return true
				}


				if time.Now().Unix() - sess.TimeS > num{
					//删除session
					fmt.Println("心跳删除用户")
					this.DelSessionById(key.(string))
				}
				return true
			})
		}
}


约定的请求数据格式
/************************  请求数据  **************************/
// 通用请求数据格式
{
		    "Module": "",  //调用的项目模块
            "Action": "",  //调用方法
            "Data":{

            },  //请求的数据
}
// 登录请求数据
{
		    "Module": "robot",
            "Action": "Login",
            "Data":{
                "PlatId":"1",
                "UserId":"abc123456"
  },

// 心跳请求数据
{
            "Module": "robot",
            "Action": "HeartBeat"
};
/************************ 相应数据 **************************/

2.1.7 使用路由的方式处理客户端的请求数据

使用路由的方式处理由客户端发送过来的请求数据
以后添加请求类型以后就可以用类是用http相类似的方式(router-controller)去处理
1.注册websocket路由
main.go	启动注册
    route.WsInit()
开始注册
robot/app/inits register.go
func WebSocketInit(){
   router.RegisterRobotStructFun("robot",&Robot{})
   //
}
3.注册过程 -** inits/robotRegister.go注册路由
3.1 申明结构体
    type RoutersMap struct {
        Controller  map[string]map[string]func( client interface{}, data map[string]interface{})bool
    }
3.2 注册结构体对应的方法 通过reflect获取结构所有的方法 并存入map
   func(this *RoutersMap) RegisterRobotStructFun(actionName string,mod interfaces.RobotFace)bool{
       this.Controller[actionName] = make(map[string]func( interface{}, map[string]interface{})bool)
       temval := reflect.ValueOf(mod)
       temType := reflect.TypeOf(mod)
       for i:=0 ;i<temType.NumMethod();i++{
           tem := temval.Method(i).Interface()
           if temFunc,ok:= tem.(func( client interface{},data map[string]interface{})bool);ok{
               this.Controller[actionName][temType.Method(i).Name] = temFunc
           }
       }
       //this.MapToJsonDemo1()
       return true
   }

3.2 Hook 结构体对应的方法
   func(this *RoutersMap)HookRobotAction(actionName,funcName string,client interface{},data map[string]interface{})bool{
       if _,exit := this.Controller[actionName];!exit{
           return false
       }
       if action,exit := this.Controller[actionName][funcName];exit{
           action(client,data)
       }
       return true
   }


   // 请求参数如下 ：此方法可以获取到robot路由对应的结构体中的Action方法
   {
               "Module": "robot",
               "Action": "HeartBeat"  //对应HookRobotAction方法 funcName参数 
   };
   
2.1.8 Goroutine管理 防止内存溢出

1.定时清理超时连接

读写的Goroutine有一个失败，则相互关闭 write()Goroutine写入数据失败，关闭c.Socket.Close()连接，会关闭read()Goroutine read()Goroutine读取数据失败，关闭close(c.Send)连接，会关闭write()Goroutine
2.1.9 Goroutine管理 防止内存溢出

关闭读写的Goroutine //关闭连接并删除session func(this *SessionM)DelSessionById(addr string){ client,ok :=this.sessions.Load(addr) if ok{ if client,oks := client.(*Session);oks{ client.close() } } fmt.Println("删除连接用户",addr) this.sessions.Delete(addr) return }

2.1.10 使用路由的方式处理客户端的请求数据

使用 pprof 分析性能、耗时

3 webSocket客户端

3.1 使用javaScript实现webSocket客户端

3.1.1 启动并注册监听程序

js 建立连接，并处理连接成功、收到数据、断开连接的事件处理
ws = new WebSocket("ws://127.0.0.1:8089/ws");

3.1.2 发送数据

需要注意:连接建立成功以后才可以发送数据
建立连接以后由客户端向服务器发送数据示例
登录:
{
		    "Module": "robot",
            "Action": "Login",
            "Data":{
                "PlatId":"1",
                "UserId":"abc123456"
            },

}

心跳:
{
            "Module": "robot",
            "Action": "HeartBeat"
};

ping 查看服务是否正常:
{
            "Module": "robot",
            "Action": "Ping"
 };

关闭连接:
ws.close();

3.1.3 图片和语言消息

发送图片消息，发送消息者的客户端需要先把图片上传到文件服务器，上传成功以后获得图片访问的 URL，然后由发送消息者的客户端需要将图片 URL 发送到 websocket，websocket 图片的消息格式发送给目标客户端，消息接收者客户端接收到图片的 URL 就可以显示图片消息。

图片消息的结构:

{
  "type": "img",
  "from": "杰杰棒",
  "url": "http://xxxx/images/home.png",
}
语言消息、和视频消息和图片消息类似，都是先把文件上传服务器，然后通过 gowebsocket 传递文件的 URL。

4、WebSocket 项目

4.1 项目说明

实现群聊,单聊的功能
支持水平部署，部署的机器之间可以相互通讯
项目架构图（未）
4.2 项目依赖

本项目使用go mod管理依赖
# 主要使用到的包
github.com/gin-gonic/gin
github.com/gorilla/websocket
google.golang.org/grpc
github.com/golang/protobuf
5、webSocket项目Nginx配置

5.1 配置Nginx

使用nginx实现内外网分离，对外只暴露Nginx的Ip(一般的互联网企业会在nginx之前加一层LVS做负载均衡)，减少入侵的可能
使用Nginx可以利用Nginx的负载功能，前端再使用的时候只需要连接固定的域名，通过Nginx将流量分发了到不同的机器
同时我们也可以使用Nginx的不同的负载策略(轮询、weight、ip_hash)
5.2 问题处理

运行nginx测试命令
6、压测

6.1 Linux内核优化

设置文件打开句柄数
6.2压测数据（未）

7 项目说明

7.1 说明

参考本项目源码 http://192.168.0.222:5555/svn/magook/trunk/server/robot_go

为了方便，http系统和webSocket系统合并在一个系统中

http系统: 获取全部在线的用户，群组,以及聊天记录，加群，删除群好友，用户登录注册，获取登录token 考虑主要是两点: 1.服务分离，让系统尽量的简单一点，不掺杂其它业务逻辑 2.除了发消息是走webSocket，其他操作全部使用http连接，使收和发送数据分离的方式，可以加快收发数据的效率

7.2 架构

项目启动注册和用户连接时序图(未)
8 已经实现的功能

gin log日志(请求日志+debug日志)
读取配置文件 完成
定时脚本，清理过期未心跳连接 完成
http接口，获取登录、连接数量 完成
http接口，发送push、查询有多少人在线 完成
grpc 程序内部通讯，发送消息 未完成
appIds 一个用户在多个平台登录
界面，把所有在线的人拉倒一个群里面，发送消息 未完成
单聊、群聊 完成
实现分布式，水平扩张 完成
压测脚本 未完成
文档整理 未完成
文档目录 未完成
架构图以及扩展 未完成
有人加入以后广播全体 未完成
引入机器人 待定
html接收到消息 显示到界面 未完成
8.1 需要完善、优化
