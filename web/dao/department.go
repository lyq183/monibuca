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

	//获取当前页中的图书
	sqlStr2 := "SELECT d_id,d_name,d_manager_id,d_description FROM department LIMIT ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}

	var departments []*model.Department
	for rows.Next() {
		department := &model.Department{}
		rows.Scan(&department.D_id, &department.D_name, &department.D_manager, &department.D_description)
		//将book添加到books中
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
