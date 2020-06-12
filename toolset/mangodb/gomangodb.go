package mangodb

import (
	"goWeakPass/define"
	"gopkg.in/mgo.v2"
	"fmt"
	"os"
)


func LoginMango(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)

	session, err := mgo.Dial("mongodb://"+config.UserName+":"+config.PassWord+"@"+config.Host+":"+config.Port+"/"+config.DbName)
	defer session.Close()
	if err != nil {
		fmt.Println(err)
		//log.Print("认证失败")
		return false
	} else {
		//log.Print("认证成功")
		define.Output(value)
		return true
	}
}
