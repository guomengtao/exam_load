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
