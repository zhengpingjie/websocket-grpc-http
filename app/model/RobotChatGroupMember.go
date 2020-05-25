package model

import (
	"errors"
	"time"
)

type RobotChatGroupMember struct {
	Id  int  `json:"id" form:"id"`
	GroupId int `json:"groupid" form:"groupid"`
	UserId   int `json:"userid"  from:"userid"`
	IsDel  int8 `json:"is_del"  from:"is_del"`
	CreateTime int  `json:"createtime" form:"createtime"`
	UpdateTime int  `json:"updatetime" form:"updatetime"`
}


func NewRobotChatGroupMember()*RobotChatGroupMember{
	return &RobotChatGroupMember{
		CreateTime:time.Now().Second(),
		UpdateTime:time.Now().Second(),
	}
}

//添加
func(this *RobotChatGroupMember)Add()bool{
	sql := "insert into robot_chat_group_member(`groupid`,`userid`,`createtime`)VALUE(?,?,?)"
	if _,err := add(sql,this.GroupId,this.UserId,this.CreateTime);err != nil{
		return false
	}else{
		return true
	}
}


func(this *RobotChatGroupMember)Exec()bool{
	sql := "UPDATE  robot_chat_group_member set is_del =? where  id=?"
	if _,err := exec(sql,this.IsDel,this.Id);err != nil{
		return false
	}else{
		return true
	}
}


//获取群列表
func(this  *RobotChatGroupMember)ListBy()(*[]map[string]string,error){
	sql := "select * from robot_chat_group_member where is_del=? and userid=?"
	fields := []interface{}{"id","name","status"}
	if list,err :=listBy(sql,fields,0,this.UserId);err != nil{
		return list,nil
	}
	return nil,errors.New("获取列表失败")
}

//获取群列表
func(this  *RobotChatGroupMember)ListByGroudId()(*[]map[string]string,error){
	sql := "select * from robot_chat_group_member where is_del=? and groupid=?"
	fields := []interface{}{"id","name","status"}
	if list,err :=listBy(sql,fields,0,this.GroupId);err != nil{
		return list,nil
	}
	return nil,errors.New("获取列表失败")
}
