package robot

import (
	"github.com/gin-gonic/gin"
	"log"
	"robot/app/model"
)

//获取模板首页数据
func Gettemp(tx *gin.Context){
	temp := model.NewOrgTemp()
	if tx.ShouldBindQuery(temp) == nil{
		log.Println("=====Only Bind by Query String=======")
		info,err :=temp.GetRow()
		if err != nil{
			log.Fatal(err.Error())
		}
		log.Fatalf("%+v",info)

	}
	//resp := response.NewGin(tx)
}