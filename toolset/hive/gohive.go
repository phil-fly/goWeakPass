package hive

import (
	"github.com/dazheng/gohive"
	"fmt"
	"goWeakPass/define"
	"os"
)

func LoginHive(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)
	//	conn, err := gohive.Connect("127.0.0.1:10000", gohive.DefaultOptions) // 无用户名、密码
	conn, err := gohive.ConnectWithUser(config.Host+":"+config.Port, config.UserName, config.PassWord, gohive.DefaultOptions) // 需要用户名、密码
	if err != nil {
		return false
	}
	conn.Close()
	define.Output(value)
	return true
}

