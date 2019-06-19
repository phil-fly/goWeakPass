package main

import (
	"flag"
	"goWeakPass/work"
	"log"
)

var proto = flag.String("proto", "null", "Weak password detection protos (ssh/telnet)")
var hostaddr = flag.String("host", "null", "Weak password detection hostaddr")
var port = flag.String("port", "null", "Weak password detection port")
var tasknum = flag.Int("p", 1, "Weak password detection Number of threads")
var confpath = flag.String("conf", "../config/conf.ini", "Weak password detection confpath")

func main() {
	flag.Parse()

	work.Taskinit(*confpath)
	work.Taskrun(*proto, *tasknum, *hostaddr, *port)
	log.Print("未检测到该服务弱口令，请再次检查确认或补充字典！")
}
