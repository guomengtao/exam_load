-- ----------------------------
-- Table structure for `ym_article`
-- ----------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
DROP TABLE IF EXISTS `ym_article`;
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
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;


-- ----------------------------
-- Sample data for `ym_article` (最早2条)
-- ----------------------------
INSERT INTO `ym_article` VALUES (52,'多个复选题目','','','','',183,83,1475704560,1),(57,'多个单选题目','','','','',62,84,1475915664,0),(58,'国庆去哪里玩','','大家国庆节去过哪些地方','','',34,80,1476065467,0),(69,'银川市中小学规范办学行为专项督导检查学生问卷','admin','请同学们根据问卷内容据实填写。',NULL,NULL,36,166,1680914290,0),(70,'银川市中小学规范办学行为专项督导检查学生问卷','admin','请同学们根据问卷内容据实填写。',NULL,NULL,0,166,1680914367,0),(71,'银川市中小学规范办学行为专项督导检查学生问卷','admin','请同学们根据问卷内容据实填写。',NULL,NULL,22,166,1680914430,0),(72,'动物知识测试','admin','',NULL,NULL,11,168,1709268327,0),(73,'植物知识测试','admin','你认识植物吗',NULL,NULL,19,169,1709268738,0),(84,'国家教育质量检测道德与法治四年级 测试卷 ','admin','',NULL,NULL,22,175,1709510406,0),(97,'中卫市初三音乐基础知识测试','中卫市教育局考试系统（演示版）','（测试版）',NULL,NULL,61,174,1709556204,0),(104,'国家教育质量检测数学模拟卷 （1）','admin',NULL,NULL,NULL,392,194,1715128744,0),(105,'国家教育质量检测心理健康模拟卷 （1）','admin',NULL,NULL,NULL,2949,196,1715150833,0),(146,'国家教育质量检测数学模拟卷 1','admin',NULL,NULL,NULL,2370,194,1715408463,0),(187,'国家教育质量检测数学模拟卷 2','admin',NULL,NULL,NULL,2941,198,1715777014,0),(205,'国家教育质量检测数学模拟卷 2','admin',NULL,NULL,NULL,1466,198,1715910032,0),(211,'实验','admin',NULL,NULL,NULL,8,199,1743678290,0),(212,'国家教育质量检测数学模拟卷 2（中宁县第九小学教育集团）','admin',NULL,NULL,NULL,47,198,1743943832,0),(213,'试验版1','admin',NULL,NULL,NULL,7,202,1743989581,0),(214,'试验版1','admin',NULL,NULL,NULL,1,202,1743993000,0),(215,'实验','admin',NULL,NULL,NULL,3,199,1743993011,0),(216,'一天上几节课','admin',NULL,NULL,NULL,11,183,1743993054,0),(217,'实验','admin','',NULL,NULL,1,204,1743993382,0),(218,'试验版1','admin',NULL,NULL,NULL,21,202,1744010988,0),(221,'试验版2','admin',NULL,NULL,NULL,87,202,1744096012,0),(222,'国家教育质量检测数学模拟卷 2（试验版）','admin',NULL,NULL,NULL,69,198,1744180414,0),(224,'国家教育质量检测英语模拟测试 1','admin',NULL,NULL,NULL,31680,206,1744680917,0),(226,'国家教育质量检测语文模拟测试 1','admin',NULL,NULL,NULL,21164,208,1744707054,0),(227,'国家教育质量检测艺术模拟测试 1','admin',NULL,NULL,NULL,27331,207,1744727389,0),(228,'国家教育质量检测数学模拟卷 （试验版）','admin',NULL,NULL,NULL,1139,198,1744774920,0),(229,'国家教育质量检测英语模拟测试 1（贺兰县第一小学）','admin',NULL,NULL,NULL,17,206,1745369768,1),(230,'国家教育质量检测英语模拟测试 1','admin',NULL,NULL,NULL,45,206,1745376127,0),(231,'实验','admin',NULL,NULL,NULL,0,199,1745753334,0),(232,'试验版1','admin','',NULL,NULL,0,201,1745753415,0),(233,'实验','admin',NULL,NULL,NULL,4,199,1745753473,0),(234,'国家教育质量检测英语模拟测试 2','admin',NULL,NULL,NULL,9,209,1746514369,0),(235,'国家教育质量检测艺术模拟测试 2','admin','',NULL,NULL,11,210,1746518544,0),(236,'国家教育质量检测语文模拟测试 2','admin',NULL,NULL,NULL,2598,211,1746633239,0),(237,'国家教育质量检测艺术模拟测试 2','admin','',NULL,NULL,0,210,1747035628,0),(238,'国家教育质量检测艺术模拟测试 2（贺兰县第一小学）','admin',NULL,NULL,NULL,2267,210,1747035651,0),(239,'国家教育质量检测艺术模拟测试 2（贺兰县第八小学）','admin',NULL,NULL,NULL,796,210,1747035672,0),(240,'国家教育质量检测英语模拟测试 2（贺兰县第一小学）','admin',NULL,NULL,NULL,2024,209,1747120617,0),(241,'国家教育质量检测英语模拟测试 2（贺兰县第八小学）','admin',NULL,NULL,NULL,798,209,1747120665,0);
