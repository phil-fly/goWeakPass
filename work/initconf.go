package work

import (
	"fmt"
	"goWeakPass/distfile"
	"goWeakPass/distsql"
	"gopkg.in/gcfg.v1"
	"log"
)

type MysqlConfStu struct {
	Enabled  bool
	Host     string
	Dbport   string
	Username string
	Password string
	Dbname   string
	Userdist string
	Passdist string
}

type FileConfStu struct {
	Enabled  bool
	Userfile string
	Passfile string
}

var MysqlConf MysqlConfStu
var FileConf FileConfStu
var userlist_sql []distsql.Userdist
var passlist_sql []distsql.Passdist

var userlist_file []distfile.Userdist
var passlist_file []distfile.Passdist

func GetConf(confPath string) {
	config := struct {
		Mysqldist struct {
			Enabled  bool
			Host     string
			Dbport   string
			Username string
			Password string
			Dbname   string
			Userdist string
			Passdist string
		}
		Filedist struct {
			Enabled  bool
			Userfile string
			Passfile string
		}
	}{}
	err := gcfg.ReadFileInto(&config, confPath)
	if err != nil {
		fmt.Println("Failed to parse config file: %s", err)
	}
	MysqlConf = config.Mysqldist
	FileConf = config.Filedist
	if MysqlConf.Enabled {
		log.Print("加载MYSQL字典:", MysqlConf.Enabled)
		userlist_sql, passlist_sql = distsql.SqlDistGet(MysqlConf.Username, MysqlConf.Password, MysqlConf.Host, MysqlConf.Dbport, MysqlConf.Dbname, MysqlConf.Userdist, MysqlConf.Passdist)
		log.Print("读取mysql字典加载完成", MysqlConf.Enabled)
	}

	if FileConf.Enabled {
		log.Print("加载文件字典:", FileConf.Enabled)
		userlist_file, passlist_file = distfile.FlieDist_Get(FileConf.Userfile, FileConf.Passfile)
	}

}
