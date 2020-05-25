package model

import (
	"log"
	"robot/app/db"
	"time"
)

type OrgTemp struct {
	Instanceid  int `json:"instanceId" form:"instanceId"`
	Deviceid  int `json:"deviceId" form:"deviceId"`
	Id  int  `json:"id" form:"id"`
	Orgid int  `json:"orgid" form:"orgid"`
	Title string  `json:"title" form:"title"`
	Bg string  `json:"bg" form:"bg"`
	Tips string  `json:"tips" form:"tips"`
	Url string  `json:"url" form:"url"`
	Enable  int8  `json:"enable" form:"enable"`
	Logo string  `json:"logo" form:"logo"`
	CreateTime int  `json:"createtime" form:"createtime"`
	Updateime int  `json:"updatetime" form:"updatetime"`
}

func NewOrgTemp()*OrgTemp{
	return &OrgTemp{}
}

//GetRow
func(this *OrgTemp)GetRow()(temp OrgTemp,err error){

	info := OrgTemp{}
	sql := "select * from robot_org_temp where orgid=? and enable = 1"
	log.Println(this.Id);
	return 
	db.SqlDb.QueryRow(sql,this.Id).Scan(&info)

	return info,err
}

//listBy
func(this *OrgTemp)GetRows()(temps []OrgTemp,err error ){
	rows, err := db.SqlDb.Query("select * from robot_org_temp")
	for rows.Next(){
		info := OrgTemp{}
		err := rows.Scan(&info)
		if err !=nil{
			log.Fatal(err.Error())
		}
		temps = append(temps,info)
	}
	rows.Close()
	return temps,nil

}

//add
func(this  *OrgTemp)Create()int64{
	sql := "INSERT into robot_org_temp (orgid,title,bg,tips,url,enable,logo,createtime)value(?,?,?,?,?,?,?,?)"
	rs,err :=db.SqlDb.Exec(sql,this.Orgid,this.Title,this.Bg,this.Tips,this.Enable,this.Logo,time.Now().Unix())
	if err != nil{
		log.Fatal(err.Error())
	}
	id,err := rs.LastInsertId()
	if err != nil{
		log.Fatal(err.Error())
	}
	return id
}

//updBy
func(this *OrgTemp)Update()int64{
	rs, err := db.SqlDb.Exec("update robot_org_temp set orgid = ? where id = ?", this.Orgid, this.Id)
	if err != nil{
		log.Fatal(err.Error())
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
//delBy

func Delete(id int) int64  {
	rs, err := db.SqlDb.Exec("delete from robot_org_temp where id = ?", id)
	if err != nil {
		log.Fatal()
	}
	rows, err := rs.RowsAffected()
	if err != nil {
		log.Fatal()
	}
	return rows
}