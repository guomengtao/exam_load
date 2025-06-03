-- ----------------------------
-- Table structure for `tm_exam_papers`
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
DROP TABLE IF EXISTS `tm_exam_papers`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tm_exam_papers` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `uuid` char(36) DEFAULT NULL COMMENT '试卷UUID唯一标识',
  `template_id` bigint(20) DEFAULT NULL COMMENT '来源模板ID',
  `title` varchar(255) NOT NULL COMMENT '试卷标题',
  `description` text COMMENT '试卷描述',
  `cover_image` varchar(255) DEFAULT NULL COMMENT '封面图片URL',
  `total_score` int(11) NOT NULL DEFAULT '0' COMMENT '试卷总分',
  `questions` json NOT NULL COMMENT '题目列表（JSON数组）',
  `view_count` int(11) DEFAULT '0' COMMENT '浏览量',
  `category_id` bigint(20) DEFAULT NULL COMMENT '分类ID',
  `publish_time` int(10) DEFAULT '0' COMMENT '最近一次发布时间',
  `status` tinyint(1) DEFAULT '1' COMMENT '状态：1=正常，0=关闭',
  `creator` varchar(150) DEFAULT NULL COMMENT '创建人用户名',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间（NULL=未删除）',
  `time_limit` int(11) DEFAULT '0' COMMENT '考试时间限制（分钟），0表示无限时',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COMMENT='试卷';
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
-- Sample data for `tm_exam_papers` (最早2条)
-- ----------------------------
INSERT INTO `tm_exam_papers` VALUES (1,NULL,NULL,'修改后的试卷标题','这是更新后的试卷描述','',120,'[{\"id\": 1, \"type\": \"single\", \"score\": 5, \"title\": \"1+1=？\", \"options\": [\"1\", \"2\", \"3\"], \"image_url\": null, \"correct_answer\": \"2\"}]',0,0,0,0,'','2025-04-28 03:09:13','2025-04-28 03:09:13',NULL,0),(2,NULL,1,'修改后的试卷标题','这是更新后的试卷描述','',120,'[{\"id\": 1, \"type\": \"single\", \"score\": 5, \"title\": \"1+1=？\", \"options\": [\"1\", \"2\", \"3\"], \"image_url\": null, \"correct_answer\": \"2\"}]',0,0,0,0,'','2025-04-28 03:36:41','2025-04-28 03:36:41',NULL,0),(3,NULL,1,'修改后的试卷标题','这是更新后的试卷描述','',120,'[{\"id\": 1, \"type\": \"single\", \"score\": 5, \"title\": \"1+1=？\", \"options\": [\"1\", \"2\", \"3\"], \"image_url\": null, \"correct_answer\": \"2\"}]',0,0,0,0,'','2025-04-28 04:11:40','2025-04-28 04:11:40',NULL,0),(4,'aec21b84-d9fc-4390-acb6-d36cb85677ab',1,'修改后的试卷标题','这是更新后的试卷描述','',120,'[{\"id\": 1, \"type\": \"single\", \"score\": 5, \"title\": \"1+1=？\", \"options\": [\"1\", \"2\", \"3\"], \"image_url\": null, \"correct_answer\": \"2\"}]',0,0,0,0,'','2025-04-28 04:13:14','2025-04-28 04:13:14',NULL,0),(5,'4b7f9269-95f1-4274-ac59-49d8f4f27654',8,'历史','','',52,'[{\"type\": \"single\", \"score\": 52, \"title\": \"\", \"options\": [\"\", \"\", \"\", \"\"], \"image_url\": \"\", \"correct_answer_bitmask\": 0}]',0,0,0,0,'','2025-05-07 23:59:14','2025-05-07 23:59:14',NULL,0),(6,'2228868b-2b5b-4893-be94-c05fab7716a8',8,'历史','','',52,'[{\"type\": \"single\", \"score\": 52, \"title\": \"\", \"options\": [\"\", \"\", \"\", \"\"], \"image_url\": \"\", \"correct_answer_bitmask\": 0}]',0,0,0,0,'','2025-05-08 00:22:35','2025-05-08 00:22:35',NULL,0),(7,'2494a417-7af2-4dda-b6ce-aff96c1f2839',7,'2025年数学竞赛','','',150,'[{\"id\": 1, \"type\": \"multi\", \"score\": 10, \"title\": \"那些是水果\", \"options\": [\"西瓜\", \"电脑\", \"老虎\", \"火龙果\"], \"image_url\": null, \"correct_answer_bitmask\": 11}, {\"id\": 1746600355571, \"type\": \"single\", \"score\": 122, \"title\": \"567-10=?\", \"options\": [\"557\", \"4\", \"5\", \"666\"], \"image_url\": \"\", \"correct_answer_bitmask\": 1}, {\"id\": 1746600957041, \"type\": \"judge\", \"score\": 52, \"title\": \"\", \"options\": [\"\", \"\"], \"image_url\": \"\", \"correct_answer_bitmask\": 0}]',0,0,0,0,'教师B','2025-05-08 00:22:47','2025-05-08 00:22:47',NULL,0),(8,'21abb9a4-7530-4003-8b1d-62c0ed4e4844',6,'2025年期中考试','','',100,'[{\"id\": 1, \"type\": \"multi\", \"score\": 10, \"title\": \"‘Hello’这个单词是什么意思？\", \"options\": [\"你好\", \"再见\", \"谢谢\", \"不客气\"], \"image_url\": null, \"correct_answer\": [\"B\"], \"correct_answer_bitmask\": 2}]',0,0,0,0,'教师A','2025-05-08 00:23:25','2025-05-08 00:23:25',NULL,0),(9,'1aa5d0d3-adc2-4f58-9005-1489bbfa46bf',6,'2025年期中考试','','',100,'[{\"id\": 1, \"type\": \"multi\", \"score\": 10, \"title\": \"‘Hello’这个单词是什么意思？\", \"options\": [\"你好\", \"再见\", \"谢谢\", \"不客气\"], \"image_url\": null, \"correct_answer\": [\"B\"], \"correct_answer_bitmask\": 2}]',0,0,0,0,'教师A','2025-05-08 00:35:42','2025-05-08 00:35:42',NULL,0),(10,'a75d9756-d057-49b4-b3bf-a7ede078c8ba',7,'2025年数学竞赛','','',150,'[{\"id\": 1, \"type\": \"multi\", \"score\": 10, \"title\": \"那些是水果\", \"options\": [\"西瓜\", \"电脑\", \"老虎\", \"火龙果\"], \"image_url\": null, \"correct_answer_bitmask\": 11}, {\"id\": 1746600355571, \"type\": \"single\", \"score\": 5, \"title\": \"567-10=?\", \"options\": [\"557\", \"4\", \"5\", \"666\"], \"image_url\": \"\", \"correct_answer_bitmask\": 1}, {\"id\": 1746600957041, \"type\": \"multi\", \"score\": 2, \"title\": \"42222*5>?\", \"options\": [\"4\", \"55\", \"23\", \"2\"], \"image_url\": \"\", \"correct_answer_bitmask\": 15}]',0,0,0,0,'教师B','2025-05-08 01:35:29','2025-05-08 01:35:29',NULL,0),(11,'ffd468aa-cc52-477a-83ff-1be3958773a8',10,'测试题目','','',50,'[{\"type\": \"single\", \"score\": 30, \"title\": \"单选题,12+13=?\", \"options\": [\"22\", \"25\", \"21\", \"-1\"], \"image_url\": \"\", \"correct_answer_bitmask\": 2}, {\"type\": \"multi\", \"score\": 20, \"title\": \"多选题,哪些是单词\", \"options\": [\"cat\", \"老虎\", \"西瓜\", \"fly\"], \"image_url\": \"\", \"correct_answer_bitmask\": 9}]',0,0,0,0,'','2025-05-11 12:25:59','2025-05-11 12:25:59',NULL,0),(12,'234aa735-188f-4d57-a350-2df243780235',11,'带id的题目','','',5,'[{\"id\": 1746969239826, \"type\": \"multi\", \"score\": 5, \"title\": \"12+12大于几\", \"options\": [\"24\", \"22\", \"3\", \"33\"], \"image_url\": \"\", \"correct_answer_bitmask\": 6}]',0,0,0,0,'','2025-05-11 13:15:03','2025-05-11 13:15:03',NULL,0),(13,'73278397-19a4-4135-a08d-f14342de0552',11,'带id的题目','','',5,'[{\"id\": 1746969239826, \"type\": \"multi\", \"score\": 5, \"title\": \"12+12大于几\", \"options\": [\"24\", \"22\", \"3\", \"33\"], \"image_url\": \"\", \"correct_answer_bitmask\": 6}, {\"id\": 1747002892004, \"type\": \"single\", \"score\": 7, \"title\": \"单选题,3+5=?\", \"options\": [\"3\", \"2\", \"8\", \"4\"], \"image_url\": \"\", \"correct_answer_bitmask\": 4}]',0,0,0,0,'','2025-05-11 22:35:44','2025-05-11 22:35:44',NULL,0),(14,'0b1071fc-7fab-41f7-b61d-042a9f7f3731',11,'带id的题目','','',5,'[{\"id\": 1746969239826, \"type\": \"multi\", \"score\": 5, \"title\": \"12+12大于几\", \"options\": [\"24\", \"22\", \"3\", \"33\"], \"image_url\": \"\", \"correct_answer_bitmask\": 6}, {\"id\": 1747002892004, \"type\": \"single\", \"score\": 7, \"title\": \"单选题,3+5=?\", \"options\": [\"3\", \"2\", \"8\", \"4\"], \"image_url\": \"\", \"correct_answer_bitmask\": 4}]',0,0,0,0,'','2025-05-12 00:22:53','2025-05-12 00:22:53',NULL,0),(15,'40a89290-0ff9-4c01-a803-57155a24985c',11,'带id的题目','','',5,'[{\"id\": 1746969239826, \"type\": \"multi\", \"score\": 5, \"title\": \"12+12大于几\", \"options\": [\"24\", \"22\", \"3\", \"33\"], \"image_url\": \"\", \"correct_answer_bitmask\": 6}, {\"id\": 1747002892004, \"type\": \"single\", \"score\": 7, \"title\": \"单选题,3+5=?\", \"options\": [\"3\", \"2\", \"8\", \"4\"], \"image_url\": \"\", \"correct_answer_bitmask\": 4}]',0,0,0,0,'','2025-05-12 02:02:40','2025-05-12 02:02:40',NULL,0),(16,'3269763b-cdde-4eca-8a5f-a19a81ebf52c',12,'测试总分','','',12,'[{\"id\": 1747019451218, \"type\": \"single\", \"score\": 6, \"title\": \"12+13=?\", \"options\": [\"2\", \"25\", \"22\", \"33\"], \"image_url\": \"\", \"correct_answer_bitmask\": 2}, {\"id\": 1747019480623, \"type\": \"multi\", \"score\": 6, \"title\": \"大于20的数\", \"options\": [\"2\", \"22\", \"3\", \"56\"], \"image_url\": \"\", \"correct_answer_bitmask\": 10}]',0,0,0,0,'','2025-05-12 03:12:19','2025-05-12 03:12:19',NULL,0);
