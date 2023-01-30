package main

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	// 邮件服务器地址
	SMTP_MAIL_HOST  = "smtp.126.com"
	// 端口
	SMTP_MAIL_PORT  = "25"
	// 发送邮件用户账号
	SMTP_MAIL_USER  = "yy7546389@126.com"
	// 授权密码
	SMTP_MAIL_PWD   = "Yangjianwei01"
	// 发送邮件昵称
	SMTP_MAIL_NICKNAME  = "SMTPMail"
)

func main() {
	to_addresses := []string{"yy7546389@gmail.com"}
	SendMail(to_addresses)
}


func SendMail(to_addresses []string) {
	// 通常身份应该是空字符串, 填充用户名.
	auth := smtp.PlainAuth("", SMTP_MAIL_USER, SMTP_MAIL_PWD, SMTP_MAIL_HOST)
	for _, to_address := range to_addresses {
		addr := fmt.Sprintf("%v:%v", SMTP_MAIL_HOST, SMTP_MAIL_PORT)
		msg := []byte("This is the email body.")
		err := smtp.SendMail(addr, auth, SMTP_MAIL_USER, []string{to_address}, msg)
		if err != nil {
			log.Fatal(err)
		}
	}
}
