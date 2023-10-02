package sha256

import (
	"crypto/sha256"
	"encoding/hex"

	"main/services/user/config"
)

func Encrypt(password string) string {
	// 密码加盐
	hash := sha256.New()
	hash.Write([]byte(password + config.Config.Auth.Salt))
	return hex.EncodeToString(hash.Sum(nil))
}