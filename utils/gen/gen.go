package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "gin-go-test/utils/genlib"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "text/template"
    "github.com/joho/godotenv"
    "strings"
    meta "gin-go-test/utils/gen/meta"
)

var db *sqlx.DB
var moduleName string
var tablePrefix string

func init() {
    // 先加载 .env 文件
    errEnv := godotenv.Load()
    if errEnv != nil {
        log.Printf("⚠️ 无法加载 .env 文件: %v", errEnv)
    }
    // 从环境变量读取模块名，默认值为gin-go-test
    moduleName = os.Getenv("MODULE_NAME")
    if moduleName == "" {
        moduleName = "gin-go-test"
    }

    // 从环境变量读取表前缀
    tablePrefix = os.Getenv("TABLE_PREFIX")
    if tablePrefix == "" {
        tablePrefix = "tm_"
    }

    // 连接数据库，参数从环境变量读取
    mysqlUser := os.Getenv("MYSQL_USER")
    mysqlPassword := os.Getenv("MYSQL_PASSWORD")
    mysqlHost := os.Getenv("MYSQL_HOST")
    mysqlPort := os.Getenv("MYSQL_PORT")
    mysqlDB := os.Getenv("MYSQL_DB")
    if mysqlUser == "" { mysqlUser = "root" }
    if mysqlPassword == "" { mysqlPassword = "123456" }
    if mysqlHost == "" { mysqlHost = "127.0.0.1" }
    if mysqlPort == "" { mysqlPort = "3306" }
    if mysqlDB == "" { mysqlDB = "gin_go_test" }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        mysqlUser, mysqlPassword, mysqlHost, mysqlPort, mysqlDB)

    var err error
    db, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        log.Fatalf("❌ 数据库连接失败: %v", err)
    }
}

func printSuccess(msg string, path string) {
    fmt.Printf("✅ %-20s: %s\n", msg, path)
}

func main() {
    // 解析命令行参数
    tableName := flag.String("table", "", "表名")
    cmd := flag.String("cmd", "", "命令")
    flag.Parse()

    if *tableName == "" || *cmd == "" {
        fmt.Println("❌ 请指定表名和命令，例如: go run utils/gen/gen.go -table=users -cmd=a")
        os.Exit(1)
    }

    // 检查表是否存在
    exists, err := checkTableExists(*tableName)
    if err != nil {
        fmt.Printf("❌ 检查表 %s 失败: %v\n", *tableName, err)
        os.Exit(1)
    }
    if !exists {
        fmt.Printf("❌ 表 %s 不存在，请先创建表\n", *tableName)
        os.Exit(1)
    }

    // 根据命令生成代码
    switch *cmd {
    case "a":
        generateAll(*tableName)
    case "c":
        generateController(*tableName)
    case "b":
        generateBiz(*tableName)
    case "s":
        generateService(*tableName)
    case "m":
        generateModel(*tableName)
    default:
        fmt.Printf("❌ 未知命令: %s\n", *cmd)
        os.Exit(1)
    }
}

// checkTableExists 检查表是否存在
func checkTableExists(tableName string) (bool, error) {
    // 只在这里加前缀
    if !strings.HasPrefix(tableName, tablePrefix) {
        tableName = tablePrefix + tableName
    }
    var count int
    query := fmt.Sprintf("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = '%s'", tableName)
    err := db.QueryRow(query).Scan(&count)
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

// generateAll 生成所有代码
func generateAll(tableName string) {
    fmt.Printf("\n🚀 开始生成 %s 相关代码...\n", tableName)
    generateController(tableName)
    generateBiz(tableName)
    generateService(tableName)
    generateModel(tableName)
    fmt.Printf("\n✅ %s 相关代码生成完成\n", tableName)
    // 打印生成汇总
    fmt.Printf("\n📊 生成汇总:\n")
    fmt.Printf("  ✅ 控制器: 1 个\n")
    fmt.Printf("  ✅ 业务逻辑: 1 个\n")
    fmt.Printf("  ✅ 服务层: 1 个\n")
    fmt.Printf("  ✅ 数据模型: 1 个\n")
    fmt.Printf("  ✅ 接口总数: 5 个\n\n")
    fmt.Printf("🎉 生成成功，祝你编程愉快！🧠⚙️\n")
}

// generateController 生成控制器代码
func generateController(tableName string) {
    fmt.Printf("\n📝 生成控制器: %s\n", tableName)
    err := genlib.GenerateControllerWithAppend(tableName, moduleName)
    if err != nil {
        log.Fatalf("❌ 控制器生成失败: %v", err)
    } else {
        printSuccess("控制器生成成功", fmt.Sprintf("app/controllers/%s_controller.go", tableName))
        // 获取表注释
        var tableComment string
        err := db.QueryRow("SELECT table_comment FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?", tablePrefix+tableName).Scan(&tableComment)
        if err != nil || tableComment == "" {
            tableComment = tableName // 没有注释就用表名兜底
        }
        // ANSI 颜色
        green := "\033[32m"
        blue := "\033[34m"
        reset := "\033[0m"
        // 打印注册的接口和接口地址
        fmt.Printf("🔗 注册接口:\n")
        fmt.Printf(green+"  - GET    /api/%s/         # 获取%s列表\n"+reset, strings.ToLower(tableName), tableComment)
        fmt.Printf(blue+"  - GET    /api/%s/:id      # 获取%s详情\n"+reset, strings.ToLower(tableName), tableComment)
        fmt.Printf(green+"  - POST   /api/%s/         # 创建%s\n"+reset, strings.ToLower(tableName), tableComment)
        fmt.Printf(blue+"  - PUT    /api/%s/:id      # 更新%s\n"+reset, strings.ToLower(tableName), tableComment)
        fmt.Printf(green+"  - DELETE /api/%s/:id      # 删除%s\n"+reset, strings.ToLower(tableName), tableComment)
    }

    err = genlib.GenerateControllerSkeleton(db.DB, tableName, moduleName, true)
    if err != nil {
        log.Fatalf("❌ 控制器骨架生成失败: %v", err)
    } else {
        printSuccess("控制器骨架生成成功", fmt.Sprintf("utils/generated/controller/%s_skeleton.go", tableName))
    }
}

// generateBiz 生成业务层代码
func generateBiz(tableName string) {
    fmt.Printf("\n📝 生成业务层: %s\n", tableName)
    err := genlib.GenerateBiz(tableName, true)
    if err != nil {
        log.Fatalf("❌ 业务层生成失败: %v", err)
    } else {
        printSuccess("业务层生成成功", fmt.Sprintf("app/biz/%s_biz.go", tableName))
    }

    err = genlib.GenerateBizSkeleton(tableName, true)
    if err != nil {
        log.Fatalf("❌ 业务骨架生成失败: %v", err)
    } else {
        printSuccess("业务骨架生成成功", fmt.Sprintf("utils/generated/biz/%s_biz_skeleton.go", tableName))
    }
}

// generateService 生成服务层代码
func generateService(tableName string) {
    fmt.Printf("\n📝 生成服务层: %s\n", tableName)
    err := genlib.GenerateServiceFromTable(tableName)
    if err != nil {
        log.Fatalf("❌ 服务层生成失败: %v", err)
    } else {
        printSuccess("服务层生成成功", fmt.Sprintf("app/services/%s_service.go", tableName))
    }

    // 注册 camelCase 函数
    funcMap := template.FuncMap{
        "camelCase": meta.CamelCase,
    }
    tmpl, err := template.New("service_skeleton.tpl").Funcs(funcMap).ParseFiles("utils/gen/templates/service_skeleton.tpl")
    if err != nil {
        log.Fatalf("❌ 加载服务骨架模板失败: %v", err)
    }
    err = genlib.GenerateServiceSkeleton(db.DB, tableName, tmpl)
    if err != nil {
        log.Fatalf("❌ 服务骨架生成失败: %v", err)
    } else {
        printSuccess("服务骨架生成成功", fmt.Sprintf("utils/generated/service/%s_service_skeleton.go", tableName))
    }
}

// generateModel 生成模型代码
func generateModel(tableName string) {
    fmt.Printf("\n📝 生成模型: %s\n", tableName)
    err := genlib.GenerateModelFromTable(tableName)
    if err != nil {
        log.Fatalf("❌ 模型生成失败: %v", err)
    } else {
        printSuccess("模型生成成功", fmt.Sprintf("app/models/%s.go", tableName))
    }
}
