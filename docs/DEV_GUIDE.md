# 开发指南

**⚠️ 代码生成器使用规范：**

- 严禁直接修改任何生成器生成的文件（如 app/biz、app/services、utils/generated 等目录下的自动生成文件）。
- 如需调整生成内容，请修改模板（如 utils/gen/templates/）或生成器源码。
- 修改模板后，请通过 `go run utils/gen/gen.go -table=表名 -cmd=a` 等命令重新生成并覆盖。

// ... existing code ... 