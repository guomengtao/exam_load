# 开发文档

## 架构说明

本项目采用 ✅ Clean Architecture（清晰架构 / 干净架构）结构分配代码。
- 各层职责分明，依赖关系清晰，便于扩展和维护。
- 核心接口代码几乎只存在于 service 骨架层，其他文件（如 controller、biz）主要负责调用和转发，biz 层目前业务代码较少，主要依赖 service 层。

## 文件统计

| 类型 | 数量 |
|------|------|
| 新加文件 | 2 |
| 改动文件 | 4 |
| 总数 | 6 |
| 生成文件总数 | 8 |

## 文件职责说明

### 生成文件职责

1. **控制器 (`app/controllers/{{.VarName}}_controller.go`)**
   - 职责：处理 HTTP 请求，调用业务层，返回响应。
   - 主要功能：路由注册、参数解析、调用业务层、错误处理、响应封装。

2. **服务层 (`app/services/{{.VarName}}_service.go`)**
   - 职责：实现业务逻辑，调用数据层，处理事务。
   - 主要功能：批量创建、批量更新、数据验证、事务管理。
   - **关键点**：核心接口代码主要在 service 骨架层，其他文件基本是调用。

3. **业务层 (`app/biz/{{.VarName}}_biz.go`)**
   - 职责：实现核心业务逻辑，调用服务层。
   - 主要功能：业务规则校验、调用服务层、错误处理。
   - **关键点**：业务代码较少，主要依赖 service 层。

4. **模型层 (`app/models/{{.VarName}}.go`)**
   - 职责：定义数据结构，映射数据库表。
   - 主要功能：字段定义、验证标签、数据库映射。

5. **验证器 (`app/validators/{{.VarName}}_validator.go`)**
   - 职责：实现字段验证逻辑。
   - 主要功能：必填验证、长度验证、格式验证。

6. **骨架层**
   - `utils/generated/controller/{{.VarName}}_skeleton.go`：控制器骨架，提供基础接口。
   - `utils/generated/biz/{{.VarName}}_biz_skeleton.go`：业务层骨架，提供基础接口。
   - `utils/generated/service/{{.VarName}}_service_skeleton.go`：服务层骨架，提供基础接口。

7. **单元测试 (`app/controllers/{{.VarName}}_controller_test.go`)**
   - 职责：验证控制器功能。
   - 主要功能：测试批量创建接口、参数验证、错误处理。

### 文件依赖关系

- **控制器** 依赖 **业务层**：控制器调用业务层处理请求。
- **业务层** 依赖 **服务层**：业务层调用服务层实现核心逻辑。
- **服务层** 依赖 **模型层** 和 **验证器**：服务层使用模型层定义数据结构，调用验证器进行字段验证。
- **单元测试** 依赖 **控制器**：单元测试验证控制器功能。

## 协作规则

1. **仅模板修改原则**
    - 所有新功能（如批量 CRUD、验证器）必须通过编辑或创建模板文件实现。
    - 仅允许修改或创建以下文件：
        - `utils/gen/templates/controller/controller.tpl`
        - `utils/gen/templates/service/service.tpl`
        - `utils/gen/templates/biz/biz.tpl`
        - `utils/gen/templates/model.tpl`
        - `utils/gen/templates/validator/validator.tpl`（新建）
        - `utils/gen/templates/validator/validator.go`（新建）
    - 除非明确讨论并批准，否则不得修改 main.go、骨架模板或任何其他文件。

2. **不改变生成模式或命令**
    - 代码生成命令和 main.go 逻辑不得更改。
    - 所有生成的文件路径和命名约定必须保持不变。

3. **增量功能添加**
    - 仅添加新功能（如批量创建、批量更新、验证），不得删除或重构现有功能。
    - 现有接口、路由和骨架逻辑必须保留。
    - **特别注意**：原有 count 和 list 接口必须保留，模板中不得删除相关代码。

4. **超出范围变更的审批**
    - 如果功能无法在上述模板文件中实现，必须提供明确分析并征得项目所有者书面批准，方可进行超出允许范围的变更。

5. **文档和沟通**
    - 所有新功能、变更和规则必须记录在本文件中。
    - 任何混淆或歧义必须在书面确认后方可继续。

6. **禁止因调试修改其他文件**
    - 禁止因调试或测试而修改、增加任何其他文件。
    - 所有调试和测试必须严格在允许的模板文件范围内进行。

## 开发总目标

- 实现批量创建 API（支持事务、部分成功、最多 30 条记录、简单验证）
- 实现基础字段验证（必填、最大长度等）
- 实现批量操作的详细错误报告
- 最终生成 8 个主要文件，分布于如下目录：
    - `app/controllers/{{.VarName}}_controller.go`
    - `app/services/{{.VarName}}_service.go`
    - `app/biz/{{.VarName}}_biz.go`
    - `app/models/{{.VarName}}.go`
    - `app/validators/{{.VarName}}_validator.go`
    - `utils/generated/controller/{{.VarName}}_skeleton.go`
    - `utils/generated/biz/{{.VarName}}_biz_skeleton.go`
    - `utils/generated/service/{{.VarName}}_service_skeleton.go`
    - 单元测试：`app/controllers/{{.VarName}}_controller_test.go`
- 支持标准 RESTful 增删改查接口，查接口支持简单搜索：
    - 查询（带搜索）：`GET /api/{{.RoutePath}}/list?keyword=xxx&page=1&pageSize=10`
    - 计数：`GET /api/{{.RoutePath}}/count`
    - 创建：`POST /api/{{.RoutePath}}/{{.TableName}}`
    - 更新：`PUT /api/{{.RoutePath}}/{{.TableName}}`
    - 删除：`DELETE /api/{{.RoutePath}}/{{.TableName}}/:id`

## 生成文件路径和名称

- 控制器：`app/controllers/{{.VarName}}_controller.go`
- 服务层：`app/services/{{.VarName}}_service.go`
- 业务层：`app/biz/{{.VarName}}_biz.go`
- 模型层：`app/models/{{.VarName}}.go`
- 验证器：`app/validators/{{.VarName}}_validator.go`
- 骨架层：`utils/generated/controller/{{.VarName}}_skeleton.go`、`utils/generated/biz/{{.VarName}}_biz_skeleton.go`、`utils/generated/service/{{.VarName}}_service_skeleton.go`
- 单元测试：`app/controllers/{{.VarName}}_controller_test.go`（新建，用于验证最小目标功能）

## 路由地址

> 说明：以下所有 `role` 仅为演示用表名，实际生成时会根据不同表名自动替换。

- 原有接口：
  - 计数：`GET /api/role/count`
  - 列表：`GET /api/role/list`
- 新增接口：
  - 批量创建：`POST /api/role/role`
  - 批量更新：`PUT /api/role/role`
  - 删除：`DELETE /api/role/role/:id`
  - 查询（带搜索）：`GET /api/role/list?keyword=xxx&page=1&pageSize=10`

## 生成器执行命令

> 说明：`role` 仅为演示用表名，实际请替换为你要生成的表名。

生成器执行命令示例：
```bash
go run utils/gen/gen.go -table=role
```

## 最小第一步目标

- 在 controller.tpl 中添加批量创建接口路由和处理函数
- 在 service.tpl 中添加批量创建服务方法
- 在 biz.tpl 中添加批量创建业务逻辑
- 在 model.tpl 中添加基础字段验证标签
- 创建 validator.tpl 和 validator.go 模板，实现基础验证逻辑
- 创建单元测试文件 `app/controllers/{{.VarName}}_controller_test.go`，验证最小目标功能

---

**本文档是开发规则和协作的唯一参考依据。** 