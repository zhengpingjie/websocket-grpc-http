package wsrobot

import (
	"robot/app/inits"
)

var router = inits.NewRoutersMap()
func WebSocketInit(){
	router.RegisterRobotStructFun("robot",&Robot{})
}