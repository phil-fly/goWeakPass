package distsql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
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
var db *sql.DB

//查
func getUserdist(dbtable string) []Userdist {
	var node Userdist
	rows, err := db.Query("select username  from ?  where 1",dbtable);
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		//获取请求的方法
		rows.Scan(&node.Username);
		entrypoint := Userdist{Username: node.Username}
		Userlist = append(Userlist, entrypoint)
	}
	return Userlist
}

//查
func getPassdist(dbtable string) []Passdist {
	var node Passdist
	rows, err := db.Query("select password  from ?  where 1",dbtable);
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		//获取请求的方法d
		rows.Scan(&node.Password);
		entrypoint := Userdist{Username: node.Password}
		Userlist = append(Userlist, entrypoint)
	}
	return Passlist
}


//连接数据库加载字典
func Sqlinit(user,pass,host,Dbport,dbname,userdist,passdist string) {
	db, err := sql.Open("mysql",
		user+":"+ pass +"@tcp("+ host +":"+Dbport+")/"+dbname+"?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()


	Userlist = getUserdist(userdist)
	Passlist = getPassdist(passdist)
}
