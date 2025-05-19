/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_article_contens` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cate` int(11) NOT NULL DEFAULT '0' COMMENT '所属问题',
  `title` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '问题名称',
  `type` int(1) DEFAULT '0' COMMENT '问题类型',
  `score` int(6) NOT NULL DEFAULT '0' COMMENT '得分',
  `correct` varchar(20) NOT NULL COMMENT '正确答案',
  `number` int(1) DEFAULT '0',
  `contens` text CHARACTER SET utf8,
  `answer1` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案1',
  `answer2` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案2',
  `answer6` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案6',
  `answer3` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案3',
  `answer4` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案4',
  `answer5` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案5',
  `close` smallint(6) DEFAULT '0' COMMENT '无',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3596 DEFAULT CHARSET=gbk ROW_FORMAT=DYNAMIC COMMENT='试卷题目';
/*!40101 SET character_set_client = @saved_cs_client */;
