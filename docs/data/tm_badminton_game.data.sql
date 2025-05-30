-- ----------------------------
-- Table structure for `tm_badminton_game`
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
DROP TABLE IF EXISTS `tm_badminton_game`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tm_badminton_game` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '记录ID',
  `player1` varchar(50) NOT NULL COMMENT '选手1姓名',
  `player2` varchar(50) NOT NULL COMMENT '选手2姓名',
  `score1` int(11) DEFAULT '0' COMMENT '选手1得分',
  `score2` int(11) DEFAULT '0' COMMENT '选手2得分',
  `location` varchar(100) DEFAULT NULL COMMENT '比赛地点',
  `match_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '比赛时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=79 DEFAULT CHARSET=utf8mb4 COMMENT='羽毛球比赛记录表';
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
-- Sample data for `tm_badminton_game` (最早2条)
-- ----------------------------
INSERT INTO `tm_badminton_game` VALUES (3,'Tom','Jack',21,15,'','0000-00-00 00:00:00'),(4,'Alice','Bob',44,4,'','0000-00-00 00:00:00'),(5,'Tom','Jack',55,5,'','0000-00-00 00:00:00'),(6,'Alice','Bob',18,21,'','0000-00-00 00:00:00'),(7,'Tom','Jack',21,15,'','0000-00-00 00:00:00'),(8,'Alice','Bob',18,21,'','0000-00-00 00:00:00'),(9,'Tom','Jack',21,15,'','0000-00-00 00:00:00'),(10,'Alice','Bob',18,21,'','0000-00-00 00:00:00'),(11,'Tom','Jack',21,15,'','0000-00-00 00:00:00'),(12,'Alice','Bob',18,21,'','0000-00-00 00:00:00'),(13,'Tom','Jack',21,15,'','0000-00-00 00:00:00'),(14,'Alice','Bob',18,21,'','0000-00-00 00:00:00'),(15,'Tom','Jack',21,15,'','0000-00-00 00:00:00'),(16,'Alice','Bob',18,21,'','0000-00-00 00:00:00'),(17,'Tom','Jack',21,15,'','0000-00-00 00:00:00'),(18,'Alice','Bob',18,21,'','0000-00-00 00:00:00'),(73,'Tom','Jack',21,15,NULL,NULL),(74,'Alice','Bob',18,21,NULL,NULL),(75,'Tom','Jack',21,15,NULL,NULL),(76,'Alice','Bob',18,21,NULL,NULL),(77,'Tom','Jack',21,15,NULL,NULL),(78,'Alice','Bob',18,21,NULL,NULL);
