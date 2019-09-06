package tool

import (
	"fmt"
	"database/sql"
	//_ "github.com/denisenkom/go-mssqldb"
	"log"
)
func Loginmssql(host, username, password string, port int) string {
	dataSourceName := fmt.Sprintf("server=%v;port=%v;user id=%v;password=%v;database=%v", host,
		port, username, password, "master")
	db, err := sql.Open("mssql", dataSourceName)
	if err !=nil {
		log.Print("用户名：", username, "    密码: ", password, "      ", "false")
		return "false"
	}
	if err := db.Ping(); err == nil {
		log.Print("用户名：", username, "    密码: ", password, "      ", "true")
		return "true"
	}
	log.Print("用户名：", username, "    密码: ", password, "      ", "false")
	return "false"
}

