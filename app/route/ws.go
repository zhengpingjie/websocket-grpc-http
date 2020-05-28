package route

import (
	"fmt"
	"robot/app/controller/wsrobot"
	"time"
)



func WsRoute(){
	wsrobot.WsRoute()
}


func WsInit() {
	wsrobot.WebSocketInit()
}

func Task(){
	go func() {
		//首次延迟时间
		t := time.NewTimer(3*time.Second)
		defer t.Stop()
		for{
			select {
			case <-t.C:
				fmt.Println("开始清理")
				wsrobot.TimerClean()
				//间隔30秒执行
				t.Reset(20*time.Second)
			}
		}

	}()

}