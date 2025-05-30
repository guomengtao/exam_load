package utils

import (
	"errors"
	"fmt"
	"os"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

// JWT相关错误
var (
	ErrInvalidToken = errors.New("invalid token")
	ErrInvalidPassword = errors.New("invalid password")
)

// GetJWTInfo decodes a raw JWT string and returns its claims
func GetJWTInfo(tokenString string) (map[string]interface{}, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, errors.New("JWT secret not set in environment")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Convert to map[string]interface{}
		result := make(map[string]interface{})
		for k, v := range claims {
			result[k] = v
		}
		return result, nil
	}

	return nil, errors.New("invalid token")
}
 

// ExtractToken 提取 Authorization 头中的 Bearer Token
func ExtractToken(authHeader string) string {
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return authHeader
}

// ParseJWT 解析JWT token
func ParseJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte("your-secret-key"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
