# goWeakPass
## 项目简介
    使用golang编写的	服务弱口令检测,可指定并发检测线程数量
#### 现阶段支持协议

|序号|支持协议|
|:---|:---:|
|1|ftp|
|2|telnet|
|3|ssh|
|4|Mysql|
|5|Smtp|
|6|Smb|
|7|Mssql|
|8|Postgres|


## 使用方式
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
