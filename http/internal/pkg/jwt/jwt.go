package jwt

import (
	"errors"
)

// GenerateJWT 生成一个 JWT 令牌
func GenerateJWT(userID int64) (string, error) {
	return "", nil
}

// VerifyJWT 验证 JWT 令牌并返回用户ID
func VerifyJWT(tokenString string) (int64, error) {

	return 0, errors.New("invalid token")
}
