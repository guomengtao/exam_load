# Gin-Go-Test Project v1.0.2（围绕1万高并发设计）

**⚠️ 重要开发规范：**

- 禁止直接修改任何由代码生成器生成的文件（如 app/biz、app/services、utils/generated 等目录下的自动生成文件）。
- 如需调整生成内容，请修改模板（如 utils/gen/templates/）或生成器源码。
- 修改模板后，请通过 `go run utils/gen/gen.go -table=表名 -cmd=a` 等命令重新生成并覆盖。

---

## 目录
- [一、项目简介](#一项目简介)
- [二、项目目录结构](#二项目目录结构)
- [三、安装流程](#三安装流程)
- [四、环境配置 (.env 示例)](#四环境配置-env-示例)
- [五、API 测试](#五api-测试)
- [六、开发指南](#六开发指南) (后续准备实现中)
- [七、备注](#七备注) (后续准备实现中)
- [八、后续可考虑扩展](#八后续可考虑扩展) (后续准备实现中)
- [九、版本更新日志](#九版本更新日志)
- [十、正式发布公告](#十正式发布公告)

---

## 一、项目简介

Gin-Go-Test 是一个基于 Go + Gin 框架开发的网站系统，当前架构设计以支持**1万高并发**为核心目标，适合在线教育、考试系统等场景。

---

## 二、项目目录结构

```plaintext
your-app/
├── main.go                        # 入口文件
├── .env                           # 环境变量（MySQL、Redis）
├── go.mod                         # Go 模块管理

├── config/                        # 配置读取
├── routes/                        # 路由挂载
├── handlers/                      # 各功能API接口
├── templates/                     # HTML模板
├── static/                        # 生成的静态HTML文件
├── utils/                         # 工具类（数据库、Redis、HTML生成）
├── tasks/                         # 定时任务
├── middleware/                    # 中间件（如鉴权）
└── README.md                      # 项目说明文档
```

---

## 三、安装流程

1. 确认已安装 Go (推荐 1.18+)
2. 确认已安装 MySQL 和 Redis
3. 运行以下命令：

```bash
go mod tidy
go build -o hello main.go
./hello
```

4. 配置 `.env` 文件

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

- MySQL 状态检测：
  ```bash
  GET http://127.0.0.1:8080/api/mysql
  ```

- Redis 状态检测：
  ```bash
  GET http://127.0.0.1:8080/api/redis
  ```

- 更多接口请参考 routes/handlers。

---

## 六、开发指南

- 所有接口逻辑编写在 `handlers/`
- 所有路由注册在 `routes/`
- 数据库及Redis连接统一使用 `utils/db.go` 和 `utils/redis.go`
- Redis 缓存和高并发逻辑已预置，直接调用

---

## 七、备注

- 项目以模块化划分功能，清晰易维护
- 统一使用 `.env` 进行环境管理
- 使用 Git 保持定期提交，方便版本管理和回滚

---

## 八、后续可考虑扩展

- JWT 用户鉴权
- Docker 镜像封装部署
- Redis 集群支持
- MySQL 读写分离
- Prometheus + Grafana 性能监控

---

## 九、版本更新日志

### v1.0.2

- 新增试卷模板创建接口（写入 MySQL）
- 新增试卷生成接口（基于模板生成考试试卷，同时写入 MySQL 和 Redis）
- 新增 Redis 查询接口（支持分页查询和根据 UUID 查询单条记录）
- 严格遵循 RESTful API 设计规范
- 接口采用模块名+动作名分层设计
- 插入数据时统一生成 UUID，保证 MySQL 与 Redis 一致关联
- 系统采用 Gin 框架，轻量高效，适合中高并发场景
- **更新表结构**：在 `ym_exam_answers` 表中新增 `uid` (用户唯一标识) 和 `score` (用户实际得分) 两个字段
- **更新接口逻辑**：答题记录写入 Redis 时设置七天过期时间，MySQL 同步存储用户答题详情

### 并发支持情况评估

| 并发量级       | 支持情况 | 备注                                           |
|:---------------|:---------|:-----------------------------------------------|
| 5,000 并发     | ✅ 良好支持 | 通过Gin+Redis缓存机制，读请求处理流畅              |
| 10,000 并发    | ✅ 支持   | 需合理设置Redis连接池及MySQL连接池大小，防止阻塞    |
| 20,000 并发    | ⚠️ 需优化 | 需要进一步增加Redis节点/分片，MySQL做读写分离优化  |

> 当前网站架构设计主要以1万高并发为目标，适合中大型在线考试、教育平台场景。

---

## 十、正式发布公告

🎉 **Gin-Go-Test v1.0.2 正式发布！**

- 发布日期：**2025年4月28日**
- 主要内容：支持试卷模板管理、试卷生成、Redis加速查询
- 核心目标：**围绕 1万并发构建高性能系统**
- 适合场景：教育平台、在线考试系统

🔔 未来将继续扩展功能模块，优化高并发性能，打造更强大稳定的平台！
