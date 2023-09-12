package test

import (
	"fmt"
	"main/config"
	"main/server/utils/email"
	"testing"
)

func TestSend(t *testing.T) {
	fmt.Printf("config.Email: %v\n", config.ConfServer.Email)
	err := email.Send("注册验证码", "<h1>注册验证码</h1><p>123456</p>", "1993584108@qq.com")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestSendCaptcha(t *testing.T) {
	err := email.SendCaptcha("123456", "1993584108@qq.com")
	if err != nil {
		fmt.Println(err.Error())
	}
}
