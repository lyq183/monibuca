package dao

import (
	"github.com/lyq183/monibuca/v3/web/model"
	"github.com/lyq183/monibuca/v3/web/utils"
)

func AddSession(sess *model.Session) error {
	sqlStr := "insert into sessions values(?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sess.Session_id, sess.Permissions, sess.User_id)
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
