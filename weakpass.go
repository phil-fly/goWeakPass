package main

import (
	"flag"
	"./work"
	"log"
)



var proto = flag.String("proto", "null", "Weak password detection protos (ssh/telnet)")
var hostaddr = flag.String("host", "null", "Weak password detection hostaddr")
var tasknum = flag.Int("p", 1, "Weak password detection Number of threads")
var confpath = flag.String("conf", "../config/conf.ini", "Weak password detection confpath")

func main() {
	//获取命令行参数
	//  main.exe -host 192.168.0.92 -proto ssh -p 50 -conf  confpath/conf.ini
	flag.Parse()

	work.Taskinit(*confpath)
	work.Taskrun(*proto,*tasknum,*hostaddr)
	log.Print("未检测到该服务弱口令，请再次检查确认或补充字典！")
}