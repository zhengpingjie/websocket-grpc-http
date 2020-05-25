package model

import "time"

type RobotChatInfo struct {
	Id  int  `json:"id" form:"id"`
	AppId int8 `json:"appId" form:"appId"`
	UserId   int `json:"userid"  from:"userid"`
	Status   int8 `json:"status"  from:"status"`
	IsDel  int8 `json:"is_del"  from:"is_del"`
	CreateTime int  `json:"createtime" form:"createtime"`
	UpdateTime int  `json:"updatetime" form:"updatetime"`
}


func NewRobotChatInfo()*RobotChatInfo{
	return &RobotChatInfo{
		CreateTime:time.Now().Second(),
		UpdateTime:time.Now().Second(),
	}
}