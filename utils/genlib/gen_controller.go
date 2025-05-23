package genlib

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// GenerateControllerWithAppend 路由追加版
func GenerateControllerWithAppend(tableName, moduleName string) error {
	tmplPath := "utils/gen/templates/controller.tpl"
	outputPath := fmt.Sprintf("app/controllers/%s_controller.go", strings.ToLower(tableName))

	controllerName := ToCamelCase(tableName)
	routePath := strings.ToLower(tableName)

	data := map[string]string{
		"ModuleName":     moduleName,
		"ControllerName": controllerName,
		"RoutePath":      routePath,
		"PackageName":    "controllers",
	}

	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		return fmt.Errorf("加载模板失败: %v", err)
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("创建文件失败: %v", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("渲染模板失败: %v", err)
	}

	fmt.Println("✅ 控制器及路由已生成到:", outputPath)
	fmt.Printf("🚀 路由访问示例: /api/%s/\n", routePath)
	return nil
}