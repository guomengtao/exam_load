package models

import (
    "time"
    "gorm.io/gorm"
)

// RouteStatus 路由状态模型
type RouteStatus struct {
    ID            uint      `gorm:"primarykey" json:"id"`                              // 主键ID
    Method        string    `gorm:"size:10;not null" json:"method"`                    // HTTP请求方法
    Path          string    `gorm:"size:255;not null" json:"path"`                     // API路由路径
    Handler       string    `gorm:"size:255" json:"handler"`                           // 处理函数名称
    Status        string    `gorm:"type:varchar(20);default:'active'" json:"status"` // Route status: active, paused, missing, deprecated
    GroupName     string    `gorm:"size:100" json:"group_name"`                        // 路由分组名称
    Owner         string    `gorm:"size:100" json:"owner"`                             // 接口负责人
    AllowedRoles  string    `gorm:"size:255" json:"allowed_roles"`                     // 允许访问的角色列表
    IsPrivate     bool      `gorm:"default:false" json:"is_private"`                   // 是否为私有接口
    VisitCount    int       `gorm:"default:0" json:"visit_count"`                      // 访问次数统计
    LastVisitedAt time.Time `json:"last_visited_at"`                                   // 最后访问时间
    UpdatedAt     time.Time `json:"updated_at"`                                        // 更新时间
}

// TableName 指定表名
func (RouteStatus) TableName() string {
    return "tm_route_status"
}

// BeforeCreate 创建前的钩子
func (r *RouteStatus) BeforeCreate(tx *gorm.DB) error {
    r.UpdatedAt = time.Now()
    return nil
}

// BeforeUpdate 更新前的钩子
func (r *RouteStatus) BeforeUpdate(tx *gorm.DB) error {
    r.UpdatedAt = time.Now()
    return nil
} 