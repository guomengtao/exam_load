/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_cate` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `catename` varchar(200) DEFAULT '' COMMENT '栏目名称',
  `username` varchar(150) DEFAULT '' COMMENT '所属管理员',
  `desc` text COMMENT '栏目描述',
  `type` tinyint(1) DEFAULT '0' COMMENT '栏目类型0：列表 1：留言',
  `title` varchar(255) DEFAULT '' COMMENT '问卷模板标题',
  `score` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '分数',
  `time` int(10) NOT NULL COMMENT '发布时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=212 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;
