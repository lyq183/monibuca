package config

//	配置信息
var (
	//登陆监听地址
	Ip = ":8081"

	//	连接数据库
	DatabaseRoot     = "root"
	DatabasePassword = "12345"
	Database         = "monibuca"

	//发邮箱地址,定义邮箱服务器连接信息，
	My_email   = "1519560741@qq.com"
	Email_pass = "xhcsdvwnggoxbahc" //如果是阿里邮箱 pass填密码，qq邮箱填授权码
	Smtp_host  = "smtp.qq.com"
	Smtp_port  = "465"
	Email_name = "GoSee"
)
