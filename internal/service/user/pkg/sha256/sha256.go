package sha256

import (
	"crypto/sha256"
	"encoding/hex"

	"main/config"
)

func Encrypt(password string) string {
	// 密码加盐
	hash := sha256.New()
	hash.Write([]byte(password + config.ConfAuth.Salt))
	return hex.EncodeToString(hash.Sum(nil))
}
