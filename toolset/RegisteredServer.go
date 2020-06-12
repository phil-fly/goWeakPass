package toolset

import (
	"goWeakPass/toolset/ftp"
	"goWeakPass/toolset/hive"
	"goWeakPass/toolset/mangodb"
	"goWeakPass/toolset/mysql"
	"goWeakPass/toolset/postgres"
	"goWeakPass/toolset/redis"
	smblogin "goWeakPass/toolset/smb"
	"goWeakPass/toolset/smtp"
	sshLogin "goWeakPass/toolset/ssh"
	"goWeakPass/toolset/telnet"
)

func init(){
	Init()
	ManageServer.RegisteredServer("FTP",ftp.LoginFtp)
	ManageServer.RegisteredServer("HIVE",hive.LoginHive)
	ManageServer.RegisteredServer("MANGODB",mangodb.LoginMango)
	ManageServer.RegisteredServer("MYSQL",mysql.Loginmysql)
	ManageServer.RegisteredServer("POSTGRES",postgres.LoginPostgres)
	ManageServer.RegisteredServer("REDIS",redislogin.LoginRedis)
	ManageServer.RegisteredServer("SMB",smblogin.LoginSmb)
	ManageServer.RegisteredServer("SMTP",smtp.LoginSmtp)
	ManageServer.RegisteredServer("SSH",sshLogin.LoginSsh)
	ManageServer.RegisteredServer("TELNET",telnet.LoginTelnet)
}
