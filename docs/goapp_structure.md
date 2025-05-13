量身整理的最终项目结构，适合你目前使用 Gin + Gorm/sqlx + Swagger + 单元测试 + CLI 代码生成器 的 Go 项目。

⸻

✅ 推荐项目结构（整理版）

.
├── app/                            # 主业务代码目录
│   ├── controllers/               # 控制器（接口处理）
│   │   ├── admin_controller.go
│   │   └── ...
│   ├── middleware/                # 中间件（鉴权、日志等）
│   │   └── auth.go
│   ├── models/                    # 数据模型（Gorm Struct）
│   │   └── admin.go
│   ├── requests/                  # 请求参数校验 struct
│   │   └── admin_request.go
│   └── services/                  # 业务逻辑处理
│       └── admin_service.go
│
├── config/                        # 配置加载，如 env/config.go
│   └── config.go
│
├── docs/                          # Swagger 文档生成目录
│   ├── docs.go
│   ├── swagger.yaml
│   ├── swagger.json
│   └── GEN_CURD.md               # 你手动整理的 CRUD 文档说明
│
├── internal/                      # 内部模块，仅供本项目使用
│   ├── cron/                      # 定时任务
│   └── middleware/               # 公共中间件（通用处理）
│
├── routes/                        # 路由注册目录（适合放顶层）
│   └── routes.go
│
├── scripts/                       # 工具脚本、代码生成器
│   ├── gen_crud.go                # go run scripts/gen_crud.go --table=xxx
│   └── templates/                 # 模板文件
│       ├── controller.tmpl
│       ├── model.tmpl
│       ├── service.tmpl
│       └── ...
│
├── static/                        # 静态网页资源
│   ├── index.html
│   └── hello.html
│
├── storage/                       # 动态资源（上传、日志等）
│   └── uploads/
│       └── README.md              # 提示：上传资源请放此处
│
├── tests/                         # 集成/黑盒测试代码（可选）
│   └── admin_test.go              # 示例：集成测试接口逻辑
│
├── utils/                         # 工具函数、数据库、Redis 初始化
│   ├── db.go                      # database/sql 初始化
│   ├── db_gorm.go                 # GORM 初始化
│   ├── db_sqlx.go                 # sqlx 初始化
│   ├── redis.go
│   └── status.go
│
├── main.go                        # 程序入口
├── go.mod
├── go.sum
├── README.md
├── .env                           # 环境变量文件
└── TREE.md                        # 项目结构记录文档



⸻

🧠 每个目录的说明（简洁）

目录	用途说明
app/	核心业务目录，MVC结构清晰
config/	配置加载和环境变量封装
docs/	Swagger 自动生成文档文件
routes/	路由定义注册
scripts/	CLI 工具及模板
static/	静态 HTML 文件
storage/	上传文件、日志目录
tests/	统一集成测试目录（可选）
utils/	公共工具：数据库、缓存等



 