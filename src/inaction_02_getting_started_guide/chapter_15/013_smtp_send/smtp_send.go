package main

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	// 邮件服务器地址
	SMTP_MAIL_HOST  = "smtp.163.com"

	// 端口
	SMTP_MAIL_PORT  = "25"

	// 拼接 服务器和端口
	SMTP_MAIL_ADDR = "smtp.163.com:25"

	// 发送邮件用户账号
	SMTP_MAIL_USERNAME  = "yy7546389@163.com"

	// 授权密码
	SMTP_MAIL_PASS   = "xxxxx"		//  这块得去授权 发送方和接收方的设置!
)

func main() {
	dests := []string{"2017585616@qq.com"}
	DoSend(dests)
}


func DoSend(dests []string) {
	// Set up authentication information. 通常身份应该是空字符串, 填充用户名.
	auth := smtp.PlainAuth(
		"",
		SMTP_MAIL_USERNAME,
		SMTP_MAIL_PASS,
		SMTP_MAIL_HOST,
	)

	for _, dest := range dests {
		// Connect to the server, authenticate, set the sender and recipient,
		// and send the email all in one step.
		err := smtp.SendMail(
			SMTP_MAIL_ADDR,
			auth,
			SMTP_MAIL_USERNAME,
			[]string{dest},
			[]byte("this is the email body."),
		)
		if err != nil {
			fmt.Printf("send error = %v \n", err)
			log.Fatal(err)
		}
	}

}
