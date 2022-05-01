package dao

import (
	"fmt"

	"github.com/lyq183/monibuca/v3/web/model"
	"github.com/lyq183/monibuca/v3/web/utils"
)

//	往数据库中添加session，如果该用户之前使用过的session还存在，则替换
func AddSession(sess *model.Session) error {
	IF_sess, _ := GetSession_from_id(sess.User_id)
	sqlStr := ""
	var err error
	if IF_sess.User_id > 0 { //	替换
		sqlStr = "update sessions set session_id = ? WHERE user_id = ?"
		_, err = utils.Db.Exec(sqlStr, sess.Session_id, sess.User_id)
	} else {
		sqlStr = "insert into sessions values(?,?,?)"
		_, err = utils.Db.Exec(sqlStr, sess.Session_id, sess.Permissions, sess.User_id)
	}
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession 删除数据库中的Session
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSession 根据session的 Id值从数据库中查询Session
func GetSession(sessID string) (*model.Session, error) {
	sqlStr := "select session_id,permissions,user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	sess := &model.Session{} //创建Session
	//扫描数据库中的字段值为Session的字段赋值
	row.Scan(&sess.Session_id, &sess.Permissions, &sess.User_id)
	return sess, nil
}

//GetSession 根据session的 Id值从数据库中查询Session
func GetSession_from_id(userid int) (*model.Session, error) {
	fmt.Println(userid)
	sqlStr := "select session_id,permissions,user_id from sessions where user_id = ?"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(userid)
	sess := &model.Session{} //创建 Session
	row.Scan(&sess.Session_id, &sess.Permissions, &sess.User_id)
	return sess, nil
}
