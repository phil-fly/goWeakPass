package tool

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
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
	"aes128-gcm@openssh.com",
	chacha20Poly1305ID,
	"arcfour256", "arcfour128", "arcfour",
	aes128cbcID,
	tripledescbcID,
	gcmCipherID,
	"aes128-cbc",
}

func SshConnect(user, password, host string, port int) (string, error) {
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
	auth = append(auth, ssh.Password(password))

	clientConfig = &ssh.ClientConfig{
		User:    user,
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
	addr = fmt.Sprintf("%s:%d", host, port)

	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		log.Print("用户名：", user, "    密码: ", password, "      ", err.Error())
		return "", err
	}

	// create session
	if session, err = client.NewSession(); err != nil {
		log.Print("用户名：", user, "    密码: ", password, "      ", err.Error())
		return "", err
	}
	session.Close()
	return "ok", nil
}
