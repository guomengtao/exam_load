# 方案：delete-handlers-use-generator

## 🎯 最终目标

彻底删除 `handlers` 文件夹，将其中的接口统一升级为由生成器自动生成的接口逻辑，保持架构一致性，提高可维护性。

### 📐 分层拆分原则与约束

为了确保接口迁移后的代码架构合理、职责清晰，所有接口需遵循以下分层规范：

#### controller 层（控制器）

- 仅负责请求接收、参数绑定与响应输出。
- 不包含业务逻辑或数据库操作。
- 可以进行请求日志打印、简单权限判断，但不得包含业务 `if` 条件分支。

#### biz 层（业务逻辑）

- 承担核心逻辑判断与条件处理（如参数校验、数据对比、业务状态机等）。
- 如果接口中存在多层 `if` 或状态判断，**必须**迁移至 biz 层处理。
- 不负责数据库或缓存访问，仅调用 service 提供的能力。

#### service 层（服务层）

- 负责数据库（GORM）和缓存（Redis）的具体读写。
- **不允许出现复杂 `if` 判断或业务状态逻辑**，所有判断应提前由 biz 层完成。
- 封装常见数据库与 Redis 操作，提供通用方法。

#### Redis 操作规范

- 所有 Redis 调用必须通过统一工具包封装，位于 `utils/redis_helper.go`。
- 不允许在 controller 或 biz 中直接调用 Redis。
- Redis key 命名应规范（如 `prefix:module:id`），并配置超时时间。
- 尽量采用结构化数据（如 JSON 编码存储）以便调试。

#### 接口迁移拆分要素示例

- ❌ 错误：controller 中包含 `if 参数 == xxx` 判断。
- ✅ 正确：controller 调用 `biz.HandleXxx(params)`，所有判断在 biz 中完成。
- ❌ 错误：service 中对不同请求参数执行逻辑分支。
- ✅ 正确：service 仅暴露 `FindXxx`, `UpdateXxx`，不关心请求场景。

- ✅ 删除 `/handlers/*.go`
- 🔄 部分通用工具函数或上传接口可迁移到 `/controllers` 或 `/utils`
- 📦 接口迁移后统一由 `RegisterGeneratedRoutes(router)` 自动注册

---

## 🧭 项目评估

### ✅ 适合用生成器升级的模块（已具备模型、字段配置）

- exam_template
- exam_paper
- answer
- admin
- role
- user
- file_info
- teacher
- badminton_game

这些模型具备标准的 CURD 操作和结构，可以直接运行：

```bash
go run utils/gen/gen.go model_name
```

---

## ⚠️ 风险与注意事项

1. **接口路径变化**：
   - 如果前端使用的是老的 handlers 中定义的路径，替换为生成器逻辑可能会导致路径或字段名不兼容。
   - 解决方案：
     - 保持旧路由兼容性（手动映射）
     - 或要求前端重新对接（推荐长远方案）

2. **接口逻辑差异**：
   - 原 handlers 中接口可能存在特殊逻辑（如分页、Redis 缓存、过滤字段）。
   - 生成器仅提供通用逻辑，部分功能需手动扩展 `Skeleton` 层。

3. **批量接口为主**：
   - 当前系统接口大多为批量接口，如 `BatchCreate`, `BatchUpdate`, `BatchDelete`。
   - 生成器支持这些操作，但前端如传参格式变化，需同步更新调用方式。

---

## ✅ 建议落地路径

1. 用生成器重建所需模块的 CURD 接口
2. 用测试脚本确保生成接口功能完整
3. 删除 handlers 中的冗余老接口文件
4. routes 中全面切换为 `RegisterGeneratedRoutes`
5. 前端逐步适配新接口（推荐将字段结构对比写入文档）

---

## 📌 总结

删除 `handlers` 是本项目架构升级的核心步骤，预计中低风险，投入产出比高。配合生成器使用可实现接口统一、逻辑复用、维护简单。前端如能配合更新接口调用方式，将进一步降低技术债。

---

## 🚫 不需要生成器的接口迁移名单

以下接口功能简单、与业务模型解耦，适合**手动迁移**或**保留至 `/controllers` 目录**，无需生成器参与：

| 文件名 | 推荐迁移位置 | 理由 | 迁移难度 |
|--------|----------------|------|-----------|
| `dbinfo.go` | `/controllers/dbinfo_controller.go` | 查询数据库结构信息，非模型操作 | 🟢 低 |
| `status.go` | `/controllers/status_controller.go` | 系统状态接口，通常仅返回 JSON | 🟢 低 |
| `version.go` | `/controllers/version_controller.go` | 返回系统版本信息，逻辑极简 | 🟢 低 |
| `upload.go` | `/controllers/upload_controller.go` | 包含文件上传逻辑，需迁移至统一上传入口 | 🟡 中（涉及 multipart/form-data） |
| `source_check.go` | `/controllers/source_controller.go` | 数据源校验工具类接口 | 🟡 中 |
| `hello.go` | （可删除或合并至示例控制器） | 演示用途，无实际依赖 | 🟢 低 |

> 💡 提示：这些接口可以直接重命名/移动至 `controllers/`，并注册进 `routes.go` 中，不会影响业务逻辑或数据库结构。

# 
---

## 🎯 迁移目标定位与执行细节更新

本次迁移将遵循以下几点原则和策略，重点明确接口结构调整与职责划分：

### 🧭 迁移核心原则

1. **生成器作为辅助工具**：生成器是提升效率的辅助工具，**不是强制**。适合结构标准、字段已建模的接口；对复杂逻辑接口（如答题相关）不建议使用生成器。
2. **保证功能平移优先**：接口迁移的首要目标是功能等效，不添加无关优化逻辑，确保业务平稳过渡。
3. **鼓励分层重构**：即使不使用生成器，所有迁移接口也建议按 `controller → biz → service` 拆分，提升可测试性与逻辑清晰度。

### 📦 分层拆分细则（补充）

| 逻辑 | 建议所在层 | 说明 |
|------|-------------|------|
| 请求接收、参数绑定、响应格式 | controller | 不包含 if/业务判断 |
| 参数合法性判断、状态流转、核心 if/else 结构 | biz | 控制逻辑必须搬到这里 |
| GORM 操作、Redis 操作（通过工具包） | service | 禁止出现业务分支 |
| Redis get/set/del 封装 | utils/redis_helper.go | 所有缓存操作统一走工具包 |

### 🧱 示例对照

- ❌ 错误做法：controller 里存在 `if req.Type == "xxx"` 分支
- ✅ 正确做法：controller 调用 `biz.HandleXXX(req)`，判断逻辑在 biz 内处理
- ❌ 错误做法：service 里区分多个参数逻辑
- ✅ 正确做法：service 提供单一功能点方法，如 `SaveAnswer(redisKey, value)`，由 biz 控制调用流程

### 🔁 不用生成器但仍建议分层的接口（可独立迁移）

针对 `handlers/exam_answer.go` 等核心接口，**不使用生成器生成 Skeleton**，而是独立迁移为：

```
controllers/exam_answer_controller.go
biz/exam_answer_biz.go
services/exam_answer_service.go
```

此类接口为项目核心，逻辑复杂、影响范围大，采用手动迁移是更安全和灵活的方式。

> 🚨 注：答题流程接口涉及 Redis、答题提交、评分、记录，禁止引入生成器，避免功能损坏。