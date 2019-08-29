package distfile

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"io"
)

type Userdist struct {
	Username string
}

type Passdist struct {
	Password string
}

//定义字典内存存储
var Userlist []Userdist
var Passlist []Passdist

func read_userfile_line(userfile string) []Userdist {
	file, err := os.OpenFile(userfile, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open ",userfile," file error!", err)
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("User dist file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		entrypoint := Userdist{Username: line}
		Userlist = append(Userlist, entrypoint)
		if err != nil {
			if err == io.EOF {
				fmt.Println("User dist File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return nil
			}
		}
	}
	return Userlist
}

func read_passfile_line(Passfile string) []Passdist {
	file, err := os.OpenFile(Passfile, os.O_RDWR, 0666)
	if err != nil {
		fmt.Println("Open ",Passfile," file error!", err)
		return nil
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}

	var size = stat.Size()
	fmt.Println("Pass dist file size=", size)

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		entrypoint := Passdist{Password: line}
		Passlist = append(Passlist, entrypoint)
		if err != nil {
			if err == io.EOF {
				fmt.Println("User dist File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return nil
			}
		}
	}
	return Passlist
}


//读取文件加载字典
func FlieDist_Get(Userfile ,Passfile string) ([]Userdist, []Passdist) {
	Userlist = read_userfile_line(Userfile)
	Passlist = read_passfile_line(Passfile)
	return Userlist, Passlist
}


