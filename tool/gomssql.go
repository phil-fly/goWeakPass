package tool

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"log"
)

func Loginmssql(host, username, password string, port int) string {
	db, err := gorm.Open("mssql", "sqlserver://"+username+":"+password+"@"+host+":"+"1433"+"?database=comsharp-cms")

//	dataSourceName := fmt.Sprintf("server=%s;port=%d;database=%s;trusted_connection=yes;", host,
//		port, "sa", password,)
	if err != nil {
		log.Print("用户名：", username, "    密码: ", password, "     ", "false", err.Error())
		return "false"
	}else{
		log.Print("用户名：", username, "    密码: ", password, "      ", "true")
		defer db.Close()
		return "true"
	}
}
