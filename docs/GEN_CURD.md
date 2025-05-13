 完整方案的文档整理（无代码版本），重点说明每个文件、路径及职责。该方案采用：

✅ Gorm
✅ 接口基础5个：增删改查 + 搜索
✅ 测试用表加 _test 后缀，同库隔离测试数据
✅ 自动生成 Swag 文档 + 单元测试 + 接口测试模板

⸻

📘 项目规范方案文档：基于 Gorm 的 CRUD 接口生成体系

💡 方案概述

本方案旨在为每个数据库表生成一套完整的、可维护的 REST API 接口，包含：
	•	5 个核心接口（增、删、改、查、搜索）
	•	Swag 自动文档
	•	接口自动化测试（HTTP）
	•	单元测试（Service 层）
	•	测试表采用同库、同结构，表名加 _test 后缀
	•	基于 Gorm 自动迁移能力，自动生成表结构

⸻

📁 目录结构

app/
├── controllers/        // 控制器（接口处理）
│   ├── member_controller.go
│   └── member_controller_test.go       # 接口自动化测试
├── models/             // 数据模型
│   ├── member.go
│   └── member_test.go                  # 测试模型（使用 _test 表）
├── services/           // 业务逻辑层
│   ├── member_service.go
│   └── member_service_test.go          # 单元测试
scripts/
└── gen_crud.go         # 自动生成 CLI 工具
docs/
└── swagger.yaml        # 自动生成的 Swagger 文档



⸻

📌 各模块说明

1. models/member.go
	•	定义正式表 tm_member 的结构
	•	提供 Gorm 模型定义
	•	使用 TableName() 设置表前缀

⸻

2. models/member_test.go
	•	定义测试用表 tm_member_test 的结构
	•	同步模型结构，用于单元测试时自动迁移
	•	使用 TableName() 指向 _test 后缀表

⸻

3. services/member_service.go
	•	封装业务逻辑
	•	通常包含：查询列表、搜索分页、新增、更新、删除等方法

⸻

4. services/member_service_test.go
	•	单元测试文件
	•	使用 _test 表操作
	•	测试服务层逻辑是否正确，如字段校验、数据结构、返回值等

⸻

5. controllers/member_controller.go
	•	接收请求 → 调用服务逻辑 → 返回 JSON 响应
	•	包括 5 个基本接口：
	•	GET /member/list
	•	POST /member/create
	•	PUT /member/update/:id
	•	DELETE /member/delete/:id
	•	POST /member/search

⸻

6. controllers/member_controller_test.go
	•	使用 httptest.NewRecorder() 发起真实 HTTP 请求
	•	断言响应结果，适合 CI 自动化测试
	•	测试 API 层返回结构和路由正确性

⸻

7. scripts/gen_crud.go
	•	命令行工具：自动生成指定表的 CRUD 全套结构
	•	用法：

go run scripts/gen_crud.go --table=member



⸻

8. docs/swagger.yaml
	•	通过 swag init 自动生成
	•	接口注释规范统一由 controllers/member_controller.go 中维护
	•	支持在线文档、自动 Try it out 接口测试

⸻

🧪 测试机制说明

✅ 测试表设计（推荐）
	•	在 .env 文件设置前缀：TABLE_PREFIX=tm_
	•	测试时通过模型 TableName() 自动拼接为 tm_member_test
	•	使用 gorm.AutoMigrate() 自动创建测试表
	•	测试数据写入后自动清理（建议）

⸻

✅ 单元测试运行方式

go test ./app/services -v

✅ 接口测试运行方式

go test ./app/controllers -v



⸻

🏁 总结：方案优势

特性	描述
接口标准化	每张表生成结构一致的接口与文档
测试隔离安全	_test 表数据不污染正式数据
高度可维护	接口、模型、服务层职责分明
自动生成能力	CLI 工具快速生成完整结构
Swag + 测试双支持	一键文档 + 自动化测试



 
---

## 🔁 路由生成说明

本方案还支持自动生成对应的路由注册代码。

### 📄 路由生成格式（以 member 为例）：

```go
r.GET("/member/list", controllers.GetMemberList)
r.POST("/member/create", controllers.CreateMember)
r.PUT("/member/update/:id", controllers.UpdateMember)
r.DELETE("/member/delete/:id", controllers.DeleteMember)
r.POST("/member/search", controllers.SearchMember)
```

这些代码将自动写入：

```
routes/
└── auto_routes.go
```

或追加到统一的 `routes.go` 中。

### ✅ 自动路由注册方案

- 使用 `scripts/gen_crud.go` 工具自动追加路由代码。
- 可根据模板动态拼接控制器函数名与路径。
- 保持统一命名规范，方便维护。

---

## 🛠️ 自动生成前的准备工作

请准备以下信息：

1. **表名（如 `member`）**
2. **数据库字段结构**：
   - 字段名（如 `username`, `email`）
   - 字段类型（如 `string`, `int`, `time.Time`）
   - 是否为主键、自增、nullable 等属性
3. **可选项**：
   - 是否需要软删除支持
   - 是否包含 created_at / updated_at 等标准字段

---

## 🧠 自动迁移说明（Gorm）

自动迁移是 Gorm 的一个核心功能，支持将模型结构同步到数据库表结构：

### ✨ 支持的功能

| 功能 | 说明 |
|------|------|
| 自动创建表 | 根据模型自动创建新表 |
| 补充字段 | 新增字段将自动同步到数据库 |
| 自动索引 | 根据 tag 自动创建索引 |
| 自动设置字段类型 | 自动匹配合适的数据库字段类型 |
| 字段更新 | 修改 tag 后字段属性更新（谨慎使用） |

### ⚠️ 注意事项

- **字段删除不会自动执行**，防止误删数据。
- 若字段类型修改，可能不会完全同步，需手动迁移。
- 更复杂的字段变更建议手写 migration 脚本。
- 可搭配 `gormigrate` 等库进行版本化迁移管理。

---

下面是对你文档中提到的 gen_crud.go CLI 工具的完整生成细节设计和建议方案，适合作为文档的“生成器实现说明”章节：

⸻

🔧 CLI 工具生成细节说明（gen_crud.go）

📌 目标

执行命令：

go run scripts/gen_crud.go --table=member

自动生成以下内容：
	•	GORM 模型文件（含 _test 版本）
	•	Service 层代码和测试代码
	•	Controller 层接口和接口测试代码
	•	Swag 注释自动注入
	•	路由注册代码自动追加
	•	Swag 文档重新生成（可选）

⸻

🧱 文件结构模板说明

✅ 文件路径规则（根据表名 member）：

层	路径和文件	文件作用
Model	app/models/member.go	GORM 模型定义
	app/models/member_test.go	_test 测试表模型
Service	app/services/member_service.go	CRUD 逻辑函数
	app/services/member_service_test.go	单元测试逻辑
Controller	app/controllers/member_controller.go	接口定义和注释
	app/controllers/member_controller_test.go	接口测试
Route	routes/auto_routes.go	自动追加路由



⸻

✍️ CLI 工具生成流程

1. 解析参数

table := flag.String("table", "", "数据库表名（不含前缀）")

2. 模板渲染
	•	使用 text/template 模板引擎。
	•	模板文件建议放在 scripts/templates/ 目录中，支持可自定义修改。

模板文件	对应生成目标文件
model.tmpl	models/{{.Table}}.go
model_test.tmpl	models/{{.Table}}_test.go
service.tmpl	services/{{.Table}}_service.go
service_test.tmpl	services/{{.Table}}_service_test.go
controller.tmpl	controllers/{{.Table}}_controller.go
controller_test.tmpl	controllers/{{.Table}}_controller_test.go
route_append.tmpl	追加到 routes/auto_routes.go

3. 命名转换规则

原始值	示例	转换结果
表名	member	类名 Member
字段名	user_id	Go 字段名 UserID



⸻

📜 自动追加路由的规则
	•	追加到 routes/auto_routes.go，或统一注册函数中。
	•	采用 AppendFile() 方式向末尾追加以下内容：

r.GET("/member/list", controllers.GetMemberList)
r.POST("/member/create", controllers.CreateMember)
r.PUT("/member/update/:id", controllers.UpdateMember)
r.DELETE("/member/delete/:id", controllers.DeleteMember)
r.POST("/member/search", controllers.SearchMember)



⸻

🧼 生成后建议操作
	•	✅ 自动执行 swag init，生成更新的 Swagger 文档。
	•	✅ 自动运行单元测试或接口测试验证生成结果（可配置）。
	•	✅ 可加上生成后输出指引路径与“下一步操作提示”。

⸻

📌 支持字段结构（后续可拓展）

目前以手动字段定义为主，如：

type Member struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"type:varchar(100);not null" json:"username"`
	Email     string         `gorm:"type:varchar(100);unique" json:"email"`
	CreatedAt time.Time      `json:"created_at"`
}

👉 可后续拓展为自动从数据库读取字段结构（通过 INFORMATION_SCHEMA）实现更完整的自动化。

⸻

✨ 拓展功能建议

功能	描述
字段模板参数	自动生成字段结构
同步 Swag 注释	在 controller 中写入注释生成文档
表名驼峰转换	从 member_activity_log → MemberActivityLog
软删除支持	是否自动加入 gorm.DeletedAt 字段
指定字段选择	支持命令行指定生成哪些字段



⸻

是否需要我为你生成一套 scripts/templates/*.tmpl 文件和 gen_crud.go 框架代码？可以立刻开始。

