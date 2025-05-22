package genlib

import (
    "fmt"
    "os"
    "text/template"
)

func GenerateController(tableName,moduleName string) error {
    tmplPath := "utils/gen/templates/controller.tpl"

    // 读取模板
    tmpl, err := template.ParseFiles(tmplPath)
    if err != nil {
        return fmt.Errorf("❌ 加载模板失败: %v", err)
    }

    outputPath := "app/controllers/hello_controller.go"

    // 判断是否覆盖
    if _, err := os.Stat(outputPath); err == nil {
        var input string
        fmt.Printf("⚠️ 文件已存在 %s，是否覆盖？(y/n): ", outputPath)
        fmt.Scanln(&input)
        if input != "y" {
            fmt.Println("✅ 已跳过生成。")
            return nil
        }
    }

    outFile, err := os.Create(outputPath)
    if err != nil {
        return fmt.Errorf("❌ 创建文件失败: %v", err)
    }
    defer outFile.Close()

    data := map[string]string{
        "ModuleName":     moduleName,
        "ControllerName": "HelloController",
    }

    if err := tmpl.Execute(outFile, data); err != nil {
        return fmt.Errorf("❌ 渲染模板失败: %v", err)
    }

    fmt.Println("✅ 控制器已生成到:", outputPath)
    return nil
}