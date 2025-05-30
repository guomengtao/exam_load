# Gin Go Test

使用 Gin 框架的 Go 项目，包含代码生成工具和完整的测试。

## 功能特点

- 使用 Gin 框架的 RESTful API
- 模型、服务和控制器的代码生成工具
- 完整的测试覆盖
- SQLite 数据库支持
- 使用指针字段的批量更新操作
- 清晰的分层架构

## 项目结构

```
.
├── app/
│   ├── controllers/    # API 控制器
│   ├── models/        # 数据模型
│   ├── services/      # 业务逻辑
│   └── validators/    # 请求验证
├── utils/
│   └── gen/          # 代码生成工具
└── tests/            # 测试文件
```

## 代码生成

项目包含代码生成工具，可以生成：
- 带有指针字段的模型（用于可选更新）
- 包含 CRUD 操作的服务
- RESTful 端点的控制器
- 服务和控制器的测试

### 使用方法

```bash
# 生成模型
go run utils/gen/gen.go -table=表名 -cmd=m

# 生成服务
go run utils/gen/gen.go -table=表名 -cmd=s

# 生成控制器
go run utils/gen/gen.go -table=表名 -cmd=c
```

## 测试

项目包含完整的测试：
- 服务层操作
- 批量更新功能
- 指针字段处理
- 数据库操作

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定测试
go test -v ./app/services/...
```

## 开发

1. 克隆仓库
2. 安装依赖
3. 运行开发服务器

```bash
go mod download
go run main.go
```

## 许可证

MIT 