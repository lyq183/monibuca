package utils

//	连接数据库
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lyq183/monibuca/v3/web/config"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	//	func Open(driverName, dataSourceName string) (*DB, error)
	//		Open打开一个dirverName指定的数据库，dataSourceName指定数据源，一般包至少括数据库文件名和（可能的）连接信息
	//		大多数用户会通过数据库特定的连接帮助函数打开数据库，返回一个*DB。	Go标准库中没有数据库驱动。
	//		格式：sql.Open("mysql", "用户名:密码@tcp(ip:端口)/数据库?charset=utf8")

	root := config.DatabaseRoot
	password := config.DatabasePassword
	database := config.Database
	dataSourcename := root + ":" + password + "@tcp(localhost:3306)/" + database

	//dataSourcename := "root:12345@tcp(localhost:3306)/library"
	db, err := sql.Open("mysql", dataSourcename)
	if err != nil {
		fmt.Println("错误！！！sql.Open err:", err)
		log.Fatal("错误！！！sql.Open err:", err)
	} else {
		Db = db
	}
}
