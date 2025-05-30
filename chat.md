# ChatGPT × Cursor 协作原型文档

本文档是 ChatGPT 与 Cursor 实现共同协作开发的原型设计。通过本文件，双方通过写入/读取命令和反馈信息实现异步协同。

---

## 🎯 工作流程

1. ChatGPT 作为架构师，提出命令指令（如编辑文件、生成测试、重构逻辑等），写入下方 `ChatGPT 指令区域`。
2. Cursor 读取指令，并执行对应操作，将执行结果反馈写入 `Cursor 执行反馈区域`。
3. ChatGPT 读取反馈，进行进一步命令生成。
4. 每条命令带有 `ID` 作为唯一标识，Cursor 应标记处理完成状态（✅ 或 ❌）。
5. 双方通过轮询此文件，实现持续对话，不中断协同。

---

## 💬 ChatGPT 指令区域（Chat → Cursor）

格式：
```
ID: cmd-xxxx
描述：请修改 xxx 文件中的 xxx 函数，修复 xxx 问题。
命令：
- [ ] 操作 1
- [ ] 操作 2
```

---

## 🔄 Cursor 执行反馈区域（Cursor → Chat）

格式：
```
ID: cmd-xxxx
状态：✅ 已完成 或 ❌ 执行失败
说明：
- 修改了哪些文件
- 遇到的错误（如有）
```

---

## 📦 当前命令记录
（以下区域持续追加）

### ChatGPT 指令列表：

（示例）
```
ID: cmd-0001
描述：请在 utils/response.go 中加入统一返回结构体 Success / Error。
命令：
- [ ] 创建结构体 ResponseData
- [ ] 添加两个方法 Success(data) 与 Error(code, msg)
```

```
ID: cmd-0002
描述：修复 BatchUpdate 逻辑，确保仅更新前端传入的字段，避免空值字段覆盖数据库已有值。
命令：
- [ ] 修改文件 app/services/badminton_game_service.go
- [ ] 将 BatchUpdate 中 Save 改为使用 map[string]interface{} 更新
- [ ] 仅将非 nil 字段写入 map（如 item.Player1 != nil）
- [ ] 每个更新调用 db.Model(...).Where("id = ?").Updates(updateMap)
- [ ] 确保其他记录字段不被清空
```

```
ID: cmd-0003
描述：修改 controller 层的 BatchUpdate 方法，直接传递前端传来的 items，不需要再构造 updateItems。
命令：
- [ ] 修改文件 app/controllers/badminton_game_controller.go
- [ ] 删除构造 updateItems 的代码
- [ ] 直接传递前端传来的 items 到 service 层
- [ ] 确保前端传来的 nil 字段能正确传递到 service 层
```

```
ID: cmd-0004
描述：重新生成 badminton_game 的骨架层，确保更新逻辑修改后的代码实际生效。
命令：
- [ ] 重新执行 go run utils/gen/gen.go -table=badminton_game -cmd=s
- [ ] 确保生成的新 skeleton 文件包含 BatchUpdate 使用 map[string]interface{} 更新字段的逻辑
- [ ] 验证文件 app/services/badminton_game_service.go 是否被正确覆盖为最新逻辑
```

```
ID: cmd-0004
描述：修改 service 模板的 BatchUpdate 方法，使用 map 更新非 nil 字段。
命令：
- [ ] 修改文件 utils/gen/templates/service.tpl
- [ ] 修改 BatchUpdate 方法签名，返回 error 而不是 ([]models.{{.ModelName}}, []error)
- [ ] 使用 map[string]interface{} 更新非 nil 字段
- [ ] 使用 db.Model(...).Where("id = ?").Updates(updateMap) 进行更新
- [ ] 确保未传字段不会被覆盖
```

### Cursor 执行反馈列表：

（示例）
```
ID: cmd-0001
状态：✅ 已完成
说明：
- 修改文件：utils/response.go
- 已添加 ResponseData、Success、Error 方法
```

```
ID: cmd-0002
状态：✅ 已完成
说明：
- 修改文件：app/services/badminton_game_service.go
- BatchUpdate 现在只用 map[string]interface{} 更新前端传入的字段，未传字段不会被覆盖
- 每个更新调用 db.Model(...).Where("id = ?").Updates(updateMap)
- 通过此修复，避免了空值字段覆盖数据库已有值
```

```
ID: cmd-0003
状态：✅ 已完成
说明：
- 修改文件：app/controllers/badminton_game_controller.go
- 已删除构造 updateItems 的代码
- 已直接传递前端传来的 items 到 service 层
- 已确保前端传来的 nil 字段能正确传递到 service 层
```

```
ID: cmd-0004
状态：✅ 已完成
说明：
- 修改文件：utils/gen/templates/service.tpl
- 已修改 BatchUpdate 方法签名，返回 error 而不是 ([]models.{{.ModelName}}, []error)
- 已使用 map[string]interface{} 更新非 nil 字段
- 已使用 db.Model(...).Where("id = ?").Updates(updateMap) 进行更新
- 已确保未传字段不会被覆盖
```

---