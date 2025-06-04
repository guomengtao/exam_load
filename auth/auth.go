package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gin-go-test/app/services"
	"gin-go-test/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

const jwtSecret = "your-secret-key"

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler 处理登录
func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "参数错误",
			"data":    nil,
		})
		return
	}

	cacheKey := "admin:" + req.Username
	adminData, err := utils.RedisClient.HGetAll(context.Background(), cacheKey).Result()

	// Fallback to DB if Redis miss or empty
	if err != nil || len(adminData) == 0 {
		adminService := services.NewAdminService(utils.GormDB)
		admin, err := adminService.GetAdminByUsername(req.Username)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "用户不存在",
				"data":    nil,
			})
			return
		}

		// Save to Redis
		roleIDStr := strconv.Itoa(admin.RoleId)
		utils.RedisClient.HSet(context.Background(), cacheKey, map[string]interface{}{
			"id":       admin.Id,
			"username": admin.Username,
			"password": admin.Password,
			"role_id":  admin.RoleId,
		})
		utils.RedisClient.Expire(context.Background(), cacheKey, 24*time.Hour)

		adminData = map[string]string{
			"id":       strconv.Itoa(admin.Id),
			"username": admin.Username,
			"password": admin.Password,
			"role_id":  roleIDStr,
		}
	}

	// Verify password
	if !utils.CheckPassword(req.Password, adminData["password"]) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "密码错误",
			"data":    nil,
		})
		return
	}

	// Generate JWT
	adminID, _ := strconv.Atoi(adminData["id"])
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"admin_id": adminID,
		"role_id":  adminData["role_id"],
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "生成令牌失败",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "登录成功",
		"data": gin.H{
			"token": tokenString,
		},
	})
}

// Auth 中间件
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "未提供访问令牌",
				"data":    nil,
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
				"code":    401,
				"message": "无效的访问令牌",
				"data":    nil,
			})
			return
		}

		c.Next()
	}
}

func GetAdminFromCache(adminID string) (*AdminInfo, error) {
	cacheKey := fmt.Sprintf("admin:info:%s", adminID)

	// Try to get from Redis
	adminData, err := utils.RedisHGetAll(cacheKey)
	if err != nil {
		return nil, fmt.Errorf("获取管理员缓存失败: %v", err)
	}

	if len(adminData) > 0 {
		// Parse from Redis
		admin := &AdminInfo{}
		if err := json.Unmarshal([]byte(adminData["data"]), admin); err == nil {
			return admin, nil
		}
	}

	// If not in Redis, get from database
	admin, err := GetAdminFromDB(adminID)
	if err != nil {
		return nil, err
	}

	// Save to Redis
	adminJSON, err := json.Marshal(admin)
	if err == nil {
		utils.RedisHSet(cacheKey, "data", string(adminJSON))
		utils.RedisExpire(cacheKey, 24*time.Hour)
	}

	return admin, nil
}

// AdminInfo 表示管理员信息
type AdminInfo struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// GetAdminFromDB 从数据库获取管理员信息
func GetAdminFromDB(adminID string) (*AdminInfo, error) {
	// 这里实现从数据库获取管理员信息的逻辑
	return &AdminInfo{
		ID:       1,
		Username: "admin",           // 实际项目中应使用真实用户名
		Password: "hashed_password", // 实际项目中应使用加密后的密码
	}, nil
}
