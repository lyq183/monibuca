package dao

import (
	"github.com/lyq183/monibuca/v3/web/model"
	"github.com/lyq183/monibuca/v3/web/utils"
)

//CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	sqlStr := "select id,username,password,department_id from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.Uid, &user.Username, &user.Password, &user.Department_id)
	return user, nil
}

//CheckUserName 根据用户名和密码从数据库中查询一条记录
func CheckUserName(username string) (*model.User, error) {
	sqlStr := "select id,username,password,power from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.Uid, &user.Username, &user.Password, &user.Department_id)
	return user, nil
}

//AddUser 向数据库中插入用户信息
func AddUser(username string, password string, department_id int) error {
	//写sql语句
	sqlStr := "insert into users(username,password,department_id) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, department_id)
	if err != nil {
		return err
	}
	return nil
}
