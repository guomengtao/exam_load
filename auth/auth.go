package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// 全局JWT密钥（生产环境应从环境变量读取）
const jwtSecret = "your-secret-key"

// LoginRequest 登录请求结构
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler 处理登录
func LoginHandler(c *gin.Context) {
	// 临时跳过验证，直接返回测试Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": 1,                  // 固定用户ID
		"role":     "admin",            // 固定角色
		"exp":      time.Now().Add(24 * time.Hour).Unix(), // 24小时有效期
	})
	tokenString, _ := token.SignedString([]byte(jwtSecret))
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

// Auth 认证中间件（重命名以避免冲突）
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供访问令牌"})
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("非法的令牌签名方法")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "无效的访问令牌"})
			return
		}

		c.Next()
	}
}