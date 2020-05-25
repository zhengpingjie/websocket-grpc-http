package route

import (
	"github.com/gin-gonic/gin"
	"robot/app/controller/robot"
	"robot/app/util"
)
func HttpRouter(engine *gin.Engine){
	
	//设置路由中间件
	
	//404
	engine.NoRoute(func(tx *gin.Context) {
		resp := util.NewGin(tx)
		resp.Resp(404,"请求方法不存在",nil)
	})
	//测试链路追踪
	
	//


	RobotTouter := engine.Group("/api")
	{
		//添加群组
		RobotTouter.GET("addGroup",robot.AddGroup)
		//删除|修改群组
		RobotTouter.GET("delGroup",robot.DelGroup)
		//删除|修改群组机器人
		RobotTouter.GET("delGroupRobot",robot.DelGroupRobot)
		//添加机器人到群组
		RobotTouter.GET("addGroupRobot",robot.AddGroupRobot)
		//获取群组机器人列表
		RobotTouter.GET("getGroupRobot",robot.GetGroupRobot)
		//获取所有的机器人
		RobotTouter.GET("getAllRobot",robot.GetAllRobot)
		//获取机器人消息
		RobotTouter.GET("getRobotMsg",robot.GetRobotMsg)
		//设备绑定
		RobotTouter.GET("devicebind",robot.GetRobotMsg)
		//获取首页配置
		RobotTouter.GET("getTempIndex",robot.Gettemp)
		//获取问题推荐列表
		//RobotTouter.GET("getQList",robot.GetqList)
		////根据问题ID获取问题详情
		//RobotTouter.GET("getQinfo",robot.GetqInfo)
		////根据问题获取问题详情
		//RobotTouter.GET("getQlike",robot.GetqLike)
		////获取设备配置(空闲超时,  接待点，巡逻)
		//RobotTouter.GET("getdevice",robot.Getdeviceinfo)
		////获取导览讲解配置
		//RobotTouter.GET("getTourList",)
		////上报位置
		//RobotTouter.GET("uploadLocations")
		//// 获取问路引领配置
		//RobotTouter.GET("getGuideList")
		////获取用户信息
		//RobotTouter.GET("getUsers")

	}

}


