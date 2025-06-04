package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// var jwtSecret = "your-secret-key" // 确保你在真实项目中用安全方式设置

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.TrimPrefix(c.GetHeader("Authorization"), "Bearer ")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "未提供访问令牌",
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("非法的令牌签名方法")
			}
			return []byte(jwtSecret), nil
		})
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    http.StatusUnauthorized,
				"message": "无效的访问令牌",
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Printf("Token claims: %+v\n", claims)
			c.Set("adminID", claims["admin_id"])
			c.Set("role", claims["role_id"])
		}
		c.Next()
	}
}

func PermissionMiddleware(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roleID, exists := c.Get("role")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "权限不足，未找到角色信息",
			})
			return
		}

		// 将 roleID 转换为字符串
		roleIDStr := fmt.Sprintf("%v", roleID)

		// 如果是管理员角色（role_id = "1"），直接放行
		if roleIDStr == "1" {
			c.Next()
			return
		}

		// 其他角色需要根据具体权限判断
		if requiredPermission == "exam:manage" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "需要管理员权限",
			})
			return
		}

		c.Next()
	}
}
