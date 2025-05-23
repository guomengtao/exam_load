package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strings"

    "github.com/joho/godotenv"
    "gin-go-test/utils/genlib"
)

func main() {
    // 先加载 .env 文件
    err := godotenv.Load(".env")
    if err != nil {
        log.Printf("⚠️ 无法加载 .env 文件: %v", err)
    }

    // 从环境变量读取模块名，默认值为gin-go-test
    moduleName := os.Getenv("MODULE_NAME")
    if moduleName == "" {
        moduleName = "gin-go-test"
    }

    cmd := flag.String("cmd", "", "命令组合: a=all, c=controller, r=route, k=skeleton")
    table := flag.String("table", "", "要生成的表名，例如: member")
    flag.Parse()

    if *table == "" {
        log.Fatal("❌ 请使用 -table 参数指定表名，例如: -table=member")
    }
    tableName := *table

    // 支持组合命令
    cmdMap := make(map[rune]bool)
    for _, ch := range *cmd {
        cmdMap[ch] = true
    }

    // a=all -> 执行所有子命令
    if cmdMap['a'] {
        cmdMap['c'] = true
        cmdMap['r'] = true
        cmdMap['k'] = true
        cmdMap['m'] = true
        cmdMap['s'] = true
        cmdMap['b'] = true
    }

    if cmdMap['c'] {
        if err := genlib.GenerateControllerWithAppend (tableName, moduleName); err != nil {
            log.Fatalf("GenerateControllerWithAppend  error: %v", err)
        }
        fmt.Println("✅ 控制器生成成功")
        fmt.Printf("📁 生成文件路径: app/controllers/%s_controller.go\n", tableName)
    }

    if cmdMap['r'] {
        routes := []genlib.RouteInfo{
            {PackageName: tableName, RegisterFunc: "Register" + strings.Title(tableName) + "Routes"},
        }
         if err := genlib.GenerateGenRoutesFile(routes,moduleName); err != nil {
            log.Println("❌ 生成路由注册失败:", err)
        } else {
            fmt.Println("📁 路由注册文件: routes/gen_routes.go")
        }
    }

    if cmdMap['k'] {
        overwrite := true // 可根据实际需要设为 false
        if err := genlib.GenerateSkeletonWithAppend(tableName, moduleName, overwrite); err != nil {
            log.Println("❌ 生成骨架失败:", err)
        }
    }
    if cmdMap['m'] {
        err := genlib.GenerateModelFromTable(tableName)
        if err != nil {
            log.Fatalf("生成模型失败: %v", err)
        } else {
            log.Println("✅ 模型生成成功")
        }
    }
    if cmdMap['s'] {
        err := genlib.GenerateServiceFromTable(tableName)
        if err != nil {
            log.Fatalf("生成服务失败: %v", err)
        } else {
            log.Println("✅ 服务生成成功")
        }

        if err := genlib.GenerateServiceSkeleton(tableName); err != nil {
            log.Fatalf("生成服务骨架失败: %v", err)
        } else {
            log.Println("✅ 服务骨架生成成功")
        }
    }
    if cmdMap['b'] {
        err := genlib.GenerateBiz(tableName)
        if err != nil {
            log.Fatalf("生成业务层失败: %v", err)
        } else {
            log.Println("✅ 业务层生成成功")
        }

        err = genlib.GenerateBizSkeleton(tableName)
        if err != nil {
            log.Fatalf("生成业务骨架失败: %v", err)
        } else {
            log.Println("✅ 业务骨架生成成功")
        }
    }
}
