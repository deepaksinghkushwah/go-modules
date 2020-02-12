-- MySQL dump 10.16  Distrib 10.1.44-MariaDB, for debian-linux-gnu (x86_64)
--
-- Host: localhost    Database: shop
-- ------------------------------------------------------
-- Server version	10.1.44-MariaDB-0+deb9u1

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

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `email` varchar(100) DEFAULT NULL,
  `username` varchar(100) NOT NULL,
  `password` varchar(255) NOT NULL,
  `first_name` varchar(100) DEFAULT NULL,
  `last_name` varchar(100) DEFAULT NULL,
  `access_token` varchar(255) DEFAULT NULL,
  `access_token_expire_date` date DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_username_IDX` (`username`),
  KEY `user_id_IDX` (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'deepak@localhost.com','deepak','$2a$14$P5GGNo/KQwUl0dzrWI.fRePok747uX17YtHDfLo6m6lHVypkbX.Q6','deepak','kushwah','28a597cd-d557-44f7-8826-dbd1ac156e1c','2020-04-12'),(2,'test1@localhost.com','test1','$2a$14$8FB4QKVNn4LSioTOIMgFz.zuCnER2Sb8G906xEXWTNOJlWMN8S/z2','test','1','b05554e2-b3db-49d5-9e3b-4abf3034deea','2020-04-12'),(3,'test2@localhost.com','test2','$2a$14$veNXsGdgZYEdOXnpSeyPOuvYWDQqW1AiOW5uIlO7O.YHEA/Lx2x7K','test','2',NULL,NULL),(4,'test3@localhost.com','test3','$2a$14$rY.o0/2Uzog9YqqXb0H4s.cMTq/y5L1Vi109Iv/X9jXZg6lZbwpKe','Test','3','775efa2a-9f1d-4c65-9893-06cefd5092e5','2020-04-12'),(5,'test4@localhost.com','test4','$2a$14$ZLhfTVIxxN3bFgsQvloj2.1jIoKTb3bMrYEWQLzo6jqq32zesZ7cu','Test','4','213feb27-b03e-42b9-b70a-72ac0f0db6d2','2020-04-12');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping routines for database 'shop'
--
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2020-02-12 17:58:14
