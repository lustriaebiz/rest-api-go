-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.4.11-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             10.3.0.5771
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- Dumping database structure for merchant_go
CREATE DATABASE IF NOT EXISTS `merchant_go` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `merchant_go`;

-- Dumping structure for table merchant_go.merchant
CREATE TABLE IF NOT EXISTS `merchant` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) DEFAULT NULL,
  `telp` varchar(20) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- Dumping data for table merchant_go.merchant: ~2 rows (approximately)
/*!40000 ALTER TABLE `merchant` DISABLE KEYS */;
INSERT INTO `merchant` (`id`, `name`, `telp`, `created_at`, `updated_at`) VALUES
	(1, 'Kranji', '081233343214', '2020-11-11 13:26:51', '2020-11-11 13:26:54'),
	(2, 'Sarjito', '02120000241', '2020-11-11 13:26:52', '2020-11-11 13:26:56'),
	(3, 'Rewash', '081222', '2020-11-11 13:26:53', '2020-11-11 13:26:57');
/*!40000 ALTER TABLE `merchant` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
