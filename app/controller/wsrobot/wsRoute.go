package wsrobot

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"robot/app/config"
	"time"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:1024,
	WriteBufferSize:1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var Ser = NewSessionM()

func WsRoute(){
	http.HandleFunc("/ws",wsStart)
	//启动session master
	go Ser.Start()
	http.ListenAndServe(config.WebSocketPort, nil)
}

func wsStart(w http.ResponseWriter, req *http.Request){

		con, err :=  upgrader.Upgrade(w,req,nil)
		//defer func() {
		//	con.Close()
		//}()
		if err != nil{
			log.Println("升级为websocket失败",err.Error())
			return
		}
		//ser := wsrobot.NewSessionM()
		client :=  NewSession()
		client.Addr = con.RemoteAddr().String()
		client.Conn = con
		client.TimeS = time.Now().Unix()
		//defer func() {
		//	Ser.Unregister <- client
		//}()
		//Ser.SetSession(client.Addr,client)
		//fmt.Println("内存地址。。。。。。。。",&Ser)
		//心跳检测
		//go Ser.HeartBeat(500)
		//session := ser.GetSessionById(client.Addr)
		//if session == nil{
		//	return
		//}
		//go client.Write()
		go client.Read()
	    go client.Write()
		//用户注册
	    Ser.Register <- client
}


func TimerClean(){
	defer func() {
		if r:=recover();r!=nil{
			fmt.Println("停止清理超时",)
		}
	}()
	Ser.HeartBeat(60)
	return
}