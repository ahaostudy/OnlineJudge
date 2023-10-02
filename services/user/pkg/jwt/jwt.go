package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"

	"main/config"
)

// Claim Token结构声明
type Claims struct {
	ID int64
	jwt.RegisteredClaims
}

// GenerateToken 生成Token
func GenerateToken(id int64) (string, error) {
	// 创建Claims对象
	claims := Claims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(config.ConfAuth.Expire) * time.Hour)),
			Issuer:    config.ConfAuth.Issuer,
		},
	}
	// 获取加密token
	return jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.ConfAuth.Key))
}

// ParseToken 解析Token
func ParseToken(token string) (int64, bool) {
	claims := new(Claims)
	t, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.ConfAuth.Key), nil
	})
	if !t.Valid || err != nil || claims == nil {
		return 0, false
	}
	return claims.ID, true
}