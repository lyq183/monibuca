package main

import (
	"fmt"
	"github.com/lyq183/monibuca/v3/configs"
	"github.com/lyq183/monibuca/v3/web/common"
	"github.com/lyq183/monibuca/v3/web/dao"
	"log"
	"net/http"
	//"html/template"
	"text/template"

	"github.com/lyq183/monibuca/v3/web/config"
	"github.com/lyq183/monibuca/v3/web/controller"
)

var (
	// 过滤器
	filter_user  = common.NewFilter() //普通用户权限函数
	filter_admin = common.NewFilter() //管理员权限函数
)

func Webindex() {
	stripPrefix()                          //	加载静态文件
	handlefuncAll()                        //	注册路由
	http.HandleFunc("/", controller.Index) //先登陆

	fmt.Println("登陆用户：http://localhost" + config.Ip)
	if err := http.ListenAndServe(config.Ip, nil); err != nil {
		log.Fatal("错误！！！ListenAndServe err:", err)
	}
}

func handlefuncAll() {
	userpower("/If_monibuca", controller.If_monibuca)             //	检查monibuca是否启动
	userpower("/monibuca", Monibuca_run)                          //	启动某个 monibuca
	userpower("/user_PasswordChange", controller.Password_Change) //	用户修改密码

	http.HandleFunc("/login", controller.Login)             //	登陆
	http.HandleFunc("/logout", controller.Logout)           //	登出
	http.HandleFunc("/monibuca_wu", controller.Monibuce_wu) //	monibuca未启动

	http.HandleFunc("/admin", controller.AdminLogin)                            //	管理员登陆
	adminpower("/department_manage", controller.Admin_department_index)         //	部门管理主界面
	adminpower("/add_department", controller.Add_department)                    //	添加新部门
	adminpower("/edit_department", controller.Edit_department)                  //	编辑部门信息
	adminpower("/Admin_portjectManagement", controller.Admin_projectManagement) //	项目管理
	adminpower("/Edit_config", controller.Edit_config)                          //	修改项目配置文件
	adminpower("/Admin_userManagement", controller.Admin_userManagement)        //	用户管理
	adminpower("/regist", controller.Regist)                                    //	注册
	adminpower("/regist_email", controller.Regist_email)                        //	发送邮箱验证码

	adminpower("/ffmpeg", controller.Ffmpeg)
	adminpower("/ffmpegPuth", controller.FfmpegPuth) //	ffmpeg推流
}

//	管理员权限 路由注册
func adminpower(route string, WebHandle func(w http.ResponseWriter, r *http.Request)) {
	filter_admin.RegisterFilterUri(route, WebHandle)                          // 注册拦截器
	http.HandleFunc(route, filter_admin.Admin_Handle(controller.Admin_Check)) // 启动服务
}

//	普通用户权限 路由注册
func userpower(route string, WebHandle func(w http.ResponseWriter, r *http.Request)) {
	filter_user.RegisterFilterUri(route, WebHandle)
	http.HandleFunc(route, filter_user.Handle(controller.Check))
}

func Monibuca_run(w http.ResponseWriter, r *http.Request) {
	config := r.PostFormValue("config")
	_, ok := configs.Monibucas[config]
	if !ok {
		configs.Monibucas[config] = true //	记录该monibuca已经启动
		go Monibuca(config)              //	启动 monibuca

		_, sess := common.IsLogin(r)
		//	根据用户名，查询名下的项目
		Pros, _ := dao.GetProjects_Byusername(sess.User_id)

		str_add := ""
		for _, v := range Pros {
			user_name, _ := dao.CheackUserId(v.P_u_id)
			ss := "{" +
				"d_name:'" + v.P_name + "'," +
				"manager: '" + user_name.Username + "'," +
				"config:'" + v.P_configName + "'," +
				"},"
			str_add += ss
		}
		fmt.Println(str_add)

		str_user := "<script>	" +
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
			"\nStartMonibuca(index, row) {" +
			"\nwindow.location.href=\"monibuca?config=\"" + "+row.config" +
			"}," +
			"\nRunMonibuca(index, row) {" +
			"\nwindow.location.href=\"If_monibuca?config=\"" + "+row.config" +
			"}" +
			"}," +
			"}" +

			"\nvar Ctor = Vue.extend(Main)" +
			"\nnew Ctor().$mount('#app')" +
			"</script>"

		t := template.Must(template.ParseFiles("web/views/pages/user/user.html"))
		t.Execute(w, str_user)
	}
}

//func zhuanyi(w http.ResponseWriter, r *http.Request) {
//	//	解析模板之前定义一个自定义函数 safe，保证传输的内容不会被安全化
//	t, _ := template.New("login.html").Funcs(template.FuncMap{
//		"safe": func(str string) template.HTML {
//			return template.HTML(str)
//		},
//	}).ParseFiles("js/login.html")
//
//	str1 := "<a href='http://bilibili.com'>b站</a>"
//	str2 := "<script>alert(123);</script>"
//	t.Execute(w, map[string]string{
//		"str1": str1,
//		"str2": str2,
//	})
//}

//func Monibuca_start(w http.ResponseWriter, r *http.Request) {
//	if !controller.Monibuca_flag {
//		controller.Monibuca_flag = true
//		fmt.Println("管理员启动monibuca引擎：")
//		go Monibuca("") //	启动 monibuca
//	}
//	t := template.Must(template.ParseFiles("web/views/pages/admin/administrator.html"))
//	t.Execute(w, map[string]string{
//		"ui":  "/ui/",
//		"str": "Had_monibuca",
//	})
//}

func stripPrefix() {
	//	http.FileServer(prefix string,h Handler) Handler：
	//		返回一个处理器，该处理器会将请求的URL.Path字 段中给定 前缀prefix 去除掉后再交由 h处理。
	//		stringPrefix会向URL.Path字段中没有给定前缀的请求回复404 page not found

	//	func FileServer(root FileSystem) Handler
	//		FileServer返回一个使用FileSystem接口root提供文件访问服务的HTTP处理器。
	//		要使用操作系统的FileSystem接口实现，可使用http.Dir：http.Handle("/", http.FileServer(http.Dir("/tmp")))

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("web/views/pages"))))
}
