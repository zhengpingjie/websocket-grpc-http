package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"robot/app/config"
	"robot/app/protobuf"
	"time"
)

func main(){
	getIndexInfo()
}

func checkonline(){
	conn,err := grpc.Dial("127.0.0.1"+config.GrpcPort,grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer conn.Close()
	c := protobuf.NewRobotServerClient(conn)
	ctx,cancle:=context.WithTimeout(context.Background(),time.Second)
	defer cancle()
	req := protobuf.UsersOnlineReq{
		ClientId:"123456",
	}
	resp,err := c.UsersOnline(ctx,&req)
	if err !=nil{
		fmt.Println("给全体用户发送消息", err)
	}

	fmt.Println(resp.GetCode())
}


func sendMsg(){
	con,err := grpc.Dial("127.0.0.1"+config.GrpcPort,grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer con.Close()
	c := protobuf.NewRobotServerClient(con)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	req := protobuf.SendMsgReq{
		Action:"sendMsg",
		Content:"212121212",
		ToClientId:"123456",
	}
	resp,err := c.SendMsg(ctx,&req)
	if err !=nil{
		fmt.Println("给全体用户发送消息", err)
	}

	fmt.Println(resp.GetCode())

}


func startCall(){
	con,err := grpc.Dial("127.0.0.1"+config.GrpcPort,grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer con.Close()
	c := protobuf.NewRobotServerClient(con)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	req := protobuf.StartCallReq{
		ClientId:"79999",
		ToClientId:"123456",
		RoomId:"f4s545sa882",
		Action:"startCall",
	}
	resp,err := c.StartCall(ctx,&req)
	if err !=nil{
		fmt.Println("给全体用户发送消息", err)
	}
	fmt.Println(resp.GetCode())
}

func ToAction(){
	con,err := grpc.Dial("127.0.0.1"+config.GrpcPort,grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer con.Close()
	c := protobuf.NewRobotServerClient(con)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	req := protobuf.ToActionReq{
		ClientId:"8456556",
		ToClientId:"123456",
		Action:"ToAction",
		Locations:"前台，门口，技术部",
	}
	resp,err := c.ToAction(ctx,&req)
	if err !=nil{
		fmt.Println("给全体用户发送消息", err)
	}
	fmt.Println(resp.GetCode())
}

func getIndexInfo(){
	con,err := grpc.Dial("127.0.0.1"+config.GrpcPort,grpc.WithInsecure())
	if err != nil{
		fmt.Println(err.Error())
		return
	}
	defer con.Close()
	c := protobuf.NewRobotServerClient(con)
	ctx,cancel := context.WithTimeout(context.Background(),time.Second)
	defer cancel()
	req := protobuf.GetIndexInfoReq{
		ClientId:"465456",
		ToClientId:"123456",
		Action:"ToServser",
	}
	resp,err := c.GetIndexInfo(ctx,&req)
	if err !=nil{
		fmt.Println("给全体用户发送消息", err)
	}
	fmt.Println(resp.GetCode())
}