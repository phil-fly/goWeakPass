# goWeakPass
## 项目简介
    使用golang编写的	服务弱口令检测,可指定并发检测线程数量
#### 现阶段支持协议

|序号|支持协议|
|:---|:---:|
|1|ftp|
|2|telnet|
|3|ssh|
|4|mysql|
|5|smtp|
|6|smb|
|7|mssql|
|8|postgres|
|9|hive|
|10|redis|
|11|mangoDB|


## 使用方式

```
[root@localhost goWeakPass]# ./goWeakPass -h
Usage of ./goWeakPass:
  -conf string
    	Weak password detection confpath (default "conf.ini")
  -database string
    	Weak password database name (default "admin")
  -host string
    	Weak password detection hostaddr
  -p int
    	Weak password detection Number of threads (default 1)
  -port string
    	Weak password detection port
  -proto string
    	Weak password detection protos (ssh/telnet)
```
### 配置文件
    修改config目录下conf.ini文件内容，配置字典数据库信息，改配置文件为默认使用配置，使用时可通过 -conf 自定义配置文件路径
#### 命令
windows:(例) `weakpass.exe -host 10.10.10.111 -proto ssh -port 22 -p 50 -conf ../config/conf.ini`
    
linux:(例) `./weakpass -host 10.10.10.111 -proto ssh -port 22 -p 50 -conf ../config/conf.ini`
##### 参数说明：
    -h      查看帮助信息
    -host   指定检测主机地址
    -proto  指定检测服务协议
    -port   指定服务端口
    -p      指定并发检测线程数量
    -conf   指定配置文件路经
    -database 指定扫描的数据库名
    
## 字典数据库
### 账户字典：
```
DROP TABLE IF EXISTS `userdist`;
CREATE TABLE `userdist` (
  `id` int(10) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;
```
### 密码字典：
```
DROP TABLE IF EXISTS `passdist`;
CREATE TABLE `passdist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `password` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;
```
