package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/lyq183/monibuca/v3/web/dao"
	"github.com/lyq183/monibuca/v3/web/utils"
)

//	管理员 添加用户
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	power := r.PostFormValue("power")

	//验证用户名和密码合法否
	user, _ := dao.CheckUserName(username)
	if user.Uid == 12 {
		t := template.Must(template.ParseFiles("web/views/pages/admin/regist.html"))
		t.Execute(w, "请输入注册的用户名和密码！")
	} else if user.Uid > 0 { //用户名已存在
		t := template.Must(template.ParseFiles("web/views/pages/admin/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else { //用户名和密码合法
		var power_int = 0
		if power == "A" {
			power_int = 1
		}

		if err := dao.AddUser(username, password, power_int); err != nil {
			fmt.Println("controller/Regist()：发生错误:", err)
		} else {
			t := template.Must(template.ParseFiles("web/views/pages/err/404.html"))
			t.Execute(w, "注册成功！")
		}
	}
}

var (
	emailMap = make(map[string]string)
)

func Regist_email(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	power := r.PostFormValue("power")
	vcode := r.PostFormValue("vcode")

	//验证用户名和密码合法否
	user, _ := dao.CheckUserName(email)
	if user.Uid == 12 {
		t := template.Must(template.ParseFiles("web/views/pages/admin/email_regist.html"))
		t.Execute(w, "请输入注册的用户名和密码")
	} else if user.Uid > 0 { //用户名已存在
		t := template.Must(template.ParseFiles("web/views/pages/admin/email_regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		if _, ok := emailMap[email]; !ok { //为发送邮件
			fmt.Println(ok)
			fmt.Println(emailMap)
			str := utils.Create_verificationCode()
			emailMap[email] = str
			fmt.Println(emailMap)
			utils.SendMail(email, str) //	向注册邮箱发送邮件
			t := template.Must(template.ParseFiles("web/views/pages/admin/email_regist.html"))
			t.Execute(w, "已发送验证码")
		} else {
			if vcode == emailMap[email] {
				var power_int = 0
				if power == "A" {
					power_int = 1
				}
				if err := dao.AddUser(email, password, power_int); err != nil {
					fmt.Println("controller/Regist()：发生错误:", err)
				} else {
					t := template.Must(template.ParseFiles("web/views/pages/err/404.html"))
					t.Execute(w, "注册成功！")
				}
			} else {
				t := template.Must(template.ParseFiles("web/views/pages/admin/email_regist.html"))
				t.Execute(w, "验证码不正确")
			}

		}
	}
}
