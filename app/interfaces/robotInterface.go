package interfaces

type RobotFace interface {
	//登录
	Login(client interface{},data map[string]interface{})bool
	//心跳连接
	HeartBeat(client interface{},data map[string]interface{})bool
	//通知机器人上报数据
	UploadState(client interface{},data map[string]interface{})bool
	//加入房间
	AddLiveRoom(client interface{},data map[string]interface{})bool
	//控制机器人移动
	RobotMove(client interface{},data map[string]interface{})bool
	// 通知机器人巡逻
	RobotPartrol(client interface{},data map[string]interface{})bool
	//通知机器人说话
	RobotSpeak(client interface{},data map[string]interface{})bool
	//通知机器去指定地点
	RobotToLocation(client interface{},data map[string]interface{})bool
	//切换设备
	SwitchRobot(client interface{},data map[string]interface{})bool
}