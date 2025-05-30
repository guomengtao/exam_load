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

7. **生成文件请勿直接修改**
    - 生成的文件（如 controller、service、biz、model、validator 等）请勿直接修改。
    - 因为每次重新生成都会被覆盖，所有修改应通过模板文件实现。
    - 仅允许在临时调试或测试时短暂修改，正式代码请回归模板。

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
  - 批量创建：`POST /api/role`
  - 批量更新：`PUT /api/role`
  - 删除：`DELETE /api/role`（支持批量删除，body 传 ids）
  - 查询（带搜索）：`GET /api/role/list?keyword=xxx&page=1&pageSize=10`

  查询（带搜索）路由参数示例：
  ```
  示例：
  GET /api/role/list?page=1&pageSize=10&amp;search[name]=admin&amp;filters[status]=active&amp;sort[created_at]=desc
  支持参数：
  - page / pageSize：分页
  - search[field]=value：模糊搜索
  - filters[field]=value：精确过滤
  - sort[field]=asc|desc：排序
  - select[]=字段名：指定返回字段
  ```

## 生成器执行命令

> 说明：`role` 仅为演示用表名，实际请替换为你要生成的表名。

生成器执行命令示例：
```bash
go run utils/gen/gen.go -table=role
```
> 注：表名通过命令行传入，其它参数自动推导。

## 最小第一步目标

- 在 controller.tpl 中添加批量创建接口路由和处理函数
- 在 service.tpl 中添加批量创建服务方法
- 在 biz.tpl 中添加批量创建业务逻辑
- 在 model.tpl 中添加基础字段验证标签
- 创建 validator.tpl 和 validator.go 模板，实现基础验证逻辑
- 创建单元测试文件 `app/controllers/{{.VarName}}_controller_test.go`，验证最小目标功能


## 错误处理注意事项

- 在返回 JSON 时，不要直接返回 []error 类型，应将其转换为 []map[string]interface{} 或 []string，否则 JSON 编码会失败，返回 [{}]。
- 示例：在 BatchCreateHandler 中，将 errors 转换为 errorMessages，确保返回格式正确。

### RESTful 格式统一返回函数（建议使用）

项目统一使用 `utils/response.go` 中定义的 `Success`、`Error` 等函数来构建响应数据结构。

- 使用方法：
  - 成功响应：`return utils.Success(c, data)`
  - 错误响应：`return utils.Error(c, "错误信息")`
  - 分页响应：`return utils.PageSuccess(c, list, total)`

- 要求：
  - 所有 Controller 层应优先使用这些函数封装输出。
  - Service 层或 Biz 层若需向 Controller 报错，应返回标准 error 或 error 数组，由 Controller 再格式化。
  - 所有生成模板中应使用统一的 `utils.Success` / `utils.Error`。

---

**本文档是开发规则和协作的唯一参考依据。** 

## 生成器可升级性设计原则（Generator Upgradeability）

为了确保代码生成器在后期迭代中仍然具备可维护性和可升级性，避免因为早期设计缺陷导致大量已生成代码无法更新，特制定如下设计原则：

### 🧱 各类骨架层职责说明与限制

| 骨架层文件名 | 允许行为 | 禁止行为 | 说明 |
|--------------|----------|-----------|------|
| controller_skeleton.go | ✅ 封装 RESTful 错误格式、参数处理通用函数 | ❌ 写死具体表名或业务逻辑 | 适合统一格式封装逻辑，如错误返回、参数提取 |
| biz_skeleton.go | ✅ 编排 service 调用、轻量业务判断模板 | ❌ 写死字段判断（如 if item.Type == "X"） | 负责构建业务处理框架，保持中性和通用 |
| service_skeleton.go | ✅ 提供数据库字段遍历、数据映射框架 | ❌ 封装响应格式 ❌ 写事务处理 ❌ 写逻辑分支 | 仅负责数据读写框架，不介入业务流程 |

### ✅ 核心原则：生成文件应与生成器逻辑解耦
- 所有生成的代码文件（如 controller、service、biz、skeleton 等）在结构上应**独立完整**，不应依赖生成器中额外的函数或工具。
- 避免生成的文件中引入 generator 内部专用函数（如 camelCase、snakeCase 等工具函数），应在生成前处理好并写入最终代码。

### 📦 公共函数使用限制
公共函数按照作用域分为两类：

#### 1. 生成器专用函数（如命名转换函数）
- 示例：`camelCase`, `snakeCase` 等仅服务于模板渲染的函数。
- 使用策略：仅允许在生成过程调用（如模板执行），不允许生成后的代码中引用。

#### 2. 业务通用函数（如日志、数据库、配置工具）
- 示例：`logger.Infof`、`db.Conn()`、`config.Get()` 等属于业务通用范围。
- 使用策略：可在所有生成的文件中使用，并在项目全局维护。


### 🚫 避免的问题
- ❌ 生成文件引用生成器内未导出的 helper 工具包。
- ❌ 生成器变动导致历史生成文件报错（如函数消失、命名方式改变）。
- ❌ 骨架层中包含业务判断、事务控制、错误格式化等逻辑。

### ❗ 骨架层与业务/服务职责边界补充

为确保职责清晰与文件可独立维护，必须遵守以下规则：

#### 骨架层（Skeleton）中不允许的内容：
- ❌ 不允许编写任何业务判断逻辑（如 if item.Type == "x"）。
- ❌ 不允许写事务控制逻辑（如 tx.Begin()、tx.Commit()）。
- ❌ 不允许引入外部 helper 或 toolkit（如 custom error 处理）。
- ✅ 仅允许提供基础的结构性代码（如接口函数的空实现、字段遍历框架）。

#### 服务层（Service）中不允许的内容：
- ❌ 不允许与具体数据库事务深度绑定逻辑（应下沉到 skeleton 或 data 层）。
- ❌ 不允许处理通用格式化输出逻辑（如 RESTful 错误结构）。
- ❌ 不允许写入复杂 if-else 分支业务流，避免与 biz 层职责冲突。
- ✅ 应专注于数据交互、输入输出结构组织、验证调度等。

#### 业务层（Biz）中不允许的内容：
- ❌ 不允许直接操作数据库或事务（应通过 service 调用实现）。
- ❌ 不允许直接引用 controller 或 route 相关逻辑。
- ✅ 专注于业务规则实现、调用 service 完成复合逻辑。

**备注：如确需跨层调用逻辑，需在文档中说明并由负责人批准。**

### 🔄 升级策略建议
- 若生成器模板升级，需要明确影响范围，优先做到向后兼容。
- 已生成文件应提供可选“仅重建骨架层”的升级模式，保留业务层定制代码。
- 模板应避免重复生成 service/controller，改为手动引导开发者引用新的骨架接口。

该原则保障生成器可持续演进，避免成为“死神工具”，提升工程健壮性与长远维护能力。