package wsrobot

import (
	"fmt"
	"github.com/gorilla/websocket"
	"runtime/debug"
	"time"
)


type Session struct {
	Addr          string          // 客户端地址
	Conn *websocket.Conn
	inchan chan interface{}   //读队列
	isclose bool   //管道是否关闭
	PlatId         int32   // 登录的平台Id app/web/ios
	ClientId        string          // 用户Id，用户登录以后才有
	LoginTime     int64          // 登录时间 登录以后才有
	Sessions    *SessionM
	TimeS int64

}


func NewSession()*Session{
	return &Session{
		inchan: make(chan interface{},100),
		isclose:true,
	}
}


func(this *Session)Read(){
	defer func() {
		if r := recover();r != nil{
			fmt.Println("停止写",string(debug.Stack()),r)
		}
	}()

	defer func() {
		fmt.Println("读取客户端数据 关闭inchan")
		close(this.inchan)
		this.isclose = false
	}()


	//this.Conn.SetReadLimit(config.MaxMessageSize)
	//this.Conn.SetReadDeadline(time.Now().Add(config.PongWait))
	for{
		messType,data,err :=this.Conn.ReadMessage()
		if err != nil{
			//websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure)
			//log.Println("消息读取出现错误",this.Addr,err.Error())
			//this.Sessions.Unregister <- this
			//this.Sessions.DelSessionById(this.Addr)
			//fmt.Println("读取客户端数据 错误", this.Addr, err)
			return
		}
		//fmt.Println("读取客户端数据：",string(data))
		msg := &Msg{
			messType,
			data,
		}
		//fmt.Println("内存地址。。。。。。。。",&Ser)
		this.dealMsg(msg)
	}
}


func (this *Session) Write() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("write stop", string(debug.Stack()), r)
		}
	}()

	defer func() {
		fmt.Printf("使用'%%+v' %+v\n", this)
		fmt.Printf("使用'%%#v' %#v\n", this)
		fmt.Printf("使用'%%T' %T\n", this)
		if this.Sessions != nil{
			this.Sessions.Unregister <- this
		}
	}()

	for {
		select {
		case message,ok:= <-this.inchan:
			if !ok{
				// 发送数据错误 关闭连接
				fmt.Println("Client发送数据 关闭连接", this.Addr, "ok", ok)
				return
			}
			fmt.Println("通过管道接写数据",message)
			this.Conn.WriteJSON(message)

		}
	}
}


func(this *Session)dealMsg(msg *Msg){
	fmt.Println("数据处理阶段dealMsg",string(msg.data),this.Addr)
	defer func() {
		if r := recover();r !=nil{
			fmt.Println("处理数据停止",r)
		}
	}()
	this.Sessions = Ser
	//fmt.Println("内存地址client。。。。。。。。",&this.Sessions)
	//this.Sessions.Sessions.Range(func(key, value interface{}) bool {
	//	fmt.Println(key,"连接用户。。。。。。。。。。。。。。。")
	//	return true
	//})
	router.DealParam(msg.data,this)
}

// 关闭客户端连接
func (this *Session) close() {
	this.Conn.Close()
}

//更新心跳时间
func(this *Session)UpdTime(){
	this.TimeS = time.Now().Unix()
}

//判断用户是否登录
func(this  *Session)isLogin()bool{
	 if this.ClientId !=""{
	 	return true
	 }
	 return false
}