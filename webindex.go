package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/lyq183/monibuca/v3/web/config"
	"github.com/lyq183/monibuca/v3/web/controller"
)

var (
	login_flag = false //	标志登陆与否
)

func Webindex() bool {
	stripPrefix() //	加载静态文件

	handlefuncAll()             //	注册路由
	http.HandleFunc("/", index) //先登陆

	fmt.Println("登陆用户：http://localhost" + config.Ip)
	if err := http.ListenAndServe(config.Ip, nil); err != nil {
		log.Fatal("错误！！！ListenAndServe err:", err)
	}

	return false
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

func handlefuncAll() {
	//http.HandleFunc("/main", IndexHandler)
	http.HandleFunc("/login", controller.Login)        //登陆
	http.HandleFunc("/regist.html", controller.Regist) //注册
}

func index(w http.ResponseWriter, r *http.Request) { //	默认界面，先登陆
	t := template.Must(template.ParseFiles("web/views/pages/user/login.html"))
	t.Execute(w, "")
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//	func (t *Template) ParseFiles(filenames ...string) (*Template, error)：
	//		解析filenames 指定的文件里面的模板定义并解析结果与t关联。
	//	func (t *Template) Must(t *template, err error) *Tmplate：
	//		包装返回(*Template,error)的函数/方法调用，会在err非nil使panic，一般用于变量初始化。
	//	func (t *Template) Execute(wr io.Writer, data interface{}) error
	//		Execute方法将解析好的模板应用到data上，并输出写入wr。
	//		如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行。

	t := template.Must(template.ParseFiles("web/views/index.html"))
	t.Execute(w, "")
}
