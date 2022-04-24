package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/lyq183/monibuca/v3/web/dao"
)

//	管理员 添加普通用户
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	power := r.PostFormValue("power")

	//验证用户名和密码合法否
	user, _ := dao.CheckUserName(username)
	fmt.Println("Regist()：uid:", user.Uid, "name:", user.Username, "pass:", user.Password, "power", user.Power)

	if user.Uid == 12 {
		t := template.Must(template.ParseFiles("web/views/pages/user/regist.html"))
		t.Execute(w, "输入注册的用户名和密码！")
	} else if user.Uid > 0 { //用户名已存在
		t := template.Must(template.ParseFiles("web/views/pages/user/regist.html"))
		fmt.Println("controller/Regist()：用户名已存在！")
		t.Execute(w, "用户名已存在！")
	} else { //用户名和密码合法
		var power_int = 0
		if power == "A" {
			power_int = 1
		}

		if err := dao.AddUser(username, password, power_int); err != nil {
			fmt.Println("controller/Regist()：发生错误:", err)
		} else {
			t := template.Must(template.ParseFiles("web/views/pages/user/regist_success.html"))
			fmt.Println("controller/Regist()：注册成功！")
			t.Execute(w, "")
		}
	}
}
