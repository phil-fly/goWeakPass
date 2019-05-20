package tool

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)

func Loginmysql(host,username,password string,port int) string {

	db, err := sql.Open("mysql",username+":"+ password +"@tcp("+ host +":"+ strconv.Itoa(port) + ")/mysql?charset=utf8")
	if err != nil {
		return "opon database fail"
	}
	if err := db.Ping(); err != nil{
		return "opon database fail"
	}
	//fmt.Println("connnect success")
	defer db.Close()
	return "ok"
}


