package wsrobot

import (
	"fmt"
	. "robot/app/controller/grpcrobot"
	"robot/app/util"
	"time"
)

type Robot struct {
}

func(this *Robot)Login(client  interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		cli.ClientId = data["FromId"].(string)
		cli.LoginTime = time.Now().Unix()
		//更新心跳时间
		cli.UpdTime()
		//设置userid 与连接id得映射
		cli.Sessions.SetUser(cli.ClientId,cli.Addr)
		data := make(map[string]interface{})
		data["code"] = 0;
		data["msg"] = "登录成功"
		data["action"] = "Login"
		cli.Conn.WriteJSON(data)
		//cli.Sessions.SendMsgById(cli.Addr,data)
	}
	return true
}

func(this *Robot)HeartBeat(client interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		cli.UpdTime()
		fmt.Println("更新连接时间",cli.TimeS)
	}
	return true
}

//通知机器人上报数据
func(this *Robot)UploadState(client interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		RobotResp := NewRobotModule()
		RobotResp.Data.PlatId = 2
		RobotResp.Data.Types = 1
		RobotResp.Action = "UploadState"
		RobotResp.Data.FromId = data["fromId"].(string)
		clientId:= util.Md5V3(data["clientId"].(string))
		//clientId:= data["clientId"].(string)
		RobotResp.Data.ClientId = clientId
		Addr := cli.Sessions.GetUser(clientId)
		if Addr == ""{
			cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(1,1,"Response","用户不在线",data["msgToken"]))
			return false
		}
		cli.Sessions.GrpcSendMsgById(Addr,RobotResp)
		//通知消息发送成功
		cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(0,0,"Response","发送成功",data["msgToken"]))
		//机器人绑定设备
		cli.Sessions.SetBinds(RobotResp.Data.FromId,clientId)
		//cli.Sessions.SendMsgById(cli.Addr,appResp)
	}
	return true
}

//加入房间
func(this *Robot)AddLiveRoom(client interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		RobotResp := NewRobotModule()
		RobotResp.Data.PlatId = 2
		RobotResp.Data.Types = 1
		RobotResp.Action = "AddLiveRoom"
		RobotResp.Data.FromId = data["fromId"].(string)
		content := make(map[string]interface{})
		content["roomId"] = data["roomId"]
		RobotResp.Data.Content = content
		clientId:= util.Md5V3(data["clientId"].(string))
		//clientId:= data["clientId"].(string)
		RobotResp.Data.ClientId = clientId
		Addr := cli.Sessions.GetUser(clientId)
		if Addr == ""{
			cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(1,1,"Response","用户不在线",data["msgToken"]))
			return false
		}

		cli.Sessions.GrpcSendMsgById(Addr,RobotResp)
		//通知消息发送成功
		cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(0,0,"Response","发送成功",data["msgToken"]))
	}
	return true
}

//控制机器人移动
func(this *Robot)RobotMove(client interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		RobotResp := NewRobotModule()
		RobotResp.Data.PlatId = 2
		RobotResp.Data.Types = 1
		RobotResp.Action = "RobotMove"
		RobotResp.Data.FromId = data["fromId"].(string)

		content := make(map[string]interface{})
		content["action"] = data["action"]
		content["flag"] = data["flag"]
		RobotResp.Data.Content = content
		clientId:= util.Md5V3(data["clientId"].(string))
		//clientId:= data["clientId"].(string)
		RobotResp.Data.ClientId = clientId
		Addr := cli.Sessions.GetUser(clientId)
		if Addr == ""{
			cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(1,1,"Response","用户不在线",data["msgToken"]))
			return false
		}
		cli.Sessions.GrpcSendMsgById(Addr,RobotResp)
		//通知消息发送成功
		cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(0,0,"Response","发送成功",data["msgToken"]))
	}
	return true
}

//通知机器人巡逻
func(this *Robot)RobotPartrol(client interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		RobotResp := NewRobotModule()
		RobotResp.Data.PlatId = 2
		RobotResp.Data.Types = 1
		RobotResp.Action = "RobotPartrol"
		RobotResp.Data.FromId = data["fromId"].(string)

		content := make(map[string]interface{})
		content["location"] = data["location"]
		content["flag"] = data["flag"]
		RobotResp.Data.Content = content
		clientId:= util.Md5V3(data["clientId"].(string))
		//clientId:= data["clientId"].(string)
		RobotResp.Data.ClientId = clientId
		Addr := cli.Sessions.GetUser(clientId)
		if Addr == ""{
			cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(1,1,"Response","用户不在线",data["msgToken"]))
			return false
		}

		cli.Sessions.GrpcSendMsgById(Addr,RobotResp)
		//通知消息发送成功
		cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(0,0,"Response","发送成功",data["msgToken"]))
	}
	return true
}

//通知机器人说话
func(this *Robot)RobotSpeak(client interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		RobotResp := NewRobotModule()
		RobotResp.Data.PlatId = 2
		RobotResp.Data.Types = 1
		RobotResp.Action = "RobotSpeak"
		RobotResp.Data.FromId = data["fromId"].(string)

		content := make(map[string]interface{})
		content["word"] = data["word"]
		RobotResp.Data.Content = content
		clientId:= util.Md5V3(data["clientId"].(string))
		//clientId:= data["clientId"].(string)
		RobotResp.Data.ClientId = clientId
		Addr := cli.Sessions.GetUser(clientId)
		if Addr == ""{
			cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(1,1,"Response","用户不在线",data["msgToken"]))
			return false
		}
		cli.Sessions.GrpcSendMsgById(Addr,RobotResp)
		//通知消息发送成功
		cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(0,0,"Response","发送成功",data["msgToken"]))
	}
	return true
}

//通知机器去指定地点
func(this  *Robot)RobotToLocation(client interface{},data map[string]interface{})bool{
	if cli,ok:= client.(*Session);ok{
		RobotResp := NewRobotModule()
		RobotResp.Data.PlatId = 2
		RobotResp.Data.Types = 1
		RobotResp.Action = "RobotToLocation"
		RobotResp.Data.FromId = data["fromId"].(string)

		content := make(map[string]interface{})
		content["location"] = data["location"]
		RobotResp.Data.Content = content
		clientId:= util.Md5V3(data["clientId"].(string))
		//clientId:= data["clientId"].(string)
		RobotResp.Data.ClientId = clientId
		Addr := cli.Sessions.GetUser(clientId)
		if Addr == ""{
			cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(1,1,"Response","用户不在线",data["msgToken"]))
			return false
		}
		cli.Sessions.GrpcSendMsgById(Addr,RobotResp)
		//通知消息发送成功
		cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(0,0,"Response","发送成功",data["msgToken"]))
	}
	return true
}

//切换设备
func(this  *Robot)SwitchRobot(client interface{},data map[string]interface{})bool{
	//切换设备通知机器人 小程序端切换设备
	if cli,ok:= client.(*Session);ok{
		RobotResp := NewRobotModule()
		RobotResp.Data.PlatId = 2
		RobotResp.Data.Types = 1
		RobotResp.Action = "Disconnect"
		RobotResp.Data.FromId = data["fromId"].(string)
		clientId:= util.Md5V3(data["toOldClientId"].(string))
		//clientId:= data["toOldClientId"].(string)
		RobotResp.Data.ClientId = clientId
		Addr := cli.Sessions.GetUser(clientId)
		if Addr == ""{
			cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(1,1,"Response","用户不在线",data["msgToken"]))
			return false
		}
		cli.Sessions.GrpcSendMsgById(Addr,RobotResp)
		//通知消息发送成功
		cli.Sessions.SendMsgById(cli.Addr,util.AppReturn(0,0,"Response","发送成功",data["msgToken"]))
	 }
	return true
}
