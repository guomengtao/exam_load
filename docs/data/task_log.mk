/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `task_log` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键',
  `task_id` int(10) unsigned NOT NULL COMMENT '关联任务 ID',
  `task_name` varchar(64) NOT NULL COMMENT '任务名称冗余',
  `status` enum('success','failed') NOT NULL COMMENT '执行状态',
  `message` text COMMENT '日志信息或错误信息',
  `started_at` datetime DEFAULT NULL COMMENT '任务开始时间',
  `ended_at` datetime DEFAULT NULL COMMENT '任务结束时间',
  `duration` int(10) unsigned DEFAULT NULL COMMENT '执行时长（秒）',
  PRIMARY KEY (`id`),
  KEY `task_id` (`task_id`),
  KEY `started_at` (`started_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务执行日志表';
/*!40101 SET character_set_client = @saved_cs_client */;
