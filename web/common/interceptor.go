package common

//	拦截器
import (
	"fmt"
	"net/http"
	"strings"

	"github.com/lyq183/monibuca/v3/web/controller"
	"github.com/lyq183/monibuca/v3/web/dao"
	"github.com/lyq183/monibuca/v3/web/model"
)

type WebHandle func(w http.ResponseWriter, r *http.Request) // 声明新的函数类型

// 拦截器结构体
type Filter struct {
	filterMap map[string]WebHandle
}

//	创建拦截器
func NewFilter() *Filter {
	return &Filter{filterMap: make(map[string]WebHandle)}
}

// 注册拦截路由
func (f *Filter) RegisterFilterUri(uri string, handler WebHandle) {
	f.filterMap[uri] = handler
}

// 根据Uri获取相应的handle
func (f *Filter) GetFilterHandle(uri string) WebHandle {
	return f.filterMap[uri]
}

// 执行拦截器，返回函数类型
func (f *Filter) Handle(webHandle WebHandle) WebHandle {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("检测到请求：" + r.RequestURI)
		Flag, sess := IsLogin(r)

		if !Flag { //	没有用户登陆！
			fmt.Println("!!执行拦截:" + r.RequestURI)
			// 执行拦截业务逻辑
			webHandle(w, r)
		} else if sess.Permissions == 0 {
			controller.Not404(w, r)
		} else { //	有管理员权限，允许访问
			url := strings.Split(r.RequestURI, "?")
			for path, handle := range f.filterMap {
				if path == url[0] {
					handle(w, r)
					break
				}
			}
		}
	}
}

//	检查数据库，判断用户是否登陆
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user") //根据Cookie的name获取Cookie
	fmt.Println("!!!!", cookie)
	if cookie != nil { //	存在用户已经登陆
		cookieValue := cookie.Value //获取Cookie的value
		fmt.Println(cookieValue)
		//根据cookieValue 去数据库中查询与之对应的 Session
		if sess, _ := dao.GetSession(cookieValue); sess != nil {
			fmt.Println("用户已经登陆。")
			return true, sess
		}
	}
	return false, nil
}

// 执行拦截器，返回函数类型
func (f *Filter) Admin_Handle(webHandle WebHandle) WebHandle {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("检测到管理权限请求：" + r.RequestURI)
		Flag := Admin_IsLogin(r)

		if !Flag { //	非管理员
			fmt.Println("!!执行拦截:" + r.RequestURI)
			webHandle(w, r) // 执行拦截业务逻辑
		} else { //	有管理员权限，允许访问
			url := strings.Split(r.RequestURI, "?")
			for path, handle := range f.filterMap {
				if path == url[0] {
					handle(w, r)
					break
				}
			}
		}
	}
}

//	检查数据库，判断用户是否登陆
func Admin_IsLogin(r *http.Request) bool {
	cookie, _ := r.Cookie("admin") //根据 Cookie的name获取Cookie
	if cookie != nil {             //	存在管理员已经登陆,验证身份
		cookieValue := cookie.Value //获取 Cookie的value
		fmt.Println(cookieValue)
		//根据cookieValue 去数据库中查询 管理员的 Session
		if ok, _ := dao.Admin_GetSession(cookieValue); ok {
			return true
		}
	}
	return false
}
