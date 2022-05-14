package controller

import (
	"fmt"
	"github.com/lyq183/monibuca/v3/web/dao"
	"net/http"
	"net/url"
	"strings"
	"text/template"
)

//	管理所有部门单位
func Admin_department_index(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPaged(pageNo)

	t := template.Must(template.ParseFiles("web/views/pages/department/department_index.html"))
	t.Execute(w, page)
}

func Edit_department(w http.ResponseWriter, r *http.Request) { //	修改部门信息

}

func Edit_config(w http.ResponseWriter, r *http.Request) { //	修改项目配置文件

}

//	管理所有项目
func Admin_projectManagement(w http.ResponseWriter, r *http.Request) {
	url, _ := url.QueryUnescape(r.RequestURI)
	pp := strings.Split(url, "=")
	d_id := dao.Get_D_byname(pp[1])
	d_name := dao.Get_D_byid(d_id)
	Pros, _ := dao.Getprojects(d_id) //	查询出所有属于该部门的项目

	str_add := ""
	for _, v := range Pros {
		user_name, _ := dao.CheackUserId(v.P_u_id)
		ss := "{" +
			"p_name:'" + v.P_name + "'," +
			"manager: '" + user_name.Username + "'," +
			"config:'" + v.P_configName + "'," +
			"},"
		str_add += ss
	}
	fmt.Println(str_add)

	str := "<script>	" +
		"var Main = {" +
		"data() { " +
		"return {" +
		"tableData: [" +
		str_add +
		"]," +
		"search: ''" +
		"}" +
		"}," +
		"\nmethods: {" +
		"\nEdit_project(index, row) {" +
		"\nwindow.location.href=\"Edit_project?d_name=\"" + "+row.d_name" +
		"}," +
		"\nStartMonibuca(index, row) {" +
		"\nwindow.location.href=\"monibuca?config=\"" + "+row.config" +
		"}" +
		"}," +
		"}" +

		"\nvar Ctor = Vue.extend(Main)" +
		"\nnew Ctor().$mount('#app')" +
		"</script>"

	t := template.Must(template.ParseFiles("web/views/pages/department/project_index.html"))
	t.Execute(w, map[string]string{
		"str":    str,
		"d_name": d_name,
	})
}

//	查询所以用户信息
func Admin_userManagement(w http.ResponseWriter, r *http.Request) {
	Users, _ := dao.GetUsers()
	str_add := ""
	for _, v := range Users {
		if v.Uid == 6 {
			continue
		}
		d_name := dao.Get_D_byid(v.Department_id)
		ss := "{" +
			"u_name:'" + v.Username + "'," +
			"u_department: '" + d_name + "'," +
			"},"
		str_add += ss
	}

	str := "<script>	" +
		"var Main = {" +
		"data() { " +
		"return {" +
		"tableData: [" +
		str_add +
		"]," +
		"search: ''" +
		"}" +
		"}," +
		"methods: {" +
		"handleEdit(index, row) {" +
		"console.log(index, row);}," +
		"handleDelete(index, row) {" +
		"console.log(index, row);}," +
		"add1(){" +
		"window.location.href='regist';}," +
		"add2(){" +
		"window.location.href='regist_email';}," +
		"}," +
		"}" +
		"\nvar Ctor = Vue.extend(Main)" +
		"\nnew Ctor().$mount('#app')" +
		"</script>"

	t := template.Must(template.ParseFiles("web/views/pages/department/user_index.html"))
	t.Execute(w, str)
}
