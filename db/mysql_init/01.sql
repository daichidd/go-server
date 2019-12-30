CREATE DATABASE IF NOT EXISTS mytest;

-- add Grants to user `test`
GRANT SELECT,UPDATE,INSERT,DELETE ON `mytest`.* TO 'test'@'%';

USE mytest;

CREATE TABLE IF NOT EXISTS `todo` (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`user_id` int(10) unsigned NOT NULL,
	`title` varchar(50) NOT NULL,
	`description` varchar(255) DEFAULT NULL,
	`is_done` tinyint(3) unsigned NOT NULL DEFAULT '0',
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`id`)
) ENGINE=innodb DEFAULT CHARSET=utf8mb4;

CREATE TABLE IF NOT EXISTS `users` (
	`id` int(10) unsigned NOT NULL AUTO_INCREMENT,
	`name` varchar(50) NOT NULL,
	`created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY(`id`)
) ENGINE=innodb DEFAULT CHARSET=utf8mb4;
