package robot

import (
	"github.com/gin-gonic/gin"
	"log"
	"robot/app/model"
)

func GetqList(tx *gin.Context){
	q := model.NewCommonvList()
	if tx.ShouldBindQuery(q) == nil{
		log.Println("=====Only Bind by Query String=======");
	}
}

func GetqInfo(tx  *gin.Context){
	q := model.NewCommonvList()
	if tx.ShouldBindQuery(q) == nil{
		log.Println("=====Only Bind by Query String=======");
	}

}

func GetqLike(tx  *gin.Context){
	q := model.NewCommonvList()
	if tx.ShouldBindQuery(q) == nil{
		log.Println("=====Only Bind by Query String=======");
	}

}