package toolset

import (
	"sync"
	"reflect"
)

type Server struct {
	Service map[string]interface{}
	ServiceMutex  *sync.RWMutex

}

var ManageServer *Server

func Init(){
	if ManageServer == nil {
		//fmt.Println("honeypot ManageServer 初始化")
		ManageServer = &Server{
			Service:make(map[string]interface{}),
			ServiceMutex:new(sync.RWMutex),
		}
	}
}

func (this *Server)RegisteredServer(key string,value interface{}) {
	this.ServiceMutex.Lock()
	defer this.ServiceMutex.Unlock()
	//fmt.Println("发现服务注册:",key)
	this.Service[key] = value
}

func (this *Server)GetServer(key string) (interface{},bool){
	this.ServiceMutex.RLock()
	defer this.ServiceMutex.RUnlock()
	v,ok:=this.Service[key]
	return v,ok
}

func (this *Server)Call(server interface{}, params ... interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(server)
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}
