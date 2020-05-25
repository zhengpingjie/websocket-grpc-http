package response

import(
	"github.com/gin-gonic/gin"
)
type Response struct{
	Code int  `json:"code"`
	Msg string `json:"msg"`
	Data  interface{} `json:"data"`
}

type resp struct {
	ctx *gin.Context
}

func NewGin(tx *gin.Context)*resp{
	return &resp{
		tx,
	}
}

func(this *resp)Resp(code int,msg string,data interface{}){
	this.ctx.JSON(200,Response{
		Code : code,
		Msg  : msg,
		Data : data,
	})
}


