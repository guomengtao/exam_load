package controllers


import (
	"fmt"
	"net/http"
	"gin-go-test/app/models"
	"gin-go-test/utils"
	"golang.org/x/crypto/bcrypt"
	"github.com/gin-gonic/gin"
)

func GetAdminsHandler(c *gin.Context) {
    var admins []models.Admin
    err := utils.DBX.Select(&admins, "SELECT * FROM tm_admin")  // 直接查询管理员数据
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"message": "查询失败"})
        return
    }
    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "message": "Query successful",
        "data": admins,
    })
}

func UpdateAdminPasswordHandler(c *gin.Context) {
    var req struct {
        AdminID     int    `json:"admin_id"`
        NewPassword string `json:"new_password"`
    }
    // 绑定 JSON 参数
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": "请求参数错误"})
        return
    }

    // 调用 UpdateAdminPassword 函数
    err := UpdateAdminPassword(req.AdminID, req.NewPassword)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "密码更新成功"})
}

func UpdateAdminPassword(adminID int, newPassword string) error {
    var admin models.Admin
    err := utils.DBX.Get(&admin, "SELECT * FROM tm_admin WHERE id = ?", adminID)
    if err != nil {
        return fmt.Errorf("无法获取管理员信息: %v", err)
    }

    // 比对密码是否匹配
    err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(newPassword))
    if err == nil {
        return fmt.Errorf("新密码不能与旧密码相同")
    }

    // 加密新密码
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
    if err != nil {
        return fmt.Errorf("密码加密失败: %v", err)
    }

    // 更新数据库中的密码
    _, err = utils.DBX.Exec("UPDATE tm_admin SET password = ? WHERE id = ?", hashedPassword, adminID)
    if err != nil {
        return fmt.Errorf("更新密码失败: %v", err)
    }

    return nil
}