package controllers

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "net/http"
    "gin-go-test/app/models"
    "strconv"
)

// RouteController 路由管理控制器
type RouteController struct {
    db     *gorm.DB
    engine *gin.Engine
}

// NewRouteController 创建路由控制器实例
func NewRouteController(db *gorm.DB, engine *gin.Engine) *RouteController {
    return &RouteController{db: db, engine: engine}
}

// RefreshRoutes 刷新路由列表
// @Summary 刷新路由列表
// @Description 扫描并更新所有注册的路由
// @Tags 路由管理
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /api/route/refresh [post]
func (c *RouteController) RefreshRoutes(ctx *gin.Context) {
    // 获取所有注册的路由
    routes := c.engine.Routes()
    
    // 开始事务
    tx := c.db.Begin()
    
    // 标记所有现有路由为missing
    if err := tx.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&models.RouteStatus{}).Update("status", "missing").Error; err != nil {
        tx.Rollback()
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新路由状态失败"})
        return
    }
    
    // 更新或创建路由记录
    for _, route := range routes {
        var routeStatus models.RouteStatus
        result := tx.Where("path = ? AND method = ?", route.Path, route.Method).First(&routeStatus)
        
        if result.Error != nil {
            // 新路由，创建记录
            routeStatus = models.RouteStatus{
                Method:  route.Method,
                Path:    route.Path,
                Handler: route.Handler,
                Status:  "active",
            }
            if err := tx.Create(&routeStatus).Error; err != nil {
                tx.Rollback()
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": "创建路由记录失败"})
                return
            }
        } else {
            // 更新现有路由
            routeStatus.Status = "active"
            if err := tx.Save(&routeStatus).Error; err != nil {
                tx.Rollback()
                ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新路由记录失败"})
                return
            }
        }
    }
    
    // 提交事务
    if err := tx.Commit().Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "提交事务失败"})
        return
    }
    
    ctx.JSON(http.StatusOK, gin.H{
        "message": "路由刷新成功",
        "count":   len(routes),
    })
}

// GetRoutes 获取路由列表
// @Summary 获取路由列表
// @Description 获取所有路由的状态信息
// @Tags 路由管理
// @Accept json
// @Produce json
// @Param page query int false "页码"
// @Param size query int false "每页数量"
// @Param status query string false "状态筛选"
// @Param group query string false "分组筛选"
// @Success 200 {object} map[string]interface{}
// @Router /api/route/list [get]
func (c *RouteController) GetRoutes(ctx *gin.Context) {
    pageStr := ctx.DefaultQuery("page", "1")
    sizeStr := ctx.DefaultQuery("size", "100")
    status := ctx.Query("status")
    group := ctx.Query("group")

    page, _ := strconv.Atoi(pageStr)
    size, _ := strconv.Atoi(sizeStr)

    query := c.db.Model(&models.RouteStatus{})
    
    if status != "" {
        query = query.Where("status = ?", status)
    }
    if group != "" {
        query = query.Where("group_name = ?", group)
    }

    var total int64
    query.Count(&total)

    var routes []models.RouteStatus
    if err := query.Offset((page - 1) * size).Limit(size).Find(&routes).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取路由列表失败"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{
        "total": total,
        "list":  routes,
    })
}

// GetRouteDetail 获取路由详情
// @Summary 获取路由详情
// @Description 获取单个路由的详细信息
// @Tags 路由管理
// @Accept json
// @Produce json
// @Param id path int true "路由ID"
// @Success 200 {object} models.RouteStatus
// @Router /api/route/{id} [get]
func (c *RouteController) GetRouteDetail(ctx *gin.Context) {
    id := ctx.Param("id")
    var route models.RouteStatus
    if err := c.db.First(&route, id).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "路由不存在"})
        return
    }
    ctx.JSON(http.StatusOK, route)
}

// UpdateRouteStatus 更新路由状态
// @Summary 更新路由状态
// @Description 更新指定路由的状态
// @Tags 路由管理
// @Accept json
// @Produce json
// @Param id path int true "路由ID"
// @Param status body string true "新状态"
// @Success 200 {object} models.RouteStatus
// @Router /api/route/{id}/status [put]
func (c *RouteController) UpdateRouteStatus(ctx *gin.Context) {
    id := ctx.Param("id")
    var req struct {
        Status string `json:"status" binding:"required,oneof=active paused deprecated"`
    }
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的状态值"})
        return
    }
    
    var route models.RouteStatus
    if err := c.db.First(&route, id).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "路由不存在"})
        return
    }
    
    route.Status = req.Status
    if err := c.db.Save(&route).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新状态失败"})
        return
    }
    
    ctx.JSON(http.StatusOK, route)
}

// GetRouteGroups 获取路由分组
// @Summary 获取路由分组
// @Description 获取所有路由分组信息
// @Tags 路由管理
// @Accept json
// @Produce json
// @Success 200 {array} map[string]interface{}
// @Router /api/route/groups [get]
func (c *RouteController) GetRouteGroups(ctx *gin.Context) {
    var groups []struct {
        GroupName   string `json:"group_name"`
        RouteCount  int64  `json:"route_count"`
        Owner       string `json:"owner"`
    }
    
    if err := c.db.Model(&models.RouteStatus{}).
        Select("group_name, count(*) as route_count, owner").
        Group("group_name, owner").
        Find(&groups).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "获取分组信息失败"})
        return
    }
    
    ctx.JSON(http.StatusOK, groups)
}

// UpdateRouteGroup 更新路由分组
// @Summary 更新路由分组
// @Description 更新路由的分组信息
// @Tags 路由管理
// @Accept json
// @Produce json
// @Param id path int true "路由ID"
// @Param group body object true "分组信息"
// @Success 200 {object} models.RouteStatus
// @Router /api/route/{id}/group [put]
func (c *RouteController) UpdateRouteGroup(ctx *gin.Context) {
    id := ctx.Param("id")
    var req struct {
        GroupName string `json:"group_name" binding:"required"`
        Owner     string `json:"owner" binding:"required"`
    }
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的分组信息"})
        return
    }
    
    var route models.RouteStatus
    if err := c.db.First(&route, id).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "路由不存在"})
        return
    }
    
    route.GroupName = req.GroupName
    route.Owner = req.Owner
    if err := c.db.Save(&route).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新分组信息失败"})
        return
    }
    
    ctx.JSON(http.StatusOK, route)
}

// UpdateRoutePermission 更新路由权限
// @Summary 更新路由权限
// @Description 更新路由的权限设置
// @Tags 路由管理
// @Accept json
// @Produce json
// @Param id path int true "路由ID"
// @Param permission body object true "权限信息"
// @Success 200 {object} models.RouteStatus
// @Router /api/route/{id}/permission [put]
func (c *RouteController) UpdateRoutePermission(ctx *gin.Context) {
    id := ctx.Param("id")
    var req struct {
        AllowedRoles string `json:"allowed_roles"`
        IsPrivate    bool   `json:"is_private"`
    }
    
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "无效的权限信息"})
        return
    }
    
    var route models.RouteStatus
    if err := c.db.First(&route, id).Error; err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": "路由不存在"})
        return
    }
    
    route.AllowedRoles = req.AllowedRoles
    route.IsPrivate = req.IsPrivate
    if err := c.db.Save(&route).Error; err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "更新权限信息失败"})
        return
    }
    
    ctx.JSON(http.StatusOK, route)
}

// GetRouteStats 获取路由统计
// @Summary 获取路由统计
// @Description 获取路由访问统计信息
// @Tags 路由管理
// @Accept json
// @Produce json
// @Param start_time query string false "开始时间"
// @Param end_time query string false "结束时间"
// @Param group query string false "分组"
// @Success 200 {object} map[string]interface{}
// @Router /api/route/stats [get]
func (c *RouteController) GetRouteStats(ctx *gin.Context) {
    startTime := ctx.Query("start_time")
    endTime := ctx.Query("end_time")
    group := ctx.Query("group")
    
    query := c.db.Model(&models.RouteStatus{})
    
    if startTime != "" && endTime != "" {
        query = query.Where("last_visited_at BETWEEN ? AND ?", startTime, endTime)
    }
    if group != "" {
        query = query.Where("group_name = ?", group)
    }
    
    var stats struct {
        TotalVisits      int64 `json:"total_visits"`
        ActiveRoutes     int64 `json:"active_routes"`
        PausedRoutes     int64 `json:"paused_routes"`
        DeprecatedRoutes int64 `json:"deprecated_routes"`
    }
    
    // 获取总访问量
    var totalVisits int64
    query.Select("sum(visit_count)").Scan(&totalVisits)
    stats.TotalVisits = totalVisits
    
    // 获取各状态路由数量
    c.db.Model(&models.RouteStatus{}).Where("status = ?", "active").Count(&stats.ActiveRoutes)
    c.db.Model(&models.RouteStatus{}).Where("status = ?", "paused").Count(&stats.PausedRoutes)
    c.db.Model(&models.RouteStatus{}).Where("status = ?", "deprecated").Count(&stats.DeprecatedRoutes)
    
    // 获取分组统计
    var groupStats []struct {
        GroupName   string `json:"group_name"`
        VisitCount  int64  `json:"visit_count"`
        RouteCount  int64  `json:"route_count"`
    }
    
    c.db.Model(&models.RouteStatus{}).
        Select("group_name, sum(visit_count) as visit_count, count(*) as route_count").
        Group("group_name").
        Find(&groupStats)
    
    ctx.JSON(http.StatusOK, gin.H{
        "total_visits": stats.TotalVisits,
        "active_routes": stats.ActiveRoutes,
        "paused_routes": stats.PausedRoutes,
        "deprecated_routes": stats.DeprecatedRoutes,
        "group_stats": groupStats,
    })
} 