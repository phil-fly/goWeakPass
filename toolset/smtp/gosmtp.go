package smtp

import (
	"fmt"
	"goWeakPass/define"
	"gopkg.in/gomail.v2"
	"os"
)

func LoginSmtp(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)

	mailTo := []string{
		config.UserName, //实验邮件接收人为自己本身
	}

	subject := "弱口令测试"
	// 邮件正文
	body := ""

	mailConn := map[string]string{
		"user": config.UserName,
		"pass": config.PassWord,
		"host": config.Host,                   //"smtp.qq.com",
		"port": config.Port, //"465",
	}

	m := gomail.NewMessage()
	m.SetHeader("From", "weakpass"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                             //发送给多个用户
	m.SetHeader("Subject", subject)                          //设置邮件主题
	m.SetBody("text/html", body)                             //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], config.PortInt, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		return false
	}
	define.Output(value)
	return true

}
