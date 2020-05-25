package robot

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"robot/app/model"
	"strconv"
)

//添加群组
func AddGroup(tx *gin.Context){
  m:= model.NewRobotChatGroup()
  if err :=tx.ShouldBindJSON(m);err != nil{
	if m.Name == ""{
		tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"请求参数不能为空"})
		return
	}
	if isok := m.Add();!isok{
		tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"插入数据失败"})
		return
	}
	  tx.JSON(http.StatusAccepted,gin.H{"code":"1","msg":"成功"})
  }

	tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"插入数据失败"})
	return
}


//删除|修改群组
func DelGroup(tx *gin.Context){
	m:= model.NewRobotChatGroup()
	if err :=tx.ShouldBindJSON(m);err != nil{
		if m.Id <=0{
			tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"请求参数不能为空"})
			return
		}
		if isok := m.Exec();!isok{
			tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"删除失败"})
			return
		}
		tx.JSON(http.StatusAccepted,gin.H{"code":"1","msg":"成功"})
	}
	tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"删除群组失败"})
	return
}

//删除群组机器人
func DelGroupRobot(tx *gin.Context){
  m := model.NewRobotChatGroupMember()
	if err :=tx.ShouldBindJSON(m);err != nil{
		if m.Id <=0{
			tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"请求参数不能为空"})
			return
		}
		if isok := m.Exec();!isok{
			tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"删除失败"})
			return
		}
		tx.JSON(http.StatusAccepted,gin.H{"code":"1","msg":"成功"})
	}
	tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"删除群组成员失败"})
	return
}

//添加群组机器人
func AddGroupRobot(tx *gin.Context){
	m := model.NewRobotChatGroupMember()
	if err :=tx.ShouldBindJSON(m);err != nil{
		if m.UserId >=0 && m.GroupId >=0 {
			if isok := m.Add();!isok{
				tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"插入数据失败"})
				return
			}
			tx.JSON(http.StatusAccepted,gin.H{"code":"1","msg":"成功"})
		}
	}
	tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"插入数据失败"})
	return
}

//获取群组机器人列表
func GetGroupRobot(tx *gin.Context){
	m:= model.NewRobotChatGroupMember()
	if err :=tx.ShouldBindJSON(m);err != nil{
       //获取所有的群列表
       if m.UserId == 0{
		   tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"参数不能为空"})
		   return
	   }
       list := make([]interface{},0)
      // grouplist := make([]map[string]string,0)
      if glist,err := m.ListBy();err == nil{
      		for _,v := range *glist{
				if len(v) == 0 {
					continue
				}
				if groupid,ok:= v["groupid"];ok{
					gid,_ := strconv.Atoi(groupid)
					m.GroupId = gid
					if grouplist,err := m.ListByGroudId();err!=nil{
						list = append(list,grouplist)
					}
				}
			}
	  }

      if len(list) > 0{
		  tx.JSON(http.StatusAccepted,gin.H{"code":"1","msg":"成功","data":list})
		  return
	  }
		tx.JSON(http.StatusAccepted,gin.H{"code":"0","msg":"失败"})
	}

}
