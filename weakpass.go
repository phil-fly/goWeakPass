package main

import (
	"flag"
	"./work"
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
}