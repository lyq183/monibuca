package dao

import (
	"github.com/lyq183/monibuca/v3/web/model"
	"github.com/lyq183/monibuca/v3/web/utils"
)

//验证管理员登陆
func CheckAdmin(adminname string, password string) (*model.Admin, error) {
	sqlStr := "select id, adminname, password, session_id from admin where adminname = ? and password = ?"
	row := utils.Db.QueryRow(sqlStr, adminname, password)
	admin := &model.Admin{}
	row.Scan(&admin.Id, &admin.Adminname, &admin.Password, &admin.Session)
	return admin, nil
}

//	修改 管理员的 session
func Admin_ChangeSession(session string) error {
	var err error
	sqlStr := "update admin set session_id = ? WHERE id = 1"
	_, err = utils.Db.Exec(sqlStr, session)
	if err != nil {
		return err
	}
	return nil
}

//查询 管理员的 session
func Admin_GetSession(sessID string) (bool, error) {
	sqlStr := "select session_id from admin where id = 1"
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return false, err
	}
	//执行
	row := inStmt.QueryRow()
	sess := ""
	row.Scan(&sess)
	if sess == sessID {
		return true, nil
	}
	return false, nil
}
