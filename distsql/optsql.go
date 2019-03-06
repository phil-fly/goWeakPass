package distsql

import (
	"github.com/go-xorm/xorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Userdist struct {
	Id       int       `xorm:"INT(20) 'id'"`
	Username string    `xorm:"VARCHAR(64) 'username'"`
}

type Passdist struct {
	Id       int       `xorm:"INT(20) 'id'"`
	Password string    `xorm:"VARCHAR(64) 'password'"`
}

//定义字典内存存储
var Userlist  []Userdist
var Passlist  []Passdist

//定义orm引擎
var x *xorm.Engine

//增
func Insert(name string) (int64, bool) {
	user := new(Userdist)
	user.Username = name
	affected, err := x.Insert(user)
	if err != nil {
		return affected, false
	}
	return affected, true
}



func Del(id int64) {
	user := new(Userdist)
	x.Id(id).Delete(user)
}

//改
func update(id int64, user *Userdist) bool {
	affected, err := x.ID(id).Update(user)
	if err != nil {
		log.Fatal("错误:", err)
	}
	if affected == 0 {
		return false
	}
	return true
}
//查
func getinfo(id int) *Userdist {
	user := &Userdist{Id: id}
	is, _ := x.Get(user)
	if !is {
		log.Fatal("搜索结果不存在!")
	}
	return user
}


//连接数据库加载字典
func Sqlinit(user,pass,host,dbname string) {
	var err error
	x, err = xorm.NewEngine("mysql", user+":"+pass+"@tcp("+host+":3306)/"+dbname+"?charset=utf8")
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}
	if err := x.Sync(new(Userdist)); err != nil {
		log.Fatal("数据表同步失败:", err)
	}

	Userlist = getUserList()
	Passlist = getPassList()
}

//遍历
func getUserList() (as []Userdist){
	x.Distinct("username").Find(&as)
	return
}

func getPassList() (as []Passdist){
	x.Distinct("password").Find(&as)
	return
}

