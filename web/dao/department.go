package dao

import (
	"github.com/lyq183/monibuca/v3/web/model"
	"github.com/lyq183/monibuca/v3/web/utils"
	"strconv"
)

func GetPaged(pageNo string) (*model.Page, error) {
	//将页码转换为int64类型
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)

	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from books"
	var totalRecord int64 //设置一个变量接收总记录数

	row := utils.Db.QueryRow(sqlStr) //执行
	row.Scan(&totalRecord)

	var pageSize int64 = 4 //设置每页显示的记录数目
	var totalPageNo int64  //记录总页数
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	//获取当前页中的
	sqlStr2 := "SELECT d_id,d_name,d_manager_id,d_description FROM department LIMIT ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var departments []*model.Department
	for rows.Next() {
		department := &model.Department{}
		rows.Scan(&department.D_id, &department.D_name, &department.D_manager, &department.D_description)
		departments = append(departments, department)
	}
	//创建page
	page := &model.Page{
		Departments: departments,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func Get_D_byname(d_name string) int { //	根据 部门名称 查 id
	sql := "SELECT d_id,d_name,d_manager_id,d_description FROM department where d_name = ?"
	rows := utils.Db.QueryRow(sql, d_name)
	department := &model.Department{}
	rows.Scan(&department.D_id, &department.D_name, &department.D_manager, &department.D_description)
	return department.D_id
}
func Get_D_byid(d_id int) string { //	根据 id 查 部门名称
	sql := "SELECT d_id,d_name,d_manager_id,d_description FROM department where d_id = ?"
	rows := utils.Db.QueryRow(sql, d_id)
	department := &model.Department{}
	rows.Scan(&department.D_id, &department.D_name, &department.D_manager, &department.D_description)
	return department.D_name
}

func Add_department(dname string, uname string, ss string) {
	user, _ := CheckUserName(uname)
	if user.Uid != 6 {
		sql := "INSERT INTO department(d_name,d_manager_id,d_description) VALUES (?,?,?)"
		utils.Db.Exec(sql, dname, user.Uid, ss)
	}
}
