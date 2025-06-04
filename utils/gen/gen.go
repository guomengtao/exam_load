package main

import (
	"flag"
	"fmt"
	meta "gin-go-test/utils/gen/meta"
	"gin-go-test/utils/genlib"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"log"
	"os"
 	"strings"
	"text/template"
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
	if mysqlUser == "" {
		mysqlUser = "root"
	}
	if mysqlPassword == "" {
		mysqlPassword = "123456"
	}
	if mysqlHost == "" {
		mysqlHost = "127.0.0.1"
	}
	if mysqlPort == "" {
		mysqlPort = "3306"
	}
	if mysqlDB == "" {
		mysqlDB = "gin_go_test"
	}

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
	tableName := flag.String("table", "", "要处理的表名，使用 all 表示处理所有表")
	cmd := flag.String("cmd", "", "要执行的命令：a=生成所有代码，delete=删除代码")
	flag.Parse()

	// 验证参数
	if *tableName == "" {
		log.Fatal("请指定表名，使用 -table 参数")
	}
	if *cmd == "" {
		log.Fatal("请指定命令，使用 -cmd 参数")
	}

	// 初始化表列表管理器
	tableList, err := meta.NewGenTableList()
	if err != nil {
		log.Fatalf("初始化表列表失败: %v", err)
	}

	// 处理表
	if *tableName == "all" {
		// 获取所有活跃的表
		tables := tableList.GetActiveTables()
		if len(tables) == 0 {
			log.Fatal("没有找到需要处理的表")
		}

		// 批量处理所有表
		for _, table := range tables {
			fmt.Printf("\n🚀 开始处理表: %s\n", table)
			if err := processTable(table, *cmd, tableList); err != nil {
				fmt.Printf("❌ 处理表 %s 失败: %v\n", table, err)
				continue
			}
			fmt.Printf("✅ 表 %s 处理完成\n", table)
		}
	} else {
		// 处理单个表
		if err := processTable(*tableName, *cmd, tableList); err != nil {
			log.Fatalf("处理表失败: %v", err)
		}
	}
}

// processTable 处理单个表
func processTable(tableName, cmd string, tableList *meta.GenTableList) error {
	switch cmd {
	case "a":
		// 先校验表是否存在且合规
		if err := validateTable(tableName); err != nil {
			return fmt.Errorf("表校验失败: %v", err)
		}
		// 生成代码前先添加到表列表
		if err := tableList.AddTable(tableName); err != nil {
			return fmt.Errorf("添加表到列表失败: %v", err)
		}

		// 生成代码
		if err := generateCode(tableName); err != nil {
			return fmt.Errorf("生成代码失败: %v", err)
		}

		// 增加生成次数
		if err := tableList.IncrementGenerateCount(tableName); err != nil {
			return fmt.Errorf("更新生成次数失败: %v", err)
		}

	case "delete":
		// 删除代码
		if err := deleteCode(tableName); err != nil {
			return fmt.Errorf("删除代码失败: %v", err)
		}

		// 从表列表中移除（软删除）
		if err := tableList.RemoveTable(tableName); err != nil {
			return fmt.Errorf("从表列表移除失败: %v", err)
		}

	default:
		return fmt.Errorf("未知命令: %s", cmd)
	}

	return nil
}

// validateTable 校验表是否存在且合规（有软删除字段）
func validateTable(tableName string) error {
	// 检查表是否存在
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = DATABASE() AND table_name = ?", tablePrefix+tableName).Scan(&count)
	if err != nil {
		return fmt.Errorf("检查表是否存在失败: %v", err)
	}
	if count == 0 {
		return fmt.Errorf("表 %s 不存在", tablePrefix+tableName)
	}

	// 检查表是否有软删除字段
	var hasDeletedAt bool
	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.columns WHERE table_schema = DATABASE() AND table_name = ? AND column_name = 'deleted_at'", tablePrefix+tableName).Scan(&hasDeletedAt)
	if err != nil {
		return fmt.Errorf("检查软删除字段失败: %v", err)
	}
	if !hasDeletedAt {
		return fmt.Errorf("表 %s 缺少软删除字段 deleted_at", tablePrefix+tableName)
	}

	return nil
}

// generateCode 生成代码
func generateCode(tableName string) error {
	// 生成控制器
	if err := generateController(tableName); err != nil {
		return fmt.Errorf("生成控制器失败: %v", err)
	}

	// 生成业务层
	if err := generateBiz(tableName); err != nil {
		return fmt.Errorf("生成业务层失败: %v", err)
	}

	// 生成服务层
	if err := generateService(tableName); err != nil {
		return fmt.Errorf("生成服务层失败: %v", err)
	}

	// 生成模型
	if err := generateModel(tableName); err != nil {
		return fmt.Errorf("生成模型失败: %v", err)
	}

	return nil
}

// deleteCode 删除代码
func deleteCode(tableName string) error {
	// 删除控制器
	if err := deleteController(tableName); err != nil {
		return fmt.Errorf("删除控制器失败: %v", err)
	}

	// 删除业务层
	if err := deleteBiz(tableName); err != nil {
		return fmt.Errorf("删除业务层失败: %v", err)
	}

	// 删除服务层
	if err := deleteService(tableName); err != nil {
		return fmt.Errorf("删除服务层失败: %v", err)
	}

	// 删除模型
	if err := deleteModel(tableName); err != nil {
		return fmt.Errorf("删除模型失败: %v", err)
	}

	return nil
}

// 以下是生成和删除各个层代码的具体实现
// 这些函数需要根据你的具体需求来实现

func generateController(tableName string) error {
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
	return nil
}

func generateBiz(tableName string) error {
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
	return nil
}

func generateService(tableName string) error {
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
	return nil
}

func generateModel(tableName string) error {
	fmt.Printf("\n📝 生成模型: %s\n", tableName)
	err := genlib.GenerateModelFromTable(tableName)
	if err != nil {
		log.Fatalf("❌ 模型生成失败: %v", err)
	} else {
		printSuccess("模型生成成功", fmt.Sprintf("app/models/%s.go", tableName))
	}
	return nil
}

// deleteController 删除控制器相关代码
func deleteController(tableName string) error {
	fmt.Printf("\n🗑️ 删除控制器: %s\n", tableName)
	
	// 删除控制器文件
	controllerPath := fmt.Sprintf("app/controllers/%s_controller.go", tableName)
	if err := os.Remove(controllerPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除控制器文件失败: %v", err)
	}
	fmt.Printf("✅ 已删除控制器文件: %s\n", controllerPath)

	// 删除控制器骨架文件
	skeletonPath := fmt.Sprintf("utils/generated/controller/%s_skeleton.go", tableName)
	if err := os.Remove(skeletonPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除控制器骨架文件失败: %v", err)
	}
	fmt.Printf("✅ 已删除控制器骨架文件: %s\n", skeletonPath)

	// 清理路由注册
	if err := cleanRouteRegistration(tableName); err != nil {
		return fmt.Errorf("清理路由注册失败: %v", err)
	}

	return nil
}

// deleteBiz 删除业务层相关代码
func deleteBiz(tableName string) error {
	fmt.Printf("\n🗑️ 删除业务层: %s\n", tableName)
	
	// 删除业务层文件
	bizPath := fmt.Sprintf("app/biz/%s_biz.go", tableName)
	if err := os.Remove(bizPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除业务层文件失败: %v", err)
	}
	fmt.Printf("✅ 已删除业务层文件: %s\n", bizPath)

	// 删除业务层骨架文件
	skeletonPath := fmt.Sprintf("utils/generated/biz/%s_biz_skeleton.go", tableName)
	if err := os.Remove(skeletonPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除业务层骨架文件失败: %v", err)
	}
	fmt.Printf("✅ 已删除业务层骨架文件: %s\n", skeletonPath)

	return nil
}

// deleteService 删除服务层相关代码
func deleteService(tableName string) error {
	fmt.Printf("\n🗑️ 删除服务层: %s\n", tableName)
	
	// 删除服务层文件
	servicePath := fmt.Sprintf("app/services/%s_service.go", tableName)
	if err := os.Remove(servicePath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除服务层文件失败: %v", err)
	}
	fmt.Printf("✅ 已删除服务层文件: %s\n", servicePath)

	// 删除服务层骨架文件
	skeletonPath := fmt.Sprintf("utils/generated/service/%s_service_skeleton.go", tableName)
	if err := os.Remove(skeletonPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除服务层骨架文件失败: %v", err)
	}
	fmt.Printf("✅ 已删除服务层骨架文件: %s\n", skeletonPath)

	return nil
}

// deleteModel 删除模型相关代码
func deleteModel(tableName string) error {
	fmt.Printf("\n🗑️ 删除模型: %s\n", tableName)
	
	// 删除模型文件
	modelPath := fmt.Sprintf("app/models/%s.go", tableName)
	if err := os.Remove(modelPath); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("删除模型文件失败: %v", err)
	}
	fmt.Printf("✅ 已删除模型文件: %s\n", modelPath)

	return nil
}

// cleanRouteRegistration 清理路由注册
func cleanRouteRegistration(tableName string) error {
	routeFile := "routes/gen_routes.go"
	
	// 读取路由文件内容
	content, err := os.ReadFile(routeFile)
	if err != nil {
		return fmt.Errorf("读取路由文件失败: %v", err)
	}

	// 构建要删除的路由注册行
	routeLine := fmt.Sprintf("Register%sRoutes", strings.Title(tableName))
	
	// 按行分割内容
	lines := strings.Split(string(content), "\n")
	
	// 过滤掉包含目标路由注册的行
	var newLines []string
	for _, line := range lines {
		if !strings.Contains(line, routeLine) {
			newLines = append(newLines, line)
		}
	}

	// 写回文件
	if err := os.WriteFile(routeFile, []byte(strings.Join(newLines, "\n")), 0644); err != nil {
		return fmt.Errorf("写回路由文件失败: %v", err)
	}

	fmt.Printf("✅ 已清理路由注册: %s\n", routeLine)
	return nil
}
