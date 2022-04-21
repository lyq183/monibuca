package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/lyq183/monibuca/v3/web/dao"
)

//Login 处理用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserNameAndPassword(username, password)
	fmt.Println("login()：uid:", user.Uid, "name:", user.Username, "pass:", user.Password, "email", user.Email)
	fmt.Println()
	if user.Uid > 0 {
		//用户名和密码正确
		t := template.Must(template.ParseFiles("web/views/pages/user/login_success.html"))
		t.Execute(w, "")
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
		fmt.Println("controller/login()：登陆失败！用户或密码错误！")
		t.Execute(w, "用户名或密码不正确！")
	}
}

//Regist 处理用户的函注册数
func Regist(w http.ResponseWriter, r *http.Request) {
	fmt.Println("path:", r.URL.Path)
	t := template.Must(template.ParseFiles("web/views/pages/user/regist.html"))
	t.Execute(w, "")

	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	fmt.Println("Regist()：uid:", user.Uid, "name:", user.Username, "pass:", user.Password, "email", user.Email)
	fmt.Println()

	//if user.Uid == 0 {
	if user.Uid > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("web/views/pages/user/regist.html"))
		fmt.Println("controller/Regist()：用户名已存在！")
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可用，将用户信息保存到数据库中
		dao.SaveUser(username, password, email)
		//用户名和密码正确
		t := template.Must(template.ParseFiles("web/views/pages/user/regist_success.html"))
		fmt.Println("controller/Regist()：注册成功！")
		t.Execute(w, "")
		//}
	}
}
