package work

import (
	"sync"
	"goWeakPass/distsql"
	"goWeakPass/tool"
	"log"
	"os"
)

type Workdist struct {
	Username	string
	Password 	string
}

const (
	taskload		    = 100
)

var wg sync.WaitGroup
var host string

var userlist []distsql.Userdist
var passlist []distsql.Passdist

func Taskinit(file string){
	//加载配置
	GetConf(file)
	if MysqlConf.Enabled == true {
		distsql.Sqlinit(MysqlConf.Username,MysqlConf.Password,MysqlConf.Host,MysqlConf.Dbname)
		userlist = distsql.Userlist
		passlist = distsql.Passlist

	}

}


func Taskrun(proto string,tasknum int,hostaddr string){
	host = hostaddr
	tasks := make(chan Workdist,taskload)
	wg.Add(tasknum)

	for gr:=1;gr<=tasknum;gr++ {
		switch {
			case proto == "ssh" :
				go	sshWorker(tasks)
			case proto == "telnet":
				go	telnetWorker(tasks)
			case proto == "ftp":
				go	ftpWorker(tasks)
			case proto == "mysql":
				go	mysqlWorker(tasks)
			default:
				return
		}
	}

	for _,U := range distsql.Userlist {
		for _,P := range distsql.Passlist {
			task := Workdist{
				Username:U.Username,
				Password:P.Password,
			}
			tasks <- task
		}
	}
	close(tasks)
	wg.Wait()
}

func sshWorker(tasks chan Workdist){
	defer wg.Done()

	for{
		task,ok := <- tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		log.Print("检测ssh服务弱口令：    用户名: ",task.Username,"   密码: ",task.Password)
		_,err:= tool.SshConnect(task.Username,task.Password,host,22)
		if err == nil{
			log.Print("检测到ssh服务弱口令：    用户名: ",task.Username,"   密码: ",task.Password)
			os.Exit(1)
		}
	}

}


func telnetWorker(tasks chan Workdist){
	defer wg.Done()
	for{
		task,ok := <- tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret,_:= tool.Telnet_Creat(host,task.Username,task.Password)
		if ret == true{
			log.Print("检测到telnet服务弱口令：    用户名: ",task.Username,"   密码: ",task.Password)
			os.Exit(1)
		}
	}
}

func ftpWorker(tasks chan Workdist){
	defer wg.Done()
	for{
		task,ok := <- tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret:= tool.LoginFtp(host,task.Username,task.Password)
		if ret == "230"{
			log.Print("检测到ftp服务弱口令：    用户名: ",task.Username,"   密码: ",task.Password)
			os.Exit(1)
		}
	}
}

func mysqlWorker(tasks chan Workdist){
	defer wg.Done()
	for{
		task,ok := <- tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		log.Print("正在检测mysql服务弱口令：    用户名: ",task.Username,"   密码: ",task.Password)
		ret:= tool.Loginmysql(host,task.Username,task.Password)
		if ret == "ok"{
			log.Print("检测到mysql服务弱口令：    用户名: ",task.Username,"   密码: ",task.Password)
			os.Exit(1)
		}
	}
}
