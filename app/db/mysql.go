package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


var SqlDb *sql.DB
func nit(){
	var err error
	SqlDb,err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/app_server?charset=utf8")
	if err != nil{
		log.Fatal(err.Error())
	}
	err = SqlDb.Ping()
	if err != nil{
		log.Fatal(err.Error())
	}

}