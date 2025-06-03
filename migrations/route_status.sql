-- Create tm_route_status table
CREATE TABLE IF NOT EXISTS tm_route_status (
    id INT AUTO_INCREMENT PRIMARY KEY COMMENT '主键ID',
    method VARCHAR(10) NOT NULL COMMENT 'HTTP请求方法(GET/POST/PUT/DELETE)',
    path VARCHAR(255) NOT NULL COMMENT 'API路由路径',
    handler VARCHAR(255) COMMENT '处理函数名称',
    status ENUM('active', 'paused', 'missing', 'deprecated') DEFAULT 'active' COMMENT '路由状态(active:活跃/paused:暂停/missing:未发现/deprecated:废弃)',
    group_name VARCHAR(100) COMMENT '路由分组名称',
    owner VARCHAR(100) COMMENT '接口负责人',
    allowed_roles VARCHAR(255) COMMENT '允许访问的角色列表',
    is_private BOOLEAN DEFAULT FALSE COMMENT '是否为私有接口',
    visit_count INT DEFAULT 0 COMMENT '访问次数统计',
    last_visited_at DATETIME COMMENT '最后访问时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    UNIQUE KEY unique_route (method, path) COMMENT '路由唯一性约束'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='路由状态管理表'; 