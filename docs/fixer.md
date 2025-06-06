

# Fixer 自动修正工具方案文档

## 1. 目的
Fixer 是一个用于代码自动修正的辅助工具，主要目标是在开发过程中提升代码质量，减少低级错误，保证项目结构和依赖管理的规范性。

## 2. 功能概述
- 自动修正 Go 代码中的 package 名称，使其与文件夹名称保持一致。
- 自动检测并删除未使用的 import 导入包，保持 import 区域整洁。
- 不进行自动补全 import，也不做复杂的依赖推断，避免实现复杂度和不稳定性。
- 作为编译和测试之前的辅助工具，作为“锦上添花”的辅助，不代替人工的代码审核。

## 3. 设计原则
- **稳定优先**：只修正明显且简单的问题，避免风险和破坏已有功能。
- **不过度修正**：不进行自动补全，避免复杂的依赖分析导致误判。
- **易用集成**：可集成进打包、测试的自动化脚本流程，作为前置检查和修正步骤。
- **明确规范**：严格按照项目的开发规范执行，如 package 与文件夹名必须一致。

## 4. 核心实现思路
- 解析 Go 源码文件的 AST，定位 package 声明和 import 区域。
- 对 package 声明做名称检测，如果与所在目录名不符，自动修改。
- 分析 import 区域，找出未使用的包，删除对应的 import 语句。
- 保留带下划线或别名的 import，不删除注释掉的 import。
- 支持命令行参数，指定扫描的目录或文件，支持增量检查（比如只检查改动文件）。

## 5. 使用场景
- 在本地开发时，作为开发者的辅助工具，减少小错误反复出现。
- 在 CI/CD 流程中，作为自动修正环节，保证提交的代码符合规范。
- 在打包、测试前强制执行，提升整体项目的稳定性。

## 6. 未来扩展可能
- 扩展到对 struct 结构体定义、接口定义的规范检查和自动修正。
- 集成常用代码风格检查，自动格式化部分代码结构。
- 对更复杂的依赖关系进行安全提示，但不做自动补全。
- 增加更多修正规则，但都必须遵循稳定和可控的原则。

## 7. 风险与注意点
- 自动删除 import 必须小心处理，避免误删导致编译失败。
- 修正 package 名称必须确保对应目录名准确，防止误改。
- 过度自动补全和删除可能导致代码逻辑错误和不稳定，当前版本避免。
- 任何自动修正都应在提交之前执行，保证代码稳定后再合入。

---

> 以上方案旨在提升开发效率和代码规范，避免重复低级错误，给开发者带来更流畅的编码体验。
