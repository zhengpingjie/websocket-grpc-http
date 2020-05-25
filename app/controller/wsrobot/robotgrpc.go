package wsrobot

import (
	"robot/app/util"
)


func getKey(clientId string)string{
	Addr := Ser.GetUser(clientId)
	return Addr
}
func UploadLocation(clientId,toClientId,action,word string)int64 {
	msg := make(map[string]interface{})
	Addr := getKey(toClientId)
	if Addr == "" {
		return 0
	}
	msgId :=  util.GetMsgId()
	msg["action"] = action
	msg["fromClientId"] = clientId
	msg["word"] = word
	if Ser.SendMsgById(Addr,msg){
		return msgId
	}
	return 0
}
//上报设备状态
func UploadStatus(clientId,toClientId,action,location string,chargeState,batteryLevel,currentTask,deviceId int64)int64{
	msg := make(map[string]interface{})
	Addr := getKey(toClientId)
	if Addr == ""{
		return 0
	}
	//fromAddr := getKey(clientId)
	msgId :=  util.GetMsgId()
	msg["action"] = action
	msg["msgId"] = msgId
	msg["fromClientId"] = clientId
	msg["location"] = location
	msg["chargeState"] = chargeState
	msg["batteryLevel"] = batteryLevel
	msg["currentTask"] = currentTask
	msg["deviceId"] = deviceId
	if Ser.SendMsgById(Addr,msg){
		return msgId
	}
	return 0
}