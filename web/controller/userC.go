package controller

import (
	"fmt"
	"github.com/lyq183/monibuca/v3/web/common"
	"net/http"
	"text/template"

	"github.com/lyq183/monibuca/v3/web/dao"
)

//	默认界面，先登陆
func Index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
	t.Execute(w, "")
}

//	检测登陆与否
func Check(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
	t.Execute(w, "请先登陆！")
}

//	注销登陆
func Logout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
	common.Flag = false
	t.Execute(w, "请重新登陆！")
}

//	Login 处理用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("检测到请求：" + r.RequestURI)
	r.ParseForm()
	fmt.Println(r.PostForm)

	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("输入？:" + username + " " + password)

	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserNameAndPassword(username, password)
	fmt.Println("login()：uid:", user.Uid, "name:", user.Username, ";pass:", user.Password, ";power:", user.Power)
	fmt.Println()

	if user.Uid > 0 {
		//用户名和密码正确
		t := template.Must(template.ParseFiles("web/views/pages/user/administrator.html"))
		fmt.Println("管理用户登陆！")
		common.Flag = true //	标志拦截器暂停工作；
		t.Execute(w, "")
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
		fmt.Println("controller/login()：登陆失败！用户或密码错误！")
		t.Execute(w, "用户名或密码不正确！")
	}
}
