package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"goWeakPass/define"
	"os"
)

func LoginPostgres(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)

	dataSourceName := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", config.UserName,
		config.PassWord, config.Host, config.Port, "postgres", "disable")

	//connStr := "postgres://pqgotest:password@localhost/pqgotest?sslmode=verify-full"
	//db, err := sql.Open("postgres", connStr)

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		//log.Print("用户名：", username, "    密码: ", password, "      ", "false")
		return false
	}
	if err := db.Ping(); err == nil {
		//log.Print("用户名：", username, "    密码: ", password, "      ", "true")
		defer db.Close()
		define.Output(value)
		return true
	}
	//log.Print("用户名：", username, "    密码: ", password, "      ", "false")
	return false
}
