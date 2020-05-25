package wsrobot

type Sev struct {
	SessionMaster *SessionM
}

func NewServer()*Sev{
	wsConn := &Sev{}
	wsConn.SessionMaster = NewSessionM()
	return wsConn
}



