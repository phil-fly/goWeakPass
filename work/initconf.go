package work

import (
	"gopkg.in/gcfg.v1"
	"fmt"
	"log"
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
	log.Print("数据库配置开启状态：",MysqlConf.Enabled)
	log.Print("数据库连接地址：",MysqlConf.Host)
	log.Print("数据库名：",MysqlConf.Dbname)
	log.Print("数据库账户：",MysqlConf.Username)
	log.Print("数据库密码：",MysqlConf.Password)

}
