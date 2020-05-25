package model

import (
	"time"
)

type RobotChatGroup struct {
	Id  int  `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Status   int8 `json:"status"  from:"status"`
	IsDel  int8 `json:"is_del"  from:"is_del"`
	CreateTime int  `json:"createtime" form:"createtime"`
	UpdateTime int  `json:"updatetime" form:"updatetime"`
}


func NewRobotChatGroup()*RobotChatGroup{
	return &RobotChatGroup{
		Id:0,
		Status:0,
		IsDel:0,
		CreateTime:time.Now().Second(),
		UpdateTime:time.Now().Second(),
	}
}

//添加群组
func(this *RobotChatGroup)Add()bool{
	sql := "insert into robot_chat_group(`name`,`createtime`)VALUE(?,?)"
	if _,err := add(sql,this.Name,this.CreateTime);err != nil{
		return false
	}else{
		return true
	}
}

//删除|修改
func(this *RobotChatGroup)Exec()bool{
	sql := "UPDATE  robot_chat_group set is_del =? where  id=?"
	if _,err := exec(sql,this.IsDel,this.Id);err != nil{
		return false
	}else{
		if this.IsDel == 1{
			sql = "UPDATE  robot_chat_group_member set is_del =? where  groupid=?"
			exec(sql,this.IsDel,this.Id)
		}
		return true
	}
}

