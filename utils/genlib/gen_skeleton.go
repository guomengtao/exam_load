package genlib

import (
	"fmt"
 	"os"
	"strings"
	"text/template"
)

// GenerateControllerWithAppend 路由追加版
func GenerateSkeletonWithAppend(tableName, moduleName string, overwrite bool) error {
	tmplPath := "utils/gen/templates/skeleton.tpl"
	// Ensure directory exists
	if err := os.MkdirAll("utils/generated/controller", os.ModePerm); err != nil {
		return fmt.Errorf("创建目录失败: %v", err)
	}
	outputPath := fmt.Sprintf("utils/generated/controller/%s_skeleton.go", strings.ToLower(tableName))

	controllerName := ToCamelCase(tableName)
	routePath := strings.ToLower(tableName)

	data := map[string]string{
		"ModuleName":     moduleName,
		"ControllerName": controllerName,
		"RoutePath":      routePath,
		"PackageName":    "skeleton",
		"HelloMessage":   "hello123",
	}

	// 判断文件是否存在
	if _, err := os.Stat(outputPath); err == nil {
		if !overwrite {
			fmt.Println("⚠️ 骨架文件已存在，跳过生成。")
			return nil
		}
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

	fmt.Println("✅ 骨架已生成到:", outputPath)
	return nil
}