package work

import (
	"goWeakPass/tool"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
)

type Workdist struct {
	Username string
	Password string
	Port     int
}

const (
	taskload = 100
)

var wg sync.WaitGroup
var host string

func Taskinit(file string) {
	//加载配置
	GetConf(file)
	if MysqlConf.Enabled == true {

	}

}
func checkError(err error) {
	if err != nil {
		log.Print("检测到主机没有开启该服务，检查退出！")
		os.Exit(1)
	}
}

func checkup(hostaddr, port string) {
	address := hostaddr + ":" + port
	conn, err := net.Dial("tcp", address)
	checkError(err)
	conn.Close()
}

func Taskrun(proto string, tasknum int, hostaddr, port string) {
	host = hostaddr
	tasks := make(chan Workdist, taskload)
	wg.Add(tasknum)

	for gr := 1; gr <= tasknum; gr++ {
		switch {
		case proto == "ssh":
			checkup(host, port)
			go sshWorker(tasks)
		case proto == "telnet":
			checkup(host, port)
			go telnetWorker(tasks)
		case proto == "ftp":
			checkup(host, port)
			go ftpWorker(tasks)
		case proto == "mysql":
			checkup(host, port)
			go mysqlWorker(tasks)
		case proto == "smtp":
			checkup(host, port)
			go smtpWorker(tasks)
		case proto == "smb":
			checkup(host, port)
			go smbWorker(tasks)
		case proto == "mssql": // 1433
			checkup(host, port)
			go mssqlWorker(tasks)
		case proto == "plugins":
			checkup(host, port)
			go PostgresWorker(tasks)
		case proto == "hive":
			checkup(host, port)
			go HiveWorker(tasks)
		case proto == "redis":
			checkup(host, port)
			go RedisWorker(tasks)
		default:
			return
		}
	}
	intport, _ := strconv.Atoi(port)
	if MysqlConf.Enabled {
		for _, U := range userlist_sql {
			for _, P := range passlist_sql {

				task := Workdist{
					Username: U.Username,
					Password: P.Password,
					Port:     intport,
				}
				tasks <- task
			}
		}
	}

	if FileConf.Enabled {
		for _, U := range userlist_file {
			for _, P := range passlist_file {

				task := Workdist{
					Username: U.Username,
					Password: P.Password,
					Port:     intport,
				}
				tasks <- task
			}
		}
	}

	close(tasks)
	wg.Wait()
}

func sshWorker(tasks chan Workdist) {
	defer wg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		log.Print("检测ssh服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
		_, err := tool.SshConnect(task.Username, task.Password, host, task.Port)
		if err == nil {
			log.Print("检测到ssh服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}

}

func telnetWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret, _ := tool.Telnet_Creat(host, task.Username, task.Password, task.Port)
		if ret == true {
			log.Print("检测到telnet服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}

func ftpWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret := tool.LoginFtp(host, task.Username, task.Password, task.Port)
		if ret == "230" {
			log.Print("检测到ftp服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}

func mysqlWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret := tool.Loginmysql(host, task.Username, task.Password, task.Port)
		if ret == "ok" {
			log.Print("检测到mysql服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}
func mssqlWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret := tool.Loginmssql(host, task.Username, task.Password, task.Port)
		if ret == "true" {
			log.Print("检测到mysql服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}

func HiveWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret := tool.LoginHive(host, task.Username, task.Password, task.Port)
		if ret == "true" {
			log.Print("检测到Hive服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}

func PostgresWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		ret := tool.LoginPostgres(host, task.Username, task.Password, task.Port)
		if ret == "true" {
			log.Print("检测到Postgres服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}

//暂不支持ssl
func smtpWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		err := tool.Checksmtp(host, task.Username, task.Password, task.Port)
		if err == nil {
			log.Print("检测到smtp 服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}

//暂不支持ssl
func smbWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		log.Print("检测SMB服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
		Ret, err := tool.SmbConnect(task.Username, task.Password, host, task.Port)
		if err == nil && Ret == "true" {
			log.Print("检测到SMB服务弱口令：    用户名: ", task.Username, "   密码: ", task.Password)
			os.Exit(1)
		}
	}
}

func RedisWorker(tasks chan Workdist) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		log.Print("检测Redis服务弱口令：密码: ", task.Password)
		Ret := tool.RedisConnect(task.Password, host, task.Port)
		if Ret == "true" {
			log.Print("检测到Redis服务弱口令:"," 密码: ", task.Password)
			os.Exit(1)
		}
	}
}

