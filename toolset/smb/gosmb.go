package smblogin

import (
	"fmt"
	"github.com/stacktitan/smb/smb"
	"goWeakPass/define"
	"os"
)

func LoginSmb(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)

	options := smb.Options{
		Host:        config.Host,
		Port:        config.PortInt,
		User:        config.UserName,
		Password:    config.PassWord,
		Domain:      "",
		Workstation: "",
	}
	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			define.Output(value)
			return true
		}
	}
	//log.Print("用户名：", user, "    密码: ", password, "      ", "false")
	return false
}
