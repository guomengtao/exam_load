/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tm_exam_papers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` char(36) DEFAULT NULL COMMENT '试卷UUID唯一标识',
  `template_id` bigint(20) DEFAULT NULL COMMENT '来源模板ID',
  `title` varchar(255) NOT NULL COMMENT '试卷标题',
  `description` text COMMENT '试卷描述',
  `cover_image` varchar(255) DEFAULT NULL COMMENT '封面图片URL',
  `total_score` int(11) NOT NULL DEFAULT '0' COMMENT '试卷总分',
  `questions` json NOT NULL COMMENT '题目列表（JSON数组）',
  `view_count` int(11) DEFAULT '0' COMMENT '浏览量',
  `category_id` bigint(20) DEFAULT NULL COMMENT '分类ID',
  `publish_time` int(10) DEFAULT '0' COMMENT '最近一次发布时间',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：1=正常，0=关闭',
  `creator` varchar(150) DEFAULT NULL COMMENT '创建人用户名',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间（NULL=未删除）',
  `time_limit` int(11) DEFAULT '0' COMMENT '考试时间限制（分钟），0表示无限时',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='试卷';
/*!40101 SET character_set_client = @saved_cs_client */;
