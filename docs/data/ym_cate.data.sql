-- ----------------------------
-- Table structure for `ym_cate`
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
DROP TABLE IF EXISTS `ym_cate`;
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
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;


-- ----------------------------
-- Sample data for `ym_cate` (最早2条)
-- ----------------------------
INSERT INTO `ym_cate` VALUES (80,'国庆去哪里玩','admin','大家国庆节去过哪些地方',0,'',0.00,1709263732),(81,'你喜欢的景点','admin','列举几个可以多选哦',0,'列举几个可以多选哦',0.00,1709268271),(83,'多个复选题目','admin','',0,NULL,0.00,0),(84,'多个单选题目','admin','',0,'miaoshu',0.00,1586683233),(166,'银川市中小学规范办学行为专项督导检查学生问卷','admin',NULL,0,'请同学们根据问卷内容据实填写。',0.00,1680913597),(167,'银川市中小学规范办学行为专项督导检查学生问卷调查一下大家旅游的情况','admin',NULL,0,'请同学们根据问卷内容据实填写。',0.00,1709261650),(168,'动物知识测试','admin','你认识多少动物呢？',0,'',0.00,1709268003),(169,'植物知识测试 2','admin','你认识植物吗',0,'你认识植物吗',0.00,1709353004),(170,'一天上几节课','admin','上课情况',0,'',0.00,1709277531),(171,'一天上几节课，图','admin','上课情况',0,'上课情况',0.00,1709372688),(172,'图片题，认识植物','admin','',0,'',0.00,1709518023),(174,'中卫市初三音乐基础知识测试','中卫市教育局考试系统（演示版）','（测试版）',0,'（测试版）',0.00,1709556199),(176,'调查一下大家旅游的情况','admin',NULL,0,'',0.00,1709516644),(178,'调查一下大家旅游的情况11','admin','96',0,NULL,0.00,1709601731),(181,'一天上几节课','admin','1',0,'',0.00,1709602017),(182,'一天上几节课','admin','',0,'',0.00,1709602164),(183,'一天上几节课','admin','1',0,NULL,0.00,1743948928),(194,'国家教育质量检测数学模拟卷 1（中宁县第九小学教育集团）','admin','选择题部分',0,NULL,0.00,1716163699),(196,'国家教育质量检测心理健康模拟卷 1（中宁县第九小学教育集团）','admin','选择题部分',0,NULL,0.00,1716163714),(197,'国家教育质量检测心理健康模拟卷 2（中宁县第九小学教育集团）','admin','选择题部分',0,NULL,0.00,1716163736),(198,'国家教育质量检测数学模拟卷 （试验版）','admin','选择题部分',0,NULL,0.00,1744774911),(199,'实验','admin','',0,NULL,0.00,1743948679),(200,'调查一下大家旅游的情况','admin','',0,NULL,0.00,1743676923),(201,'试验版1','admin','',0,'',0.00,1743989395),(202,'试验版2','admin','',0,NULL,0.00,1744096005),(203,'图片测试','admin','1',0,'',0.00,1743993042),(204,'实验','admin','',0,'',0.00,1743993158),(206,'国家教育质量检测英语模拟测试 1','admin','客观题',0,NULL,0.00,1745376123),(207,'国家教育质量检测艺术模拟测试 1','admin','客观题（30道单选题+10道多选题）',0,NULL,0.00,1744727382),(208,'国家教育质量检测语文模拟测试 1','admin','客观题',0,NULL,0.00,1744707049),(209,'国家教育质量检测英语模拟测试 2（贺兰县第八小学）','admin','客观题',0,NULL,0.00,1747120630),(210,'国家教育质量检测艺术模拟测试 2（贺兰县第八小学）','admin','客观题',0,NULL,0.00,1747035664),(211,'国家教育质量检测语文模拟测试 2','admin','客观题',0,NULL,0.00,1746633234);
