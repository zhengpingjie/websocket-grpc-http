package wsrobot

import (
	"fmt"
	"robot/app/controller/grpcrobot"
	"runtime/debug"
	"sync"
	"time"
)


type SessionM struct {
	Sessions sync.Map
	Users sync.Map
	sLock sync.RWMutex       // 读写锁
	Register    chan *Session       // 连接连接处理
	Unregister  chan *Session
	Binds   sync.Map
//	Broadcast   chan *Msg
}

type Msg struct {
	messType int
	data []byte
}


func NewSessionM()*SessionM{
	return &SessionM{
		Register:make(chan *Session,100),
		Unregister:make(chan *Session),
		//Broadcast:make(chan *Msg,1000),
	}
}


func(this *SessionM)SetBinds(clientId,toclientId string){
	this.Binds.Store(clientId,toclientId)
}

func(this  *SessionM)GetBinds(clientId string)string{
	toclientId,ok := this.Binds.Load(clientId)
	if ok{
		return toclientId.(string)
	}
	return ""
}

func(this *SessionM)SetUser(userId ,Addr string){
	this.Users.Store(userId,Addr)
}

func(this *SessionM)GetUser(userId string) string{
	Addr,ok := this.Users.Load(userId)
	if ok{
			return Addr.(string)
	}
	return ""
}
func(this *SessionM)DelUser(userId string){
	this.Users.Delete(userId)
}

func(this *SessionM)SetSession(addr string,client *Session){
	this.Sessions.Store(addr,client)
}

//关闭连接并删除session
func(this  *SessionM)DelSessionById(addr string){
	client,ok :=this.Sessions.Load(addr)
	if ok{
		if client,oks := client.(*Session);oks{
			fmt.Println("连接clientid是", len(client.ClientId))
			if len(client.ClientId)>0{
				toclientId:= this.GetBinds(client.ClientId)
				if toclientId != ""{
					msg := make(map[string]interface{})
					msg["Module"] = "robot"
					msg["Action"] = "Disconnect"
					data := grpcrobot.NewRobotData()
					data.ClientId = toclientId
					data.PlatId = client.PlatId
					data.FromId = client.ClientId
					msg["Data"] = data
					if addr:= this.GetUser(toclientId);addr != ""{
						this.SendMsgById(addr,msg)
					}
					this.Binds.Delete(client.ClientId)
				}
			}
			client.close()
		}
	}
	fmt.Println("删除连接用户",addr)
	this.Sessions.Delete(addr)
	return
}

func(this *SessionM)GetSessionById(addr string)*Session{
	this.sLock.RLock()
	defer this.sLock.RUnlock()
	client,ok :=this.Sessions.Load(addr)
	if ok{
		if client,oks := client.(*Session);oks{

			return client
		}
	}
	return nil
}


//发送消息
func(this *SessionM)SendMsgById(addr string,data map[string]interface{})bool{
	defer func() {
		if r := recover();r != nil{
			fmt.Println("停止写",string(debug.Stack()),r)
		}
	}()

	client := this.GetSessionById(addr)
	if client == nil{
		return false
	}
	if !client.isclose{
		client.inchan = make(chan interface{},100)
	}
	client.inchan <- data
	//client.Conn.WriteJSON(data)
	return true
}

//grpc发送消息
func(this  *SessionM)GrpcSendMsgById(addr string,data interface{})bool{
	//defer func(fromId string) {
	//	this.Unregister <- this.GetSessionById(fromId)
	//}(fromId)
	client := this.GetSessionById(addr)
	if client == nil{
		return false
	}
	client.inchan <- data
	//client.Conn.WriteJSON(data)
	return true
}

//心跳检测 遍历所有的sess 上次接受消息的时间 如果超过num 就删除sess
func(this *SessionM)HeartBeat(num int64){
		//for{
		//	time.Sleep(time.Second)
			this.Sessions.Range(func(key, val interface{}) bool {
				sess,ok := val.(*Session)
				if !ok{
					return true
				}
				if time.Now().Unix() - sess.TimeS > num{
					//删除session
					fmt.Println("心跳删除用户")
					//if this.GetSessionById(sess.Addr) !=nil{
					this.Unregister <- sess
				//	}
				}
				return true
			})
		//}
}





func(this  *SessionM)Start(){
	for{
		select {
			case conn := <- this.Register:
				//建立连接事件
				this.RegisterEvent(conn)
			case conn := <- this.Unregister:
				//断开连接
				this.UnregisterEvent(conn)
			default:


		}
	}
}


//注册
func(this  *SessionM)RegisterEvent(client *Session){
	this.SetSession(client.Addr,client)
	fmt.Println("用户已经建立连接",client.Addr)
}

//断开事件
func(this  *SessionM)UnregisterEvent(conn *Session){
	fmt.Println("删除连接用户",conn.Addr)
	this.DelSessionById(conn.Addr)
}
