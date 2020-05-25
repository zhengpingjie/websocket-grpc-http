package grpcrobot

type RobotModule struct {
	Module string
	Action string
	Data *RobotData
}

type RobotData struct {
	PlatId int32
	ClientId string
	Types int32
	FromId string
	RoomId string
	MsgToken string
	MsgId string
	Content map[string]interface{}
}

type AppResp struct {
	Code int32
	Msg string
	Data map[string]string
}

func NewRobotData()*RobotData{
	return &RobotData{}
}

func NewRobotModule()*RobotModule{
   return &RobotModule{
	   Module:"robot",
	   Action:"",
	   Data:&RobotData{},
	}
}

func NewAppResp()*AppResp{
	return &AppResp{
		Code:1,
		Msg:"success",
		Data:make(map[string]string),
	}
}


