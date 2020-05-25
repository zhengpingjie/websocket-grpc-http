package config

import "time"

const(
	ApiMode = "debug"  //debug | test

	HttpPort = ":8091"
	WebSocketPort = ":8089"

	// 超时时间
	ApiReadTimeout  = 120
	ApiWriteTimeout = 120

	//websocket 配置项

	WriteWait = 3600 * time.Second
	PongWait  =  3600 * time.Second
	PingPeriod = (PongWait*9)/10

	MaxMessageSize = 512

	//grcp 配置项
	GrpcPort =  ":8092"

)
