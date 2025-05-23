package genlib

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// 多种文件存在检测方式，尝试多种方法确保准确检测

// FileExists 判断文件是否存在（主用）
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileExistsOpen 使用 os.Open 打开文件检测是否存在
func FileExistsOpen(path string) bool {
	f, err := os.Open(path)
	if err != nil {
		return false
	}
	defer f.Close()
	return true
}

// FileExistsLstat 使用 os.Lstat 检测文件是否存在
func FileExistsLstat(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}

// RouteInfo 用于描述单个路由注册信息
type RouteInfo struct {
	PackageName  string
	RegisterFunc string
}

// GenerateGenRoutesFile 生成或追加路由注册文件
func GenerateGenRoutesFile(routes []RouteInfo, moduleName string) error {
	const genRoutesFilePath = "routes/gen_routes.go"

	// 先检测文件是否存在，使用多种检测函数
	existsStat := FileExists(genRoutesFilePath)
	existsOpen := FileExistsOpen(genRoutesFilePath)
	existsLstat := FileExistsLstat(genRoutesFilePath)

	log.Printf("文件检测: Stat=%v, Open=%v, Lstat=%v\n", existsStat, existsOpen, existsLstat)

	content := buildRoutesContent(routes, moduleName)

	if !existsStat && !existsOpen && !existsLstat {
		// 文件不存在，创建并写入全部内容
		if err := os.WriteFile(genRoutesFilePath, []byte(content), 0644); err != nil {
			return fmt.Errorf("生成路由注册文件失败: %w", err)
		}
		log.Println("生成路由注册文件（新建）:", genRoutesFilePath)
		return nil
	}

	// 文件存在，尝试追加新路由，避免重复
	fileData, err := os.ReadFile(genRoutesFilePath)
	if err != nil {
		return fmt.Errorf("读取路由注册文件失败: %w", err)
	}
	existingContent := string(fileData)

	for _, r := range routes {
		registerLine := fmt.Sprintf("\tcontrollers.Register%sRoutes(router, utils.GormDB)", ToCamelCase(r.PackageName))
		if strings.Contains(existingContent, registerLine) {
			log.Printf("路由已存在，跳过追加: %s\n", registerLine)
			continue
		}
		// 插入到 // === GENERATED ROUTES END === 之前
		idx := strings.Index(existingContent, "// === GENERATED ROUTES END ===")
		if idx == -1 {
			// 如果标记没找到，直接追加到文件末尾
			existingContent += "\n" + registerLine + "\n"
		} else {
			existingContent = existingContent[:idx] + registerLine + "\n" + existingContent[idx:]
		}
	}

	// 写回文件
	if err := os.WriteFile(genRoutesFilePath, []byte(existingContent), 0644); err != nil {
		return fmt.Errorf("追加路由注册文件失败: %w", err)
	}
	log.Println("追加路由注册文件成功:", genRoutesFilePath)
	return nil
}

// buildRoutesContent 生成完整的路由文件内容
func buildRoutesContent(routes []RouteInfo, moduleName string) string {
	var importsBuilder strings.Builder
	var routesBuilder strings.Builder

	importsBuilder.WriteString("\t\"gin-go-test/app/controllers\"\n")
	importsBuilder.WriteString("\t\"gin-go-test/utils\"\n")

	for _, r := range routes {
		routesBuilder.WriteString(fmt.Sprintf("\tcontrollers.Register%sRoutes(router, utils.GormDB)\n", ToCamelCase(r.PackageName)))
	}

	template := `package routes

import (
	"github.com/gin-gonic/gin"
{{.Imports}}
)

// RegisterGeneratedRoutes 统一注册所有生成的路由
func RegisterGeneratedRoutes(router *gin.Engine) {
{{.Routes}}
// === GENERATED ROUTES END ===
}
`

	content := strings.ReplaceAll(template, "{{.Imports}}", importsBuilder.String())
	content = strings.ReplaceAll(content, "{{.Routes}}", routesBuilder.String())
	return content
}