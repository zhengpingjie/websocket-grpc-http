package grpcserver

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	"robot/app/config"
	"robot/app/controller/wsrobot"
	"robot/app/protobuf"
	"robot/app/util"
)

type robotserver struct {
}
func (this *robotserver)UploadLocation(ctx context.Context, req *protobuf.UploadLocationReq) (*protobuf.Rsponses, error){
	fmt.Println("grpc上报当前位置",req.String())
	msgId := wsrobot.UploadLocation(req.GetClientId(),req.GetToClientId(),req.GetAction(),req.GetWord())
	rsp := &protobuf.Rsponses{}
	if msgId == 0{
		rsp.Code=0
		rsp.Msg = "用户不在线"
		rsp.MsgId = util.GetMsgId()
	}else{
		rsp.Code=1
		rsp.Msg = "发送成功"
		rsp.MsgId = msgId
	}
	return rsp,nil
}
//上报当前机器人状态
func (this *robotserver)UploadStatus(ctx context.Context,req *protobuf.UploadReq)(*protobuf.Rsponses, error) {
	fmt.Println("grpc获取当前任务|机器人电量|当前位置|是否在线",req.String())
	msgId := wsrobot.UploadStatus(req.GetClientId(),req.GetToClientId(),req.GetAction(),req.GetLocation(),req.GetChargeState(),req.GetBatteryLevel(),req.GetCurrentTask(),req.GetDeviceId())
	rsp := &protobuf.Rsponses{}
	if msgId == 0{
		rsp.Code=0
		rsp.Msg = "用户不在线"
		rsp.MsgId = util.GetMsgId()
	}else{
		rsp.Code=1
		rsp.Msg = "发送成功"
		rsp.MsgId = msgId
	}
	return rsp,nil
}



func Init(){
	rpcPort := config.GrpcPort
	fmt.Println("rpc server 启动", rpcPort)
	lis, err :=net.Listen("tcp",rpcPort)
	if err !=nil{
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	protobuf.RegisterRobotServerServer(s ,&robotserver{})
	if  err:= s.Serve(lis);err!=nil{
		log.Fatalf("failed to serve: %v", err)
	}

}