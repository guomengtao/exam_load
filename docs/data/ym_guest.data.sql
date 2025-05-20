-- ----------------------------
-- Table structure for `ym_guest`
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
DROP TABLE IF EXISTS `ym_guest`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `ym_guest` (
  `id` mediumint(9) NOT NULL AUTO_INCREMENT,
  `nickname` varchar(30) NOT NULL,
  `email` varchar(30) NOT NULL,
  `homepage` varchar(150) NOT NULL,
  `content` text NOT NULL,
  `face` varchar(150) NOT NULL,
  `level` tinyint(1) NOT NULL DEFAULT '1',
  `artid` mediumint(9) NOT NULL DEFAULT '0',
  `pid` mediumint(9) NOT NULL DEFAULT '0',
  `time` int(10) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
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
-- Sample data for `ym_guest` (最早2条)
-- ----------------------------
INSERT INTO `ym_guest` VALUES (1,'fsadf','dfasdf@qq.com','www.122.com','sdfsdf','',1,0,0,1469539481),(2,'fsadf2','dfasd2f@qq.com','www.12222.com','sdfsdf2','',1,0,0,1469539687),(3,'童攀','tongpan@qq.com','www.12222.com','童攀童攀童攀童攀童攀童攀童攀童攀','',1,0,0,1469539721),(4,'童年','tongnian@qq.com','www.12222.com','童年童年童年童年童年童年童年童年童年童年童年童年童年','',1,0,0,1469539739),(5,'天天','tongnian@qq.com','www.12222.com','天天天天天天天天天天天天天天天天天天天天','',1,0,0,1469539760),(6,'dfsd','dfasdf@qq.com','www.12222.com','fasdfsdfsdf','',1,0,0,1469542092),(7,'dfsd','dfasdf@qq.com','www.12222.com','sdfasdfsdf','',1,0,0,1469542468);
