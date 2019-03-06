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
func Task(proto string,tasknum int,hostaddr string){
	host = hostaddr
	tasks := make(chan Workdist,taskload)
	wg.Add(tasknum)
	for gr:=1;gr<=tasknum;gr++ {
		if proto == "ssh" {
			go	sshWorker(tasks)
		}else if proto == "telnet" {
			go	telnetWorker(tasks)
		}else if proto == "ftp" {
			go	ftpWorker(tasks)
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

