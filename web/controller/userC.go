package controller

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/lyq183/monibuca/v3/web/dao"
	"github.com/lyq183/monibuca/v3/web/model"
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

//	没有访问权限
func Not404(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("web/views/pages/error/404.html"))
	t.Execute(w, "没有访问权限")
}

//	注销登陆
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	re_str := ""
	if cookie != nil {
		cookieValue := cookie.Value    //获取 cookie的 value值
		dao.DeleteSession(cookieValue) //删除数据库中与之对应的Session
		cookie.MaxAge = -1             //设置cookie失效
		http.SetCookie(w, cookie)      //将修改之后的cookie发送给浏览器
		re_str = "请重新登陆！"
	} else {
		re_str = "尚未登陆！"
	}

	t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
	t.Execute(w, re_str)
}

//	Login 处理用户登录
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("检测到请求：" + r.RequestURI)
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.Uid > 0 { //用户名和密码正确
		//	第一次向服务器发送请求是创建 session，给它一个设置唯一的 ID(可通过UUID生成)
		str := model.CreateUUID()
		sess := &model.Session{ //创建一个Session
			Session_id: str,
			User_id:    user.Uid,
		}
		dao.AddSession(sess)   //将Session保存到数据库中
		cookie := http.Cookie{ //	创建一个 Cookie，
			Name:     "user",
			Value:    str, //	将其 Value值设置为 Seesion的id
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie) //将 cookie发送给浏览器

		ui := "/ui/" //	检查 monibuca是否启动
		if !Monibuca_flag {
			ui = "/monibuca_wu"
		}

		t := template.Must(template.ParseFiles("web/views/pages/user/user.html"))
		t.Execute(w, ui)
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}
