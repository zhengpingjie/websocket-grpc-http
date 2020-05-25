package robot

import (
	"github.com/gin-gonic/gin"
	"log"
	"robot/app/model"
)


//获取模板首页数据
func GetTourList(tx *gin.Context){
	temp := model.NewOrgTemp()
	if tx.ShouldBindQuery(temp) == nil{
		log.Println("=====Only Bind by Query String=======");
	}
	//resp := response.NewGin(tx)
}

//上报位置
func upload(tx *gin.Context){

}