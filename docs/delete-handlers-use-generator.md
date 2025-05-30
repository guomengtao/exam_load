


# 方案：delete-handlers-use-generator

## 🎯 最终目标

彻底删除 `handlers` 文件夹，将其中的接口统一升级为由生成器自动生成的接口逻辑，保持架构一致性，提高可维护性。

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
