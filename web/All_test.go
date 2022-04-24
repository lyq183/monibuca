package web

import (
	"fmt"
	"github.com/lyq183/monibuca/v3/web/dao"
	"testing"
)

func TestUser(t *testing.T) {

}

func testLogin(t *testing.T) {
	user, _ := dao.CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}
func testRegist(t *testing.T) {
	user, _ := dao.CheckUserName("admin")
	fmt.Println("获取用户信息是：", user)
}
