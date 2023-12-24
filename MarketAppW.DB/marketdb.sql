-- Adminer 4.8.1 MySQL 8.0.35-0ubuntu0.22.04.1 dump

SET NAMES utf8;
SET time_zone = '+00:00';
SET foreign_key_checks = 0;
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `Products`;
CREATE TABLE `Products` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `category` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Products` (`id`, `name`, `category`) VALUES
(1,	'elma',	'meyve'),
(2,	'karpuz',	'sebze'),
(3,	'erik',	'meyve'),
(4,	'cengiz',	'insan'),
(5,	'Utku',	'insan'),
(6,	'muz',	'meyve');

DROP TABLE IF EXISTS `Sales`;
CREATE TABLE `Sales` (
  `id` int NOT NULL AUTO_INCREMENT,
  `product_id` int DEFAULT NULL,
  `sale_date` date DEFAULT NULL,
  `quantity` int DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `product_id` (`product_id`),
  CONSTRAINT `Sales_ibfk_1` FOREIGN KEY (`product_id`) REFERENCES `Products` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

INSERT INTO `Sales` (`id`, `product_id`, `sale_date`, `quantity`) VALUES
(1,	1,	'2023-12-19',	5),
(2,	1,	'2023-12-19',	35),
(3,	1,	'2023-12-19',	33),
(4,	1,	'2023-12-19',	155),
(5,	2,	'2023-12-19',	55),
(6,	3,	'2023-12-19',	35),
(7,	4,	'2023-12-20',	1),
(8,	5,	'2023-12-20',	1);

-- 2023-12-24 20:59:18
