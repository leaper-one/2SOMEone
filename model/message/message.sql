-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- 主机： mysql
-- 生成日期： 2022-07-26 02:44:42
-- 服务器版本： 5.7.38
-- PHP 版本： 8.0.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `2someone`
--

-- --------------------------------------------------------

--
-- 表的结构 `messages`
--

CREATE TABLE `messages` (
                            `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
                            `created_at` datetime(3) DEFAULT NULL,
                            `updated_at` datetime(3) DEFAULT NULL,
                            `deleted_at` datetime(3) DEFAULT NULL,
                            `phone` varchar(14) COLLATE utf8_bin DEFAULT NULL,
                            `type` tinyint(1) DEFAULT '0',
                            `content` varchar(512) COLLATE utf8_bin DEFAULT NULL,
                            `code` varchar(6) COLLATE utf8_bin DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            KEY `idx_messages_deleted_at` (`deleted_at`),
                            KEY `idx_messages_phone` (`phone`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

--
-- 转储表的索引
--

--
-- 表的索引 `messages`
--
ALTER TABLE `messages`
    ADD PRIMARY KEY (`id`),
    ADD KEY `idx_messages_deleted_at` (`deleted_at`),
    ADD KEY `idx_messages_phone` (`phone`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `messages`
--
ALTER TABLE `messages`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
