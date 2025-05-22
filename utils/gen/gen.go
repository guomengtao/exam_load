package main

import (
    "flag"
    "fmt"
    "log"
    "os"

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
    }

    if cmdMap['c'] {
        if err := genlib.GenerateController(tableName, moduleName); err != nil {
            log.Fatalf("GenerateController error: %v", err)
        }
        fmt.Println("✅ 控制器生成成功")
    }

    if cmdMap['r'] {
        // routes := []genlib.RouteInfo{
        //     {PackageName: tableName, RegisterFunc: "Register" + genlib.ToCamelCase(tableName) + "Routes"},
        // }
        // if err := genlib.GenerateGenRoutesFile(routes, moduleName); err != nil {
        //     log.Println("❌ 生成路由注册失败:", err)
        // }
    }

    if cmdMap['k'] {
        // if err := genlib.GenerateSkeleton(tableName, moduleName); err != nil {
        //     log.Println("❌ 生成骨架失败:", err)
        // }
    }
}
