your-app/
├── main.go                        # 入口文件（初始化、路由挂载）
├── .env                           # 环境变量（MySQL、Redis等配置）
├── go.mod                         # Go 模块管理文件

├── config/
│   └── config.go                  # 加载 .env，提供全局配置访问（封装 os.Getenv）

├── routes/
│   └── routes.go                  # 注册所有路由（API + HTML）

├── handlers/                      # 处理器（按功能模块或数据表拆分）
│   ├── user.go                    # 用户表的增删改查 API
│   ├── question.go                # 问题表 API
│   ├── html_exam.go              # 生成试卷 HTML 页面
│   ├── html_report.go            # 生成成绩报告 HTML 页面

├── templates/                     # HTML 模板
│   ├── exam_template.html
│   └── report_template.html

├── static/                        # 生成的静态 HTML 页面
│   ├── exams/
│   │   ├── exam_001.html
│   │   └── exam_002.html
│   ├── reports/
│   │   ├── report_001.html
│   │   └── report_002.html
│   └── news.html                  # 单独的页面

├── utils/                         # 工具类
│   ├── db.go                      # 数据库连接（读取 .env）
│   ├── redis.go                   # Redis 连接（读取 .env）
│   ├── html.go                    # 通用 HTML 生成工具
│   └── logger.go                  # 日志封装（可选）

├── tasks/                         # ✅ 定时任务 / 后台服务逻辑
│   └── sync_data.go              # 定期将 Redis 数据写入 MySQL
│   └── schedule.go               # 定时器管理（cron 定义入口）

├── middleware/                   # 中间件（可选）
│   └── auth.go                   # 鉴权中间件（JWT、API 权限等）

└── README.md                     # 项目文档（建议写）