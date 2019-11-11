package tool

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)


func RedisConnect(password, host string, port int)  (string) {

	c, err := redis.Dial("tcp", host+fmt.Sprintf("%d",port))
	if err != nil {
		return "false"
	}
	err = c.Send("auth", password)
	if err != nil {
		return "false"
	}
	return "true"
}
