package tool

import (
	"github.com/smallfish/ftp"
	"log"
)

func LoginFtp(host,username,password string,port int ) string {

	ftp := new(ftp.FTP)
	ftp.Debug = true
	ftp.Connect(host, port)
	ftp.Login(username, password)
	if ftp.Code == 530 {
		return "530"
	}
	log.Print("登录成功")
	ftp.Quit()
	return "230"
}







