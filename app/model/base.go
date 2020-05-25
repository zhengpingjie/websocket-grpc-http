package model

import (
	"fmt"
	"robot/app/db"
)

//添加
func add(sqlstr string,args  ...interface{})(int64,error){
	stmt,err:=db.SqlDb.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil{
		fmt.Printf("insert data error:%v\n",err)
		return 0,err
	}
	ret,err := stmt.Exec(args...)
	if err!=nil{
		return 0,err
	}
	return ret.LastInsertId()
}


//修改和删除
func exec(sqlstr string,args  ...interface{})(int64,error){
	stmt,err:=db.SqlDb.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil{
		fmt.Printf("insert data error:%v\n",err)
		return 0,err
	}
	ret,err := stmt.Exec(args...)
	if err!=nil{
		return 0,err
	}
	return ret.RowsAffected()
}


//取一行数据
func infoBy(sqlstr string,fields []interface{},args ...interface{})(*[]map[string]string, error) {
	stmt,err:=db.SqlDb.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil{
		fmt.Printf("select data error:%v\n",err)
		return nil,err
	}
	row,err := stmt.Query(args...)
	if err != nil {
		fmt.Printf("select data error:%v\n",err)
		return nil,err
	}
	arrmap := make([]map[string]string,0)
	for row.Next(){
		err = row.Scan(fields...)
		if err != nil{
			fmt.Printf("select data error:%v\n",err.Error())
		}
		smap := make(map[string]string,len(fields))
		var value  string
		for key,val := range fields{
			if val == nil{
				value = ""
			}else{
				value = val.(string)
			}
			smap[fields[key].(string)] = value
		}
		arrmap = append(arrmap,smap)
		break
	}
	return &arrmap,nil
}


//取出多行数据
func listBy(sqlstr string,fields []interface{},args ...interface{})(*[]map[string]string, error){
	stmt,err:=db.SqlDb.Prepare(sqlstr)
	defer stmt.Close()
	if err != nil{
		fmt.Printf("select data error:%v\n",err)
		return nil,err
	}
	row,err := stmt.Query(args...)
	if err != nil {
		fmt.Printf("select data error:%v\n",err)
		return nil,err
	}
	arrmap := make([]map[string]string,0)
	for row.Next(){
		err = row.Scan(fields...)
		if err != nil{
			fmt.Printf("select data error:%v\n",err.Error())
		}
		smap := make(map[string]string,len(fields))
		var value  string
		for key,val := range fields{
			if val == nil{
				value = ""
			}else{
				value = val.(string)
			}
			smap[fields[key].(string)] = value
		}
		arrmap = append(arrmap,smap)
	}
	return &arrmap,nil
}