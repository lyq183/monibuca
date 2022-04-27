package model

import (
	"fmt"
	"log"
	"math/rand"
)

//	记录用户登陆与否及其权限
type Session struct {
	Session_id  string
	Permissions int
	User_id     int
}

//	生成随机数
func CreateUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	u[8] = (u[8] | 0x40) & 0x7F
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}
