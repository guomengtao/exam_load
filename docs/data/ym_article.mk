/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_article` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL COMMENT '标题',
  `username` varchar(150) DEFAULT '' COMMENT '所属用户',
  `desc` varchar(255) DEFAULT '' COMMENT '描述',
  `pic` varchar(100) DEFAULT NULL COMMENT '缩略图',
  `content` text COMMENT '内容',
  `click` mediumint(9) DEFAULT '0' COMMENT '点击量',
  `cateid` mediumint(9) DEFAULT NULL,
  `time` int(10) DEFAULT '0' COMMENT '发布时间',
  `close` mediumint(9) DEFAULT '0' COMMENT '打开关闭',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=242 DEFAULT CHARSET=utf8 COMMENT='试卷';
/*!40101 SET character_set_client = @saved_cs_client */;
