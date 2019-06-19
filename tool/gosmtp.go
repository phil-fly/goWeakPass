package tool

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func Checksmtp(host, username, password string, smtpport int) error {

	mailTo := []string{
		username, //实验邮件接收人为自己本身
	}

	subject := "弱口令测试"
	// 邮件正文
	body := ""

	mailConn := map[string]string{
		"user": username,
		"pass": password,
		"host": host,                   //"smtp.qq.com",
		"port": strconv.Itoa(smtpport), //"465",
	}

	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	m.SetHeader("From", "weakpass"+"<"+mailConn["user"]+">") //这种方式可以添加别名，即“XD Game”， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code> 读者可以自行实验下效果
	m.SetHeader("To", mailTo...)                             //发送给多个用户
	m.SetHeader("Subject", subject)                          //设置邮件主题
	m.SetBody("text/html", body)                             //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	return err

}
