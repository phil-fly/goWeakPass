package main

import (
	"gopkg.in/gcfg.v1"
	"fmt"
	"./distsql"
	"flag"
	"./work"
)

type MysqlConfStu struct {
	Enabled bool
	Host	string
	Username    string
	Password    string
	Dbname		string
}

type FileConfStu struct {
	Enabled bool
	Path	string
	Userfile    string
	Passfile	string
}

var MysqlConf MysqlConfStu
var FileConf FileConfStu


func GetConf(confPath string){
	config := struct {
		Mysqldist struct {
			Enabled bool
			Host	string
			Username    string
			Password    string
			Dbname		string
		};
		Filedist struct {
			Enabled bool
			Path	string
			Userfile    string
			Passfile	string
		}
	}{}
	err := gcfg.ReadFileInto(&config, confPath)
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	MysqlConf = config.Mysqldist
	FileConf = config.Filedist
	//log.Print("数据库配置开启状态：",MysqlConf.Enabled)
	//log.Print("数据库连接地址：",MysqlConf.Host)
	//log.Print("数据库名：",MysqlConf.Dbname)
	//log.Print("数据库账户：",MysqlConf.Username)
	//log.Print("数据库密码：",MysqlConf.Password)

}

var proto = flag.String("proto", "null", "Weak password detection protos (ssh/telnet)")
var hostaddr = flag.String("host", "null", "Weak password detection hostaddr")
var tasknum = flag.Int("p", 1, "Weak password detection Number of threads")
var confpath = flag.String("conf", "config/conf.ini", "Weak password detection confpath")

func main() {
	//获取命令行参数
	//  main.exe -host 192.168.0.92 -proto ssh -p 50 -conf  confpath/conf.ini
	flag.Parse()

	//加载配置
	GetConf(*confpath)
	if MysqlConf.Enabled == true {
		distsql.Sqlinit(MysqlConf.Username,MysqlConf.Password,MysqlConf.Host,MysqlConf.Dbname)
		distsql.Userlist = distsql.GetUserList()
		distsql.Passlist = distsql.GetPassList()
	}

	work.Task(*proto,*tasknum,*hostaddr)
}