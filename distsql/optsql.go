package distsql

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"database/sql"
	"fmt"
)

type Userdist struct {
	Id       int
	Username string
}

type Passdist struct {
	Id       int
	Password string
}

//定义字典内存存储
var Userlist  []Userdist
var Passlist  []Passdist

//定义orm引擎
var db *sql.DB

//查
func getUserdist(db *sql.DB,dbtable string) []Userdist {
	var node Userdist
	sqlqyery := fmt.Sprintf("select username  from  %s  where 1", dbtable)
	rows, err := db.Query(sqlqyery);
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
func getPassdist(db *sql.DB,dbtable string) []Passdist {
	var node Passdist
	sqlqyery := fmt.Sprintf("select password  from  %s  where 1", dbtable)
	rows, err := db.Query(sqlqyery);
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		//获取请求的方法d
		rows.Scan(&node.Password);
		entrypoint := Passdist{Password: node.Password}
		Passlist = append(Passlist, entrypoint)
	}
	return Passlist
}


//连接数据库加载字典
func DistGet(user,pass,host,Dbport,dbname,userdist,passdist string) ([]Userdist,[]Passdist){
	db, err := sql.Open("mysql",
		user+":"+ pass +"@tcp("+ host +":"+Dbport+")/"+dbname+"?charset=utf8")
	if err != nil {
		log.Fatal(err)
		return nil,nil
	}
	defer db.Close()


	Userlist = getUserdist(db,userdist)
	Passlist = getPassdist(db,passdist)
	return Userlist,Passlist
}
