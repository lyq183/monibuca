package config

//	配置信息
var (
	//Config = config{"9011","root","12345","library"}

	Ip = ":8081" //登陆监听地址

	//	连接数据库
	DatabaseRoot     = "root"
	DatabasePassword = "12345"
	Database         = "library"
)

//type config struct {
//	Ip string	//登陆监听地址
//
//	DatabaseRoot string
//	DatabasePassword string
//	Database string
//}
