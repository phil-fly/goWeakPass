package tool

import (
	"gopkg.in/mgo.v2"
	"fmt"
	"strconv"
)


func LoginMango(host, username, password string, port int,database string) string{
	if database == "" {

	}
	session, err := mgo.Dial("mongodb://"+username+":"+password+"@"+host+":"+strconv.Itoa(port)+"/"+database)
	defer session.Close()
	if err != nil {
		fmt.Println(err)
		//log.Print("认证失败")
		return "Failed"
	} else {
		//log.Print("认证成功")
		return "Success"
	}
}
