-- MySQL dump 10.13  Distrib 9.2.0, for macos15 (arm64)
--
-- Host: 127.0.0.1    Database: mvc_foodopia_v2
-- ------------------------------------------------------
-- Server version	9.4.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `category`
--

DROP TABLE IF EXISTS `category`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `cat_id` bigint NOT NULL AUTO_INCREMENT,
  `cat_name` varchar(255) NOT NULL,
  `cat_description` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`cat_id`)
) ENGINE=InnoDB AUTO_INCREMENT=25 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

LOCK TABLES `category` WRITE;
/*!40000 ALTER TABLE `category` DISABLE KEYS */;
INSERT INTO `category` VALUES (1,'South Indian','Traditional dishes from South India including dosas, idlis, sambars and coconut-based curries'),(2,'North Indian','Rich and flavorful dishes from North India with creamy gravies, rotis and tandoor items'),(3,'Gujarati','Sweet and savory vegetarian dishes from Gujarat with unique flavor combinations'),(4,'Punjabi','Hearty and robust dishes from Punjab with rich gravies and bread varieties'),(5,'Bengali','Traditional Bengali vegetarian dishes with subtle flavors and fish-free preparations'),(6,'Rajasthani','Royal cuisine from Rajasthan with rich gravies, dal preparations and traditional breads'),(7,'Spicy','Hot and fiery dishes with intense heat levels and bold flavors'),(8,'Sweet','Dishes with natural or added sweetness, including desserts and sweet preparations'),(9,'Tangy','Dishes with acidic, sour or citrusy flavors that stimulate the palate'),(10,'Mild','Gentle flavors suitable for all age groups with minimal spice levels'),(11,'Light & Healthy','Nutritious, low-calorie dishes perfect for health-conscious diners'),(12,'Comfort Food','Hearty, satisfying dishes that provide emotional and physical comfort'),(13,'Street Food','Popular roadside snacks and quick bites with authentic Indian flavors'),(14,'Desserts','Traditional Indian sweets and desserts to end meals on a sweet note'),(15,'Breakfast','Morning meals and dishes traditionally consumed to start the day'),(16,'Main Course','Substantial dishes that form the centerpiece of a complete meal'),(17,'Fried','Deep-fried or pan-fried dishes with crispy textures and rich flavors'),(18,'Steamed','Healthy steamed preparations that retain natural flavors and nutrients'),(19,'Grilled','Tandoor or grill-cooked items with smoky flavors and charred textures'),(20,'Curry','Gravy-based dishes with complex spice blends and rich sauces'),(21,'Vegan','Plant-based dishes without any dairy or animal products'),(22,'Gluten-Free','Dishes made without wheat, suitable for gluten-sensitive individuals'),(23,'Jain','Vegetarian dishes prepared without onion, garlic, and root vegetables');
/*!40000 ALTER TABLE `category` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `item`
--

DROP TABLE IF EXISTS `item`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `item` (
  `item_id` bigint NOT NULL AUTO_INCREMENT,
  `item_name` varchar(255) NOT NULL,
  `cook_time_min` bigint DEFAULT NULL,
  `price` bigint NOT NULL,
  `display_pic` varchar(500) DEFAULT NULL,
  `cat_id` bigint NOT NULL,
  `subcat_id` bigint DEFAULT NULL,
  `is_available` tinyint(1) DEFAULT '1',
  PRIMARY KEY (`item_id`),
  KEY `cat_id` (`cat_id`),
  KEY `subcat_id` (`subcat_id`),
  CONSTRAINT `item_ibfk_1` FOREIGN KEY (`cat_id`) REFERENCES `category` (`cat_id`),
  CONSTRAINT `item_ibfk_2` FOREIGN KEY (`subcat_id`) REFERENCES `category` (`cat_id`)
) ENGINE=InnoDB AUTO_INCREMENT=82 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `item`
--

LOCK TABLES `item` WRITE;
/*!40000 ALTER TABLE `item` DISABLE KEYS */;
INSERT INTO `item` VALUES (1,'Masala Dosa',25,120,NULL,1,15,1),(2,'Idli Sambar',20,80,NULL,1,18,1),(3,'Rava Upma',15,70,NULL,1,11,1),(4,'Medu Vada',20,90,NULL,1,17,1),(5,'Coconut Rice',30,110,NULL,1,16,1),(6,'Filter Coffee',5,40,NULL,1,10,1),(7,'Rasam',25,60,NULL,1,9,1),(8,'Pongal',30,100,NULL,1,12,1),(9,'Paneer Butter Masala',35,180,NULL,2,20,1),(10,'Chole Bhature',30,150,NULL,2,13,1),(11,'Rajma Rice',40,140,NULL,2,12,1),(12,'Aloo Paratha',25,100,NULL,2,15,1),(13,'Dal Makhani',45,160,NULL,2,20,1),(14,'Butter Naan',15,50,NULL,2,16,1),(15,'Tandoori Roti',12,35,NULL,2,19,1),(16,'Dhokla',20,80,NULL,3,18,1),(17,'Khandvi',25,90,NULL,3,11,1),(18,'Gujarati Thali',15,220,NULL,3,16,1),(19,'Fafda Jalebi',20,110,NULL,3,8,1),(20,'Handvo',35,120,NULL,3,17,1),(21,'Undhiyu',50,180,NULL,3,7,1),(22,'Sarson Ka Saag',45,160,NULL,4,20,1),(23,'Makki Ki Roti',20,60,NULL,4,16,1),(24,'Punjabi Kadhi',35,140,NULL,4,9,1),(25,'Stuffed Kulcha',25,120,NULL,4,17,1),(26,'Amritsari Kulcha',30,130,NULL,4,13,1),(27,'Aloo Posto',30,120,NULL,5,10,1),(28,'Bengali Dal',25,100,NULL,5,20,1),(29,'Cholar Dal',35,130,NULL,5,8,1),(30,'Aloo Bhaja',15,70,NULL,5,17,1),(31,'Dal Baati Churma',60,200,NULL,6,12,1),(32,'Gatte Ki Sabji',40,150,NULL,6,7,1),(33,'Bajre Ki Roti',20,45,NULL,6,16,1),(34,'Ker Sangri',35,140,NULL,6,9,1),(35,'Mirchi Bada',20,85,NULL,7,17,1),(36,'Spicy Paneer Tikka',25,180,NULL,7,19,1),(37,'Achari Aloo',30,110,NULL,7,20,1),(38,'Gulab Jamun',15,80,NULL,8,14,1),(39,'Rasgulla',10,70,NULL,8,14,1),(40,'Kheer',30,90,NULL,8,14,1),(41,'Jalebi',15,60,NULL,8,17,1),(42,'Pani Puri',10,50,NULL,9,13,1),(43,'Bhel Puri',8,60,NULL,9,13,1),(44,'Aam Panna',5,45,NULL,9,11,1),(45,'Plain Rice',20,60,NULL,10,16,1),(46,'Khichdi',25,80,NULL,10,12,1),(47,'Buttermilk',3,30,NULL,10,11,1),(48,'Sprouts Salad',10,70,NULL,11,21,1),(49,'Vegetable Soup',15,65,NULL,11,18,1),(50,'Fruit Chaat',8,80,NULL,11,9,1),(51,'Green Salad',5,50,NULL,11,21,1),(52,'Poha',15,60,NULL,12,15,1),(53,'Maggi Noodles',8,55,NULL,12,7,1),(54,'Aloo Parantha',25,95,NULL,12,17,1),(55,'Vada Pav',12,40,NULL,13,17,1),(56,'Pav Bhaji',20,90,NULL,13,7,1),(57,'Sev Puri',8,55,NULL,13,9,0),(58,'Dahi Puri',10,65,NULL,13,8,1),(59,'Kulfi',5,60,NULL,14,8,1),(60,'Ras Malai',12,90,NULL,14,8,1),(61,'Gajar Halwa',35,100,NULL,14,8,1),(62,'Aloo Poha',15,70,NULL,15,10,1),(63,'Upma',12,60,NULL,15,1,1),(64,'Paratha with Curd',20,85,NULL,15,10,1),(65,'Samosa',18,45,NULL,17,13,1),(66,'Pakora',15,55,NULL,17,7,1),(67,'Aloo Tikki',20,70,NULL,17,13,1),(68,'Vegetable Momos',20,120,NULL,18,11,1),(69,'Kothimbir Vadi',25,80,NULL,18,11,1),(70,'Paneer Tikka',25,160,NULL,19,10,1),(71,'Grilled Vegetables',20,140,NULL,19,11,1),(72,'Mixed Vegetable Curry',35,130,NULL,20,16,1),(73,'Palak Paneer',30,150,NULL,20,11,1),(74,'Baingan Bharta',40,120,NULL,21,20,1),(75,'Aloo Jeera',20,90,NULL,21,10,1),(76,'Chana Masala',35,110,NULL,21,7,1),(77,'Jain Pav Bhaji',25,95,NULL,23,13,1),(78,'Jain Pizza',20,180,NULL,23,16,1),(79,'Jain Fried Rice',25,130,NULL,23,17,1);
/*!40000 ALTER TABLE `item` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2025-08-17  5:54:03
