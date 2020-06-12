package redislogin

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"goWeakPass/define"
	"os"
)


func LoginRedis(value interface{}) bool {
	switch value.(type) {
	case define.ServiceInfo :break
	default :
		fmt.Println("程序错误")
		os.Exit(-1)
	}
	config := value.(define.ServiceInfo)

	c, err := redis.Dial("tcp", config.Host+":"+config.Port)
	if err != nil {
		return false
	}
	err = c.Send("auth", config.PassWord)
	if err != nil {
		return false
	}
	define.Output(value)
	return true
}
