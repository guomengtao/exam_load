# Gin + Go 实战项目

## 一、项目介绍

该项目是一个基于 Gin 框架和 Go 语言开发的在线考试系统，包含用户管理、试题管理、考试管理等功能模块。

## 二、项目目录结构

```
.
├── main.go                      # 程序入口
├── db.go                        # 数据库连接
├── redis.go                     # Redis 连接
├── run.sh                       # 启动脚本
├── handlers
│   ├── user.go                  # 用户相关 API (后续计划实现)
│   ├── question.go              # 问题相关 API (后续计划实现)
│   ├── html_exam.go             # 生成试卷 HTML 页面 (后续计划实现)
│   └── html_report.go           # 生成成绩报告 HTML 页面 (后续计划实现)
├── templates                    # HTML 模板文件 (后续计划实现)
│   ├── exam.html
│   ├── report.html
│   └── ...
├── static                       # 静态页面 (后续计划实现)
│   ├── css
│   ├── js
│   └── images
├── utils
│   ├── html.go                  # HTML 生成工具 (后续计划实现)
│   ├── logger.go                # 日志封装 (后续计划实现)
│   └── ...
├── tasks                        # 定时任务模块 (后续计划实现)
│   └── ...
└── middleware
    └── auth.go                  # JWT 鉴权中间件 (后续计划实现)
```