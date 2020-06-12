package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goWeakPass/define"
	"os"
)

func Loginmysql(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)

	db, err := sql.Open("mysql", config.UserName+":"+config.PassWord+"@tcp("+config.Host+":"+config.Port+")/mysql?charset=utf8")
	if err != nil {
		return false
	}
	if err := db.Ping(); err != nil {
		return false
	}
	//fmt.Println("connnect success")
	defer db.Close()
	define.Output(value)
	return true
}
