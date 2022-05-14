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
	sqlStr := "select id,username,password,department_id from users where username = ?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.Uid, &user.Username, &user.Password, &user.Department_id)
	return user, nil
}

//	根据 id 查 name
func CheackUserId(uid int) (*model.User, error) {
	sqlStr := "SELECT id,username,PASSWORD,department_id FROM users WHERE id = ?"
	row := utils.Db.QueryRow(sqlStr, uid)
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

//	修改密码
func Password_Change(uid int, new_pw string) error {
	sql := "UPDATE users SET PASSWORD = ? WHERE id = ?"
	if _, err := utils.Db.Exec(sql, new_pw, uid); err != nil {
		return err
	}
	return nil
}

//	查询所有用户信息
func GetUsers() ([]*model.User, error) {
	sql := "SELECT id,username,PASSWORD,department_id FROM users"
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var pros []*model.User
	for rows.Next() {
		u := &model.User{}
		rows.Scan(&u.Uid, &u.Username, &u.Password, &u.Department_id)
		pros = append(pros, u)
	}
	return pros, nil
}
