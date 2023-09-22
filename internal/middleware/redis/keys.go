package redis

import "fmt"

func GenerateCaptchaKey(email string) string {
	return fmt.Sprintf("captcha:%s", email)
}

func GenerateAuthKey(id int64) string {
	return fmt.Sprintf("auth:%d", id)
}

func GenerateSubmitKey(id int64) string {
	return fmt.Sprintf("submit:%d", id)
}

func GenerateSubmitsKey() string {
	return "submits"
}

func GenerateContestUserKey(contestID, userID int64) string {
	return fmt.Sprintf("contest_user:%d_%d", contestID, userID)
}

func GenerateRankKey(contestID int64) string {
	return fmt.Sprintf("rank:%d", contestID)
}

func GenerateJudgeKey(id string) string {
	return fmt.Sprintf("judge:%s", id)
}
