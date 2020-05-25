package robot

import (
	"github.com/gin-gonic/gin"
	"log"
	"robot/app/model"
)

func Getdeviceinfo(tx  *gin.Context){
	q := model.NewRobotOrg()
	if tx.ShouldBindQuery(q) == nil{
		log.Println("=====Only Bind by Query String=======");
	}

}