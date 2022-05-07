package dao

import (
	"github.com/lyq183/monibuca/v3/web/model"
	"github.com/lyq183/monibuca/v3/web/utils"
)

func Getprojects(d_id int) ([]*model.Project, error) {
	sql := "SELECT p_id,p_name,p_uid,P_configtName FROM project where p_department = ?"
	rows, err := utils.Db.Query(sql, d_id)
	if err != nil {
		return nil, err
	}

	var pros []*model.Project
	for rows.Next() {
		p := &model.Project{}
		rows.Scan(&p.P_id, &p.P_name, &p.P_u_id, &p.P_configName)
		pros = append(pros, p)
	}
	return pros, nil
}

func GetProjects_Byusername(uid int) ([]*model.Project, error) {
	sql := "SELECT p_id,p_name,p_uid,P_configtName FROM project where p_uid = ?"
	rows, err := utils.Db.Query(sql, uid)
	if err != nil {
		return nil, err
	}

	var pros []*model.Project
	for rows.Next() {
		p := &model.Project{}
		rows.Scan(&p.P_id, &p.P_name, &p.P_u_id, &p.P_configName)
		pros = append(pros, p)
	}
	return pros, nil
}
