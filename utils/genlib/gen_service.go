

package genlib

import (
	"fmt"
	"os"
 	"text/template"
)

type ServiceTemplateData struct {
	ServiceName string
}

 

func GenerateServiceFromTable(tableName string) error {
	serviceName := toCamelCase(tableName)

	data := ServiceTemplateData{
		ServiceName: serviceName,
	}

	tmpl, err := template.ParseFiles("utils/gen/templates/service.tpl")
	if err != nil {
		return fmt.Errorf("加载模板失败: %v", err)
	}

	outputPath := fmt.Sprintf("app/services/%s_service.go", tableName)
	if err := os.MkdirAll("app/services", os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("创建服务文件失败: %v", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("渲染模板失败: %v", err)
	}

	fmt.Println("✅ 服务生成成功:", outputPath)
	return nil
}

func GenerateServiceSkeleton(tableName string) error {
	serviceName := toCamelCase(tableName)

	data := ServiceTemplateData{
		ServiceName: serviceName,
	}

	tmpl, err := template.ParseFiles("utils/gen/templates/service_skeleton.tpl")
	if err != nil {
		return fmt.Errorf("加载骨架模板失败: %v", err)
	}

	outputPath := fmt.Sprintf("utils/generated/service/%s_service_skeleton.go", tableName)
	if err := os.MkdirAll("utils/generated/service", os.ModePerm); err != nil {
		return err
	}

	outFile, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("创建骨架文件失败: %v", err)
	}
	defer outFile.Close()

	if err := tmpl.Execute(outFile, data); err != nil {
		return fmt.Errorf("渲染骨架模板失败: %v", err)
	}

	fmt.Println("✅ 骨架生成成功:", outputPath)
	return nil
}