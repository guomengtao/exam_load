## 自动读取路由并写入数据库

本系统计划支持点击按钮自动扫描项目中的 `router.go` 或相关路由定义文件，提取所有注册的 API 路径与方法（如 GET、POST、PUT、DELETE），并将其存储到 MySQL 数据库的 `route_status` 表中，作为统一管理的数据来源。

### 自动读取方式建议：
- 使用 `go/parser` 解析 Go 文件；
- 查找形如 `r.GET("/path", ...)` 的注册语句；
- 抽取路径、HTTP 方法、注释说明等元数据；
- 后端提供 `/api/route/refresh` 接口由前端触发按钮调用。

除了通过静态分析 `r.GET("/path", ...)` 这样的注册语句以外，Gin 框架提供了一个更简洁的方法：调用 `r.Routes()` 方法可列出当前所有已注册的路由，返回一个结构体数组，包含路径、方法、处理函数等信息。

### 示例代码：

```go
func PrintAllRoutes(r *gin.Engine) {
    for _, route := range r.Routes() {
        fmt.Printf("Method: %-6s Path: %-30s Handler: %s\n", route.Method, route.Path, route.Handler)
    }
}
```

此方法适合于多个控制器文件分散注册路由的情况，便于统一汇总、自动存入数据库。

---

## 同步按钮行为与“未发现”标记

前端点击“同步路由”按钮将触发 `/api/route/refresh` 接口，该接口应执行以下逻辑：

1. 获取当前所有注册路由（通过 `r.Routes()`）；
2. 将新发现的路由插入数据库；
3. 对比数据库中已有的路由记录，若某条路由在本次扫描中未出现：
   - 将该路由状态更新为 `missing`（未发现）；
   - 可选项：提示“该接口可能已被删除”。

### 增量更新策略：

- 若接口存在但字段有更新（如 handler 变化），则更新对应字段；
- 若新接口未出现在数据库中，则新增；
- 若旧接口在当前注册路由中未匹配，则标记为 `missing` 或 `deleted`;

这种增量策略可实现路由状态的长期追踪与维护。

---

## 路由暂停功能设计与探讨

### 是否有意义？
是有意义的，尤其在以下场景中：
- 某接口已废弃，但暂时不能删除；
- 某接口逻辑正在重构，暂不提供服务；
- 需要临时屏蔽某接口用于维护、灰度发布等。

### 实现方式建议：
- **非代码层注释**（不推荐）：直接注释掉 `router.go` 中注册语句容易遗漏，且不易管理；
- **状态字段控制**（推荐）：
  - `route_status` 表中添加 `status` 字段，如 `active`, `paused`, `deprecated`；
  - 请求中间件判断路由状态，如果为 `paused`，统一返回 `"接口暂停服务"` 的响应；
  - 前端页面或接口文档也可以标注“暂停中”状态，避免对接误用。

### 示例中间件逻辑：

```go
func RouteStatusMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        path := r.URL.Path
        if IsRoutePaused(path) {
            http.Error(w, "该接口已暂停服务", http.StatusServiceUnavailable)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

---
## 路由状态表结构设计（SQL 示例）

```sql
CREATE TABLE route_status (
    id INT AUTO_INCREMENT PRIMARY KEY,
    method VARCHAR(10) NOT NULL,               -- GET, POST, PUT, DELETE
    path VARCHAR(255) NOT NULL,                -- 路由路径，如 /api/user/info
    handler VARCHAR(255),                      -- Gin handler 函数名
    status ENUM('active', 'paused', 'missing', 'deprecated') DEFAULT 'active',  -- 当前状态
    group_name VARCHAR(100),                   -- 模块或分组名称（用于树结构）
    owner VARCHAR(100),                        -- 接口维护人或负责人
    allowed_roles VARCHAR(255),                -- 允许访问的角色，如 "admin,user"
    is_private BOOLEAN DEFAULT FALSE,          -- 是否私有接口
    visit_count INT DEFAULT 0,                 -- 访问次数（可选）
    last_visited_at DATETIME,                  -- 最近访问时间（可选）
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
```

---

## 什么是 panic？

在 Go 语言中，`panic` 表示程序遇到了**无法恢复的严重错误**，比如：

- 空指针访问；
- 超出数组下标；
- 手动调用 `panic("错误信息")`；
- 类型断言失败（未使用安全写法）等。

一旦触发 `panic`，当前函数会立即中断执行，逐层向上传递，**如果没有 `recover()` 捕捉，程序会直接崩溃退出。**

示例：

```go
func mayPanic() {
    panic("这里发生了严重错误！")
}

func main() {
    mayPanic()
    fmt.Println("不会执行到这里，因为 panic 崩溃了程序")
}
```

✅ 正确做法：使用 `defer` + `recover()` 捕捉 panic，避免程序崩溃：

```go
func safeHandler() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("捕捉到 panic：", err)
        }
    }()
    mayPanic()
}
```