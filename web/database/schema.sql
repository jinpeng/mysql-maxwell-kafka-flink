CREATE DATABASE  IF NOT EXISTS `contacts` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `contacts`;
CREATE USER 'contacts'@'localhost' IDENTIFIED BY 'password';
CREATE USER 'contacts'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON * . * TO 'contacts'@'%';
GRANT ALL PRIVILEGES ON * . * TO 'contacts'@'localhost';

CREATE USER 'maxwell'@'%' IDENTIFIED BY 'password-maxwell';
GRANT ALL ON maxwell.* TO 'maxwell'@'%';
GRANT SELECT, REPLICATION CLIENT, REPLICATION SLAVE ON *.* TO 'maxwell'@'%';
ALTER USER 'maxwell'@'%' IDENTIFIED WITH mysql_native_password BY 'password-maxwell';

--
-- Table structure for table `contact`
--
DROP TABLE IF EXISTS `contact`;
CREATE TABLE `contact` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `firstName` varchar(50) DEFAULT NULL,
  `lastName` varchar(50) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id_UNIQUE` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;