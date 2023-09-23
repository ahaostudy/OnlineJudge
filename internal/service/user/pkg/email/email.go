package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"

	"main/config"
)

func Send(subject, html string, toEmails ...string) error {
	conf := config.ConfUser
	e := email.NewEmail()

	// 发送方的邮箱
	e.From = conf.Email.From
	// 接收方的邮箱
	e.To = toEmails
	// 主题
	e.Subject = subject
	// 邮件内容
	e.HTML = []byte(html)
	// 服务器配置
	auth := smtp.PlainAuth("", conf.Email.Email, conf.Email.Auth, conf.Email.Host)

	// 发送消息
	return e.Send(conf.Email.Addr, auth)
}

func SendCaptcha(captcha string, toEmails ...string) error {
	subject := "【OnlineJudge】邮箱验证"
	html := fmt.Sprintf(`<div style="text-align: center;">
		<h2 style="color: #333;">欢迎使用，你的验证码为：</h2>
		<h1 style="margin: 1.2em 0;">%s</h1>
		<p style="font-size: 12px; color: #666;">请在5分钟内完成验证，过期失效，请勿告知他人，以防个人信息泄露</p>
	</div>`, captcha)
	return Send(subject, html, toEmails...)
}
