package tool


import (  
  "fmt"
  "time"
  "golang.org/x/crypto/ssh"
  "net"
	_ "github.com/go-sql-driver/mysql"
  "log"
)

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
      Ciphers: []string{"aes128-cbc"},
    },
    HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
        },
  }

  // connet to ssh
  addr = fmt.Sprintf("%s:%d", host, port)

  if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
    log.Print("用户名：",user,"    密码: ",password ,"      ",err.Error())
    return "", err
  }

  // create session
  if session, err = client.NewSession(); err != nil {
    log.Print("用户名：",user,"    密码: ",password ,"      ",err.Error())
    return "", err
  }
  session.Close()
  return "ok", nil
}