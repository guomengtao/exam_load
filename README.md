# Gin-Go-Test Project v1.0

## 一、项目简介

Gin-Go-Test 是一个基于 Go + Gin 框架开发的网站系统，包含以下功能：

- MySQL 数据库连接检测
- Redis 连接检测
- API 接口服务
- 自动生成 HTML 试卷和成绩报告
- 定时任务管理
- 环境变量配置管理

---

## 二、项目目录结构

```plaintext
your-app/
├── main.go                        # 入口文件（初始化、路由挂载）
├── .env                           # 环境变量（MySQL、Redis 配置）
├── go.mod                         # Go 模块管理

├── config/
│   └── config.go              # 读取 .env，提供全局配置访问

├── routes/
│   └── routes.go              # 注册所有路由

├── handlers/                      # 处理器（API 接口）
│   ├── user.go                # 用户相关 API
│   ├── question.go            # 问题相关 API
│   ├── html_exam.go           # 生成试卷 HTML 页面
│   └── html_report.go         # 生成成绩报告 HTML 页面

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
│   └── news.html              # 单独的页面

├── utils/                         # 工具类
│   ├── db.go                  # MySQL 连接
│   ├── redis.go               # Redis 连接
│   ├── html.go                # HTML 生成工具
│   └── logger.go              # 日志封装

├── tasks/                         # 定时任务
│   ├── sync_data.go           # 将 Redis 数据持久化到 MySQL
│   └── schedule.go            # 定时器管理

├── middleware/                    # 中间件
│   └── auth.go                # JWT 鉴权中间件

└── README.md                   # 项目文档
```

---

## 三、安装流程

1. 先确认已安装 Go (1.18+)
2. 先确认已安装 MySQL 和 Redis
3. 运行

```bash
# 追加依赖
go mod tidy

# 构建程序，-o 后面是生成的执行文件名称（可自定义）
go build -o hello main.go

# 运行
./hello
```

4. 配置 .env 文件（如：MySQL 连接参数）

---

## 四、环境配置 (.env 示例)

```env
MYSQL_USER=root
MYSQL_PASSWORD=yourpassword
MYSQL_HOST=127.0.0.1
MYSQL_PORT=3306
MYSQL_DB=yourdbname

REDIS_ADDR=127.0.0.1:6379
REDIS_PASSWORD=
REDIS_DB=0
```

---

## 五、API 测试

服务运行后，打开浏览器：

- MySQL 状态检测：

```
GET http://127.0.0.1:8080/api/mysql
```

- Redis 状态检测：

```
GET http://127.0.0.1:8080/api/redis
```

---

## 六、开发指南

- 每墙次扩展功能，先在 handlers/、routes/中添加对应逻辑
- 新添模块（如：异步实现功能），增加在 tasks/与 utils/中
- 静态 HTML 生成，使用 templates/ 配合 utils/html.go 进行模板生成

---

## 七、备注

- `-o hello` 只是生成的执行文件名，可以自定义，但一般一致性更好（一直用 hello 或 your-app名称综一）
- 同一个项目，加上 static/ 静态文件生成，一个 Git 仓库即可，无需分他 Git 仓库
- 建议每天当前展示的逻辑做一个简单备份（git commit）

---

## 八、后续可考虑扩展

- JWT 鉴权中间件，实现用户登陆鉴权
- MySQL 表结构打造，完善数据库系统
- Redis 进行缓存加速
- 用 Docker 包装安装
- 配合日志管理，实现日志分类

---

这份 README 是第一版， v1.0

