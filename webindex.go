package main

import (
	"fmt"
	"github.com/lyq183/monibuca/v3/web/common"
	"html/template"
	"log"
	"net/http"

	"github.com/lyq183/monibuca/v3/web/config"
	"github.com/lyq183/monibuca/v3/web/controller"
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
	// 1.过滤器
	filter_all := common.NewFilter()
	// 注册拦截器
	filter_all.RegisterFilterUri("/monibuca", Monibuca_start)              //	启动 monibuca
	filter_all.RegisterFilterUri("/regist", controller.Regist)             //	注册
	filter_all.RegisterFilterUri("/ffmpeg", controller.Ffmpeg)             //	ffmpeg
	filter_all.RegisterFilterUri("/ffmpegPuth", controller.FfmpegPuth)     //	ffmpeg推流
	filter_all.RegisterFilterUri("/regist_email", controller.Regist_email) //	发送邮箱验证码
	// 2.启动服务
	http.HandleFunc("/regist", filter_all.Handle(controller.Check))
	http.HandleFunc("/monibuca", filter_all.Handle(controller.Check))
	http.HandleFunc("/ffmpeg", filter_all.Handle(controller.Check))
	http.HandleFunc("/ffmpegPuth", filter_all.Handle(controller.Check))
	http.HandleFunc("/regist_email", filter_all.Handle(controller.Check))

	http.HandleFunc("/login", controller.Login)   //	登陆
	http.HandleFunc("/logout", controller.Logout) //	登出
	http.HandleFunc("/monibuca_wu", controller.Monibuce_wu)
}

func Monibuca_start(w http.ResponseWriter, r *http.Request) {
	if !controller.Monibuca_flag {
		controller.Monibuca_flag = true
		fmt.Println("管理员启动monibuca引擎：")
		go Monibuca() //	启动 monibuca

	}
	t := template.Must(template.ParseFiles("web/views/pages/admin/administrator.html"))
	t.Execute(w, map[string]string{
		"ui":  "/ui/",
		"str": "Had_monibuca",
	})
}

func zhuanyi(w http.ResponseWriter, r *http.Request) {
	//	解析模板之前定义一个自定义函数 safe，保证传输的内容不会被安全化
	t, _ := template.New("login.html").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("js/login.html")

	str1 := "<a href='http://bilibili.com'>b站</a>"
	str2 := "<script>alert(123);</script>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

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
