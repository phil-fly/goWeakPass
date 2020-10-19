package sshLogin

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goWeakPass/define"
	"golang.org/x/crypto/ssh"
	"net"
	"os"
	"time"
)
const chacha20Poly1305ID = "chacha20-poly1305@openssh.com"
const (
	gcmCipherID    = "aes128-gcm@openssh.com"
	aes128cbcID    = "aes128-cbc"
	tripledescbcID = "3des-cbc"
)

var supportedCiphers = []string{
	"aes128-ctr", "aes192-ctr", "aes256-ctr",
	chacha20Poly1305ID,
	"arcfour256", "arcfour128", "arcfour",
	aes128cbcID,
	tripledescbcID,
	gcmCipherID,
}

func LoginSsh(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)
	var (
		auth         []ssh.AuthMethod
		addr         string
		clientConfig *ssh.ClientConfig
		client       *ssh.Client
		session      *ssh.Session
		err          error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(config.PassWord))

	clientConfig = &ssh.ClientConfig{
		User:    config.UserName,
		Auth:    auth,
		Timeout: 6 * time.Second,
		//2019.6.18  golang默认配置加密方式不包括aes128-cbc  连接交换机需要使用aes128-cbc
		Config: ssh.Config{
			Ciphers: supportedCiphers,
		},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", config.Host, config.PortInt)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		//log.Print("用户名：", user, "    密码: ", password, "      ", err.Error())
		return false
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		//log.Print("用户名：", user, "    密码: ", password, "      ", err.Error())
		return false
	}
	session.Close()
	define.Output(value)
	return true
}
