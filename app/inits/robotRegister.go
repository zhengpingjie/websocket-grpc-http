package inits

import (
	"log"
	"reflect"
	"robot/app/interfaces"
	"robot/app/util"
)

type RoutersMap struct {
	Controller  map[string]map[string]func( client interface{}, data map[string]interface{})bool
}

func NewRoutersMap()*RoutersMap{
	return &RoutersMap{
		Controller:make(map[string]map[string]func( client interface{}, data map[string]interface{})bool),
	}
}

//注册test结构
func(this *RoutersMap)RegisterRobotStructFun(actionName string,mod interfaces.RobotFace)bool{
	this.Controller[actionName] = make(map[string]func( interface{}, map[string]interface{})bool)
	temval := reflect.ValueOf(mod)
	temType := reflect.TypeOf(mod)
	for i:=0 ;i<temType.NumMethod();i++{
		tem := temval.Method(i).Interface()
		if temFunc,ok:= tem.(func( client interface{},data map[string]interface{})bool);ok{
			this.Controller[actionName][temType.Method(i).Name] = temFunc
		}
	}
	//this.MapToJsonDemo1()
	return true
}

//Hook test结构
func(this *RoutersMap)HookRobotAction(actionName,funcName string,client interface{},data map[string]interface{})bool{
	if _,exit := this.Controller[actionName];!exit{
		return false
	}
	if action,exit := this.Controller[actionName][funcName];exit{
		action(client,data)
	}
	return true
}



//处理参数
func(this *RoutersMap)DealParam(data []byte,client interface{}){
	m := util.JsonByteToMap(data)
	actionName := commaStringOk("Module",m)
	funcName := commaStringOk("Action",m)
	param := commDataOk("Data",m)
	//this.MapToJsonDemo1()

	if actionName != "" && funcName != ""{
		this.HookRobotAction(actionName,funcName,client,param)
	}
	return
}


func commaStringOk(name string,m map[string]interface{})string{
	if v,exist := m[name];exist{
		ac,ok :=v.(string)
		if !ok{
			log.Printf("不存在 *string, got:", v)
			return ""
		}
		return ac
	}
	return ""
}

func commDataOk(name string,m map[string]interface{})map[string]interface{}{
	if v,exist := m[name];exist{
		ac,ok :=v.(map[string]interface{})
		if !ok{
			log.Printf("不存在 *string, got:", v)
			return nil
		}
		return ac
	}
	return nil
}



