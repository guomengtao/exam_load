/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_cate_contens` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cate` int(11) NOT NULL DEFAULT '0' COMMENT '所属问题',
  `title` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '问题名称',
  `type` int(1) NOT NULL DEFAULT '0' COMMENT '问题类型0单选1多选',
  `number` int(1) NOT NULL DEFAULT '0' COMMENT '问题序号',
  `score` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '分数',
  `correct` varchar(126) DEFAULT NULL COMMENT '正确答案',
  `contens` text CHARACTER SET utf8,
  `answer1` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '问题答案1',
  `answer2` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案2',
  `answer3` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案3',
  `answer4` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案4',
  `answer5` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案5',
  `answer6` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案6',
  `Copy_of_answer2` varchar(255) CHARACTER SET utf8 DEFAULT '' COMMENT '问题答案2',
  `time` int(10) NOT NULL DEFAULT '1' COMMENT '发布时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=614 DEFAULT CHARSET=gbk;
/*!40101 SET character_set_client = @saved_cs_client */;
