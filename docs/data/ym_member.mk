/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_member` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL COMMENT '姓名',
  `username` varchar(50) DEFAULT NULL COMMENT '所属管理员ID',
  `province` varchar(30) DEFAULT '' COMMENT '省',
  `city` varchar(50) DEFAULT NULL COMMENT '市',
  `area` varchar(150) DEFAULT '' COMMENT '县/区',
  `school` varchar(30) DEFAULT '' COMMENT '学校',
  `grade` varchar(50) DEFAULT '' COMMENT '年级',
  `class` varchar(22) DEFAULT NULL COMMENT '班级',
  `serialnumber` varchar(100) DEFAULT NULL COMMENT '学号',
  `telephone` varchar(13) DEFAULT NULL COMMENT '电话',
  `type` int(11) DEFAULT NULL COMMENT '类型（老师/家长/学生）',
  `level` int(11) DEFAULT NULL COMMENT '级别（小/初/高）',
  `number` int(11) DEFAULT '0' COMMENT '编号',
  `memberid` varchar(50) DEFAULT NULL,
  `pic` varchar(100) DEFAULT '' COMMENT '头像',
  `timee` mediumint(9) DEFAULT NULL COMMENT '添加时间',
  `time` int(11) DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2042 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;
