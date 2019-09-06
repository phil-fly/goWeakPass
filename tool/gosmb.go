package tool

import (

	"github.com/stacktitan/smb/smb"
	"log"
)

func SmbConnect(user, password, host string, port int) (string, error) {
	options := smb.Options{
		Host:        host,
		Port:        port,
		User:        user,
		Password:    password,
		Domain:      "",
		Workstation: "",
	}
	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			return "true",nil
		}
	}
	log.Print("用户名：", user, "    密码: ", password, "      ", "false")
	return "false",nil
}