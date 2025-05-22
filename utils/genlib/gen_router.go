package genlib

import (
	"fmt"
	"os"
	"strings"
)

const genRoutesFilePath = "routes/gen_routes.go"

const genRoutesTemplate = `package routes

import (
	"github.com/gin-gonic/gin"
{{.Imports}}
)

// RegisterGeneratedRoutes 统一注册所有生成的路由
func RegisterGeneratedRoutes(router *gin.Engine) {
{{.Routes}}
}
`

// RouteInfo 用于描述单个路由注册信息
type RouteInfo struct {
	PackageName string
	RegisterFunc string // 例如 RegisterMemberRoutes
}

// GenerateGenRoutesFile 生成或者更新 gen_routes.go 文件
func GenerateGenRoutesFile(routes []RouteInfo, moduleName string) error {
	
	// 准备 imports 和 routes 字符串
	var importsBuilder strings.Builder
	var routesBuilder strings.Builder

	for _, r := range routes {
		importsBuilder.WriteString(fmt.Sprintf("\t\"gin-go-test/app/controllers/%s\"\n", r.PackageName))
		routesBuilder.WriteString(fmt.Sprintf("\t%s.%s(router)\n", r.PackageName, r.RegisterFunc))
	}

	content := strings.ReplaceAll(genRoutesTemplate, "{{.Imports}}", importsBuilder.String())
	content = strings.ReplaceAll(content, "{{.Routes}}", routesBuilder.String())

	// 检查文件是否存在，备份或覆盖逻辑你可以自行加
	// 这里简单直接写文件
	if err := os.WriteFile(genRoutesFilePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("生成路由注册文件失败: %w", err)
	}
	return nil
}