package util

import (
	"encoding/json"
	"fmt"
)

func JsonByteToMap(param []byte )map[string]interface{}{
	m := make(map[string]interface{})
	err := json.Unmarshal(param,&m)
	if err !=nil{
		fmt.Println("数据错误:",err.Error())
		return nil
	}
	fmt.Println(m["Action"])
	data := m["Data"]
	if v,ok := data.(map[string]interface{});ok{
		fmt.Println(v)
	}
	return m
}