# ************************************************************
# Sequel Ace SQL dump
# 版本号： 20064
#
# https://sequel-ace.com/
# https://github.com/Sequel-Ace/Sequel-Ace
#
# 主机: localhost (MySQL 8.0.39)
# 数据库: heritage
# 生成时间: 2024-09-19 12:31:07 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
SET NAMES utf8mb4;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE='NO_AUTO_VALUE_ON_ZERO', SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# 转储表 account
# ------------------------------------------------------------

DROP TABLE IF EXISTS `account`;

CREATE TABLE `account` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '区块链地址',
  `password` varchar(255) DEFAULT NULL COMMENT '密码',
  `city` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '城市',
  `level` int DEFAULT NULL COMMENT '级别（1省级机构，2市级单位）',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



# 转储表 heritage_inheritor
# ------------------------------------------------------------

DROP TABLE IF EXISTS `heritage_inheritor`;

CREATE TABLE `heritage_inheritor` (
  `number` varchar(255) NOT NULL,
  `Inheritor` varchar(255) DEFAULT NULL,
  `InheritorImg` varchar(255) DEFAULT NULL,
  `project` varchar(255) DEFAULT NULL,
  `rank` varchar(255) DEFAULT NULL,
  `category` varchar(255) DEFAULT NULL,
  `details` text,
  `locate` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`number`),
  UNIQUE KEY `number` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



# 转储表 heritage_project
# ------------------------------------------------------------

DROP TABLE IF EXISTS `heritage_project`;

CREATE TABLE `heritage_project` (
  `number` varchar(255) NOT NULL,
  `category` varchar(255) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `rank` varchar(255) DEFAULT NULL,
  `locate` varchar(255) DEFAULT NULL,
  `details` text,
  PRIMARY KEY (`number`),
  UNIQUE KEY `number` (`number`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;



# 转储表 heritage_task
# ------------------------------------------------------------

DROP TABLE IF EXISTS `heritage_task`;

CREATE TABLE `heritage_task` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'uuid',
  `field` text COMMENT '字段',
  `type` int DEFAULT NULL COMMENT '类型',
  `create_time` varchar(255) DEFAULT NULL COMMENT '时间',
  `locate` varchar(255) DEFAULT NULL COMMENT '地区',
  `apply_level` int DEFAULT NULL COMMENT '申请级别',
  `pass_level` int DEFAULT NULL COMMENT '审核通过级别',
  `number` varchar(255) DEFAULT NULL COMMENT '审核通过后的number',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
