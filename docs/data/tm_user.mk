/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tm_user` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL COMMENT '姓名',
  `admin_id` varchar(50) DEFAULT NULL COMMENT '所属管理员ID',
  `province` varchar(30) DEFAULT '' COMMENT '省',
  `city` varchar(50) DEFAULT NULL COMMENT '市',
  `area` varchar(150) DEFAULT '' COMMENT '县/区',
  `school_name` varchar(50) DEFAULT NULL COMMENT '学校名称',
  `grade_name` varchar(50) DEFAULT NULL COMMENT '年级',
  `class_name` varchar(30) DEFAULT NULL COMMENT '班级',
  `user_id` varchar(50) DEFAULT NULL COMMENT '学号（唯一）',
  `phone` varchar(20) DEFAULT NULL COMMENT '联系电话',
  `user_type` tinyint(4) DEFAULT NULL COMMENT '用户类型（1-老师，2-家长，3-学生）',
  `school_level` tinyint(4) DEFAULT NULL COMMENT '学段（1-小学，2-初中，3-高中）',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2042 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC COMMENT='用户表';
/*!40101 SET character_set_client = @saved_cs_client */;
