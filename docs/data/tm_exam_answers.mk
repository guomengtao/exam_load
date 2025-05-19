/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tm_exam_answers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主键ID，自增',
  `answers` json NOT NULL COMMENT '题目列表（JSON数组）',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间（NULL=未删除）',
  `uuid` varchar(255) NOT NULL COMMENT '唯一回答记录标识',
  `total_score` int(3) NOT NULL COMMENT '满分',
  `score` int(11) DEFAULT '0' COMMENT '用户实际得分',
  `username` varchar(100) NOT NULL COMMENT '学生姓名',
  `user_id` varchar(50) NOT NULL COMMENT '学生编号',
  `duration` int(11) DEFAULT '0' COMMENT '答题用时（单位：秒）',
  `exam_uuid` varchar(64) DEFAULT '' COMMENT '试卷UUID',
  `exam_id` int(3) DEFAULT NULL COMMENT '试卷ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18236 DEFAULT CHARSET=utf8mb4 COMMENT='学生作答表';
/*!40101 SET character_set_client = @saved_cs_client */;
