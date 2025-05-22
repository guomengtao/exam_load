package genlib

import (
	"fmt"
	"io/ioutil"
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

	// 判断文件是否存在
	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		// 不存在，生成新文件
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

	// 文件存在，追加路由
	contentBytes, err := ioutil.ReadFile(outputPath)
	if err != nil {
		return fmt.Errorf("读取文件失败: %v", err)
	}
	content := string(contentBytes)

	funcName := fmt.Sprintf("Register%sRoutes", controllerName)
	startIdx := strings.Index(content, funcName)
	if startIdx == -1 {
		return fmt.Errorf("找不到路由注册函数 %s", funcName)
	}

	// 找函数体开始的大括号
	bodyStart := strings.Index(content[startIdx:], "{")
	if bodyStart == -1 {
		return fmt.Errorf("路由注册函数 %s 格式异常", funcName)
	}
	bodyStart += startIdx + 1

	// 找函数体结束的大括号
	bodyEnd := strings.Index(content[bodyStart:], "}")
	if bodyEnd == -1 {
		return fmt.Errorf("路由注册函数 %s 格式异常", funcName)
	}
	bodyEnd += bodyStart

	funcBody := content[bodyStart:bodyEnd]

	newRouteCode := fmt.Sprintf("\n\tgroup.GET(\"/%s\", ctrl.HelloHandler)", routePath)

	if strings.Contains(funcBody, newRouteCode) {
		fmt.Println("⚠️ 路由已存在，跳过追加。")
		return nil
	}

	newFuncBody := funcBody + newRouteCode + "\n"

	newContent := content[:bodyStart] + newFuncBody + content[bodyEnd:]

	if err := ioutil.WriteFile(outputPath, []byte(newContent), 0644); err != nil {
		return fmt.Errorf("写回文件失败: %v", err)
	}

	fmt.Println("✅ 路由已追加到:", outputPath)
	fmt.Printf("🚀 路由访问示例: /api/%s/\n", routePath)
	return nil
}