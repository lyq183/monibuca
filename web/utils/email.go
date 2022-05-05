package utils

import (
	"fmt"
	"github.com/lyq183/monibuca/v3/web/config"
	"math/rand"
	"regexp"
	"strconv"
	"time"

	"gopkg.in/gomail.v2"
)

//	邮箱注册功能实现

//	验证输入字符串是否为邮箱格式
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func SendMail(mailTo string, verification_code string) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := map[string]string{
		"name": config.Email_name,
		"user": config.My_email,
		"pass": config.Email_pass,
		"host": config.Smtp_host,
		"port": config.Smtp_port,
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int

	m := gomail.NewMessage()
	//这种方式可以添加别名为 mailConn["name"]， 也可以直接用<code>m.SetHeader("From",mailConn["user"])</code>
	m.SetHeader("From", mailConn["name"]+"<"+mailConn["user"]+">")
	m.SetHeader("To", mailTo)                 //发送给多个用户
	m.SetHeader("Subject", "go监控网站用户注册验证码")   //设置邮件主题
	m.SetBody("text/html", verification_code) //设置邮件正文

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}

func Create_verificationCode() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}
