package controller

import (
	"fmt"
	"github.com/lyq183/monibuca/v3/web/dao"
	"net/http"
	"text/template"
)

//	管理所有单位
func Admin_department_index(w http.ResponseWriter, r *http.Request) {
	//获取页码
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPaged(pageNo)
	fmt.Println("|||||", page.Departments)

	t := template.Must(template.ParseFiles("web/views/pages/department/department_index.html"))
	t.Execute(w, page)
}

func Admin_projectManagement(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	str := "<script>	" +
		"var Main = {" +
		"data() { " +
		"return {" +
		"tableData: [{" +
		"d_name: '2016-05-02'," +
		"manager: '王小虎'," +
		"}]," +
		"search: ''" +
		"}" +
		"}," +
		"methods: {" +
		"handleEdit(index, row) {" +
		"console.log(index, row);}," +
		"handleDelete(index, row) {" +
		"console.log(index, row);" +
		"}" +
		"}," +
		"}" +

		"\nvar Ctor = Vue.extend(Main)" +
		"\nnew Ctor().$mount('#app')" +
		"</script>"

	t := template.Must(template.ParseFiles("web/views/pages/department/project_index.html"))
	t.Execute(w, str)
}

func Admin_userManagement(w http.ResponseWriter, r *http.Request) {
}
