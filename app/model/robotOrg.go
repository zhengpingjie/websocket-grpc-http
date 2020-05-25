package model


type RobotOrg struct {
	Id  int  `json:"id" form:"id"`
	Orgid int  `json:"indid" form:"indid"`
	Title string  `json:"orgid" form:"orgid"`
	Bg string  `json:"packid" form:"packid"`
	Tips string  `json:"answer" form:"answer"`
	Url string  `json:"question" form:"question"`
	Enable  int8  `json:"questions" form:"questions"`
	Logo string  `json:"answer_media" form:"answer_media"`
	QuestionPackType int8  `json:"question_pack_type" form:"question_pack_type"`
	TipsVoice string  `json:"tips_voice" form:"tips_voice"`
	Type int8  `json:"type" form:"type"`
	Isdel int8 `json:"is_del" form:"is_del"`
	Sort int  `json:"sort" form:"sort"`
	Cnt int  `json:"cnt" form:"cnt"`
	Createtime int  `json:"createtime" form:"createtime"`
	Updatetime int  `json:"updatetime" form:"updatetime"`
}

func NewRobotOrg()*RobotOrg{
	return &RobotOrg{}
}