package ftp

import (
	"fmt"
	ftptool "github.com/smallfish/ftp"
	"goWeakPass/define"
	"os"
)

func LoginFtp(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)
	Aftp := new(ftptool.FTP)
	Aftp.Debug = true
	Aftp.Connect(config.Host, config.PortInt)
	Aftp.Login(config.UserName, config.PassWord)
	if Aftp.Code == 530 {
		return false
	}
	Aftp.Quit()
	define.Output(value)
	return true
}
