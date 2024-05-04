/*
 Navicat MySQL Data Transfer

 Source Server         : looklook
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : 127.0.0.1:33069
 Source Schema         : looklook_order

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 10/03/2022 17:15:38
*/
CREATE DATABASE looklook_tiktok1     DEFAULT CHARACTER SET utf8mb4     DEFAULT COLLATE utf8mb4_unicode_ci;
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;
ALTER TABLE douyin_live
ADD COLUMN open_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP;
-- ----------------------------
-- Table structure for homestay_order
-- ----------------------------
DROP TABLE IF EXISTS `douyin_live`;
CREATE TABLE `douyin_live` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `open_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `open_url` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '分享链接',
  `info` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '礼包信息',
  `name` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '主播',
  PRIMARY KEY (`id`)
  
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='抖音直播礼包';

