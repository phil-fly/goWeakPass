package tool

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func LoginPostgres(host, username, password string, port int) string {

	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", username,
		password, host, port, "postgres", "disable")

	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	//db, err := sql.Open("postgres", connStr)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Print("用户名：", username, "    密码: ", password, "      ", "false")
		return "false"
	}
	if err := db.Ping(); err == nil {
		log.Print("用户名：", username, "    密码: ", password, "      ", "true")
		defer db.Close()
		return "true"
	}
	log.Print("用户名：", username, "    密码: ", password, "      ", "false")
	return "false"
}
