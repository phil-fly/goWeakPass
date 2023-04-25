package rdp

import (
	"fmt"
	"goWeakPass/define"
	"os"
	"github.com/icodeface/grdp"
	"github.com/icodeface/grdp/glog"
)

func LoginRdp(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("参数错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)

	client := grdp.NewClient(config.Host+":"+config.Port, glog.NONE)
	err := client.Login(config.UserName, config.PassWord)
	if err != nil {
		return false
	} else {
		define.Output(value)
		return true
	}
}
