package main

import (
	"flag"
	"goWeakPass/work"
	"log"
	"os"
)

var proto = flag.String("proto", "", "Weak password detection protos (ssh/telnet)")
var hostaddr = flag.String("host", "", "Weak password detection hostaddr")
var port = flag.String("port", "", "Weak password detection port")
var tasknum = flag.Int("p", 1, "Weak password detection Number of threads")
var confpath = flag.String("conf", "conf.ini", "Weak password detection confpath")

func main() {
	flag.Parse()
	check()
	work.Taskinit(*confpath)
	work.Taskrun(*proto, *tasknum, *hostaddr, *port)
	log.Print("未检测到该服务弱口令，请再次检查确认或补充字典！")
}

func check() {
	if *proto == "" {
		log.Print("请使用 -proto 指定协议名")
		os.Exit(1)
	}
	if *hostaddr == "" {
		log.Print("请使用 -host 指定目标地址")
		os.Exit(1)
	}
	if *port == "" {
		log.Print("请使用 -port 指定目标端口")
		os.Exit(1)
	}

}
