/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_answer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cate` int(11) NOT NULL DEFAULT '0' COMMENT '所属试卷',
  `title` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '问题名称',
  `type` int(1) NOT NULL DEFAULT '0' COMMENT '问题类型',
  `time` int(10) DEFAULT NULL COMMENT '提交时间',
  `correct` varchar(20) DEFAULT NULL COMMENT '回答汇总',
  `score` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '得分',
  `answerid` int(11) DEFAULT '0' COMMENT '所属问题id',
  `answer1` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案1',
  `answer2` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案2',
  `answer3` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案3',
  `answer4` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案4',
  `answer5` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案5',
  `answer6` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '问题答案6',
  `member` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '用户',
  `name` varchar(255) CHARACTER SET utf8 DEFAULT NULL COMMENT '填写人姓名',
  `serialnumber` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '学号',
  `memberid` varchar(255) DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=746191 DEFAULT CHARSET=gbk ROW_FORMAT=DYNAMIC;
/*!40101 SET character_set_client = @saved_cs_client */;
