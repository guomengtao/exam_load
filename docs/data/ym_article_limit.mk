/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_article_limit` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `articleid` varchar(255) DEFAULT '' COMMENT '所属问卷',
  `limit` text COMMENT '允许的用户记录',
  `title` varchar(30) DEFAULT NULL COMMENT '标题',
  `username` varchar(150) DEFAULT '' COMMENT '所属用户',
  `pic` varchar(100) DEFAULT NULL COMMENT '缩略图',
  `type` varchar(100) DEFAULT '0',
  `click` mediumint(9) DEFAULT '0' COMMENT '点击量',
  `cateid` mediumint(9) DEFAULT NULL,
  `time` int(10) DEFAULT NULL COMMENT '发布时间',
  `close` mediumint(9) DEFAULT '0' COMMENT '打开关闭',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=57 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;
