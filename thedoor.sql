-- 创建用户的语句
-- CREATE USER 'thedoor'@'localhost' IDENTIFIED BY '5VB}4kcbMimyRK6^GY^J';
-- GRANT ALL PRIVILEGES ON `thedoor`.* TO 'thedoor'@'localhost' WITH GRANT OPTION;

-- 创建表的语句

-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        10.3.29-MariaDB-0ubuntu0.20.04.1 - Ubuntu 20.04
-- 服务器操作系统:                      debian-linux-gnu
-- HeidiSQL 版本:                  11.3.0.6337
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

-- 导出  表 thedoor.comments 结构
CREATE TABLE IF NOT EXISTS `comments` (
  `id` bigint(20) unsigned NOT NULL,
  `pid` bigint(20) unsigned NOT NULL DEFAULT 0,
  `fid` bigint(20) unsigned NOT NULL DEFAULT 0,
  `content` text NOT NULL DEFAULT '',
  `createtime` datetime NOT NULL DEFAULT current_timestamp(),
  `name` char(20) NOT NULL DEFAULT '',
  `email` char(80) NOT NULL DEFAULT '',
  `site` char(100) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 thedoor.posts 结构
CREATE TABLE IF NOT EXISTS `posts` (
  `id` bigint(20) unsigned NOT NULL,
  `title` char(100) NOT NULL DEFAULT '',
  `content` text NOT NULL DEFAULT '',
  `createtime` datetime NOT NULL DEFAULT current_timestamp(),
  `public` tinyint(1) NOT NULL DEFAULT 0,
  `top` bigint(20) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 thedoor.taglists 结构
CREATE TABLE IF NOT EXISTS `taglists` (
  `id` bigint(20) unsigned NOT NULL,
  `pid` bigint(20) unsigned NOT NULL DEFAULT 0,
  `tid` bigint(20) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE_TID_PID` (`tid`,`pid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

-- 导出  表 thedoor.tags 结构
CREATE TABLE IF NOT EXISTS `tags` (
  `id` bigint(20) unsigned NOT NULL,
  `name` char(20) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `UNIQUE NAME` (`name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 数据导出被取消选择。

/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
