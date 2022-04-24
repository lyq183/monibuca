package common

//	拦截器
import (
	"fmt"
	"net/http"
	"strings"
)

var Flag = false                                            //	标志拦截与否，即登陆与否
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
		if !Flag {
			// 执行拦截业务逻辑
			webHandle(w, r)
		} else {
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
