package redis

import "fmt"

func GenerateCaptchaKey(email string) string {
	return fmt.Sprintf("captcha:%s", email)
}

func GenerateAuthKey(id int64) string {
	return fmt.Sprintf("auth:%d", id)
}
