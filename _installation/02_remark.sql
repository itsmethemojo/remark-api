DROP TABLE IF EXISTS `bookmark`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bookmark` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `url` varchar(400) CHARACTER SET latin1 NOT NULL,
  `domain` varchar(30) CHARACTER SET latin1 NOT NULL,
  `extension` varchar(10) CHARACTER SET latin1 NOT NULL,
  `title` varchar(200) CHARACTER SET latin1 NOT NULL,
  `customtitle` varchar(200) CHARACTER SET latin1 NOT NULL,
  `bookmarkcount` int(11) NOT NULL,
  `clickcount` int(11) NOT NULL,
  `tagcount` int(11) NOT NULL,
  `mediatype` int(11) NOT NULL,
  `created` bigint(20) NOT NULL,
  `updated` bigint(20) NOT NULL,
  `clicked` bigint(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `bookmarktime`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `bookmarktime` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `bookmark_id` int(11) NOT NULL,
  `created` bigint(20) NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `bookmark_id` (`bookmark_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `clicktime`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `clicktime` (
  `id` bigint(20) NOT NULL DEFAULT '0',
  `user_id` int(11) NOT NULL,
  `bookmark_id` int(11) NOT NULL,
  `created` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

DROP TABLE IF EXISTS `tagtime`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tagtime` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `bookmark_id` int(11) NOT NULL,
  `name` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;