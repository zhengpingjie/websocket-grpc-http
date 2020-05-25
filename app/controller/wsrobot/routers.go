package wsrobot



type RouteMap struct {
	actions map[string]func()bool
	modules map[string]map[string]func()bool
}

//实例化RouteMap

func NewRouteMap()*RouteMap{
	return &RouteMap{
		actions:make(map[string]func()bool),
		modules:make(map[string]map[string]func()bool),
	}
}


//注册事件

//注册单个逻辑

//结构体 注册