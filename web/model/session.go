package model

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

//	记录用户登陆与否及其权限
type Session struct {
	Session_id  string
	Permissions int
	User_id     int
}

////	生成随机数
//func CreateUUID() (uuid string) {
//	u := new([16]byte)
//	_, err := rand.Read(u[:])
//	if err != nil {
//		log.Fatalln("Cannot generate UUID", err)
//	}
//	u[8] = (u[8] | 0x40) & 0x7F
//	u[6] = (u[6] & 0xF) | (0x4 << 4)
//	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
//	return
//}

func CreateUUID() string {
	return UniqueId()
}

//生成 32 位 md5字串
func GetMd5String(s string) string {
	h := md5.New() //	返回一个新的使用 md5校验的hash.Hash接口。
	//fmt.Println(h)
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func UniqueId() string {
	b := make([]byte, 48)
	//	rand.Reader ：全局共享的密码用的强随机数生成器。
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return "生成Guid字串失败"
	}
	//	URLEncoding：标准base64编码字符集，用于URL和文件名。
	//	EncodeToString：返回将 b 编码后的字符串。
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}
