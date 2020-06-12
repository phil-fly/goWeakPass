package work

import (
	"fmt"
	"goWeakPass/define"
	"goWeakPass/toolset"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

type Workdist struct {
	Username string
	Password string
	Port     int
	Database string
	Server	 string
}

const (
	taskload = 100
)

var wg sync.WaitGroup
var host string

func Taskinit(file string) {
	//加载配置
	GetConf(file)

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

func Taskrun(proto string, tasknum int, hostaddr, port ,database string) {
	Server,ok := toolset.ManageServer.GetServer(strings.ToUpper(proto))
	if !ok {
		fmt.Println("服务不存在!")
		os.Exit(1)
	}
	checkup(host, port)

	host = hostaddr
	tasks := make(chan define.ServiceInfo, taskload)
	wg.Add(tasknum)

	for gr := 1; gr <= tasknum; gr++ {
		go ServerWorker(tasks,Server)
	}
	intport, _ := strconv.Atoi(port)
	if MysqlConf.Enabled {
		for _, U := range userlist_sql {
			for _, P := range passlist_sql {

				task := define.ServiceInfo{
					Host: hostaddr,
					Port:     port,
					PortInt: intport,
					UserName: U.Username,
					PassWord: P.Password,
					DbName: database,
				}
				tasks <- task
			}
		}
	}

	if FileConf.Enabled {
		for _, U := range userlist_file {
			for _, P := range passlist_file {
				task := define.ServiceInfo{
					Host: hostaddr,
					Port:     port,
					PortInt: intport,
					UserName: U.Username,
					PassWord: P.Password,
					DbName: database,
				}
				tasks <- task
			}
		}
	}

	close(tasks)
	wg.Wait()
}

func ServerWorker(tasks chan define.ServiceInfo,Server interface{}){
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			//log.Print("通道关闭")
			return
		}
		toolset.ManageServer.Call(Server,task)
	}
}

