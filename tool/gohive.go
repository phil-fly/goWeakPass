package tool

import (
	"github.com/dazheng/gohive"
	"fmt"
)

func LoginHive(host, username, password string, port int) string {
	//	conn, err := gohive.Connect("127.0.0.1:10000", gohive.DefaultOptions) // 无用户名、密码
	conn, err := gohive.ConnectWithUser(host+":"+fmt.Sprintf("%d",port), username, password, gohive.DefaultOptions) // 需要用户名、密码
	if err != nil {
		return "err"
	}
	conn.Close()
	return "true"
}

