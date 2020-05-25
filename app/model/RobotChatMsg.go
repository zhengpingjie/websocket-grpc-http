package model

import "time"

type RobotChatMsg struct {
	Id  int  `json:"id" form:"id"`
	AppId int8 `json:"appId" form:"appId"`
	FromUserId   int `json:"from_userid"  from:"from_userid"`
	FromGroupId   int `json:"from_groupid"  from:"from_groupid"`
	ToUserId  int `json:"to_userid"  from:"to_userid"`
	MsgType  int8 `json:"msgtype"  from:"msgtype"`
	Content  string `json:"content"  from:"content"`
	Status   int8 `json:"status"  from:"status"`
	CreateTime int  `json:"createtime" form:"createtime"`
	UpdateTime int  `json:"updatetime" form:"updatetime"`
}


func NewRobotChatMsg()*RobotChatMsg{
	return &RobotChatMsg{
		CreateTime:time.Now().Second(),
		UpdateTime:time.Now().Second(),
	}
}