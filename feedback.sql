/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : localhost:3306
 Source Schema         : feedback

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 27/06/2021 00:26:49
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for message
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `uid` int(10) NOT NULL COMMENT '用户ID',
  `content` varchar(200) NOT NULL DEFAULT '' COMMENT '留言内容',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='消息表';

-- ----------------------------
-- Records of message
-- ----------------------------
BEGIN;
INSERT INTO `message` VALUES (1, 0, '123123123', '2021-06-27 00:26:02');
INSERT INTO `message` VALUES (2, 0, 'asfasfasf ', '2021-06-27 00:26:15');
INSERT INTO `message` VALUES (3, 0, 'asdfasdfasf', '2021-06-27 00:26:22');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `username` varchar(40) NOT NULL COMMENT '用户名',
  `password` varchar(100) NOT NULL COMMENT '密码',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT ' 注册时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `is_admin` tinyint(2) DEFAULT '0' COMMENT '是否为管理员',
  `status` tinyint(2) DEFAULT '1' COMMENT '是否可用',
  `black` tinyint(2) DEFAULT '0' COMMENT '是否为黑名单',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'flc112', '$2a$10$el.0gbAN9S2aBBy2wy6Q9eGB/fBCRkRlgaEukjxNIKSBpeQuNMTXS', '2021-06-15 23:40:58', '2021-06-15 23:40:58', 0, 1, 0);
INSERT INTO `user` VALUES (2, 'asdf', '$2a$10$U4TJfB6vJHtUtZqsQNsvA.ggE9NWEdLH3JHD0Q36rtJr464uIoTmm', '2021-06-26 21:56:37', '2021-06-26 21:56:37', 0, 1, 0);
INSERT INTO `user` VALUES (3, '123123', '$2a$10$N50RQf.neDKiAJaB/Rh13.Ot1xA/MVkHkLHKBD2pbn6PZPyST9kTy', '2021-06-26 22:14:03', '2021-06-26 22:14:03', 0, 1, 0);
INSERT INTO `user` VALUES (4, 'aaaaa', '$2a$10$UG.i7FztETcL.azX8RUy4eN5QyE5H3BYPUGjlCZFSCyF9ug1hLtou', '2021-06-26 22:14:45', '2021-06-26 22:14:45', 0, 1, 0);
INSERT INTO `user` VALUES (5, 'flc1121111', '$2a$10$L5e4bz/xqEmBbsOIhs4lne/1sVge5e98e.wu6gMcDkbfZAHPQF4j6', '2021-06-26 22:52:03', '2021-06-26 22:52:03', 0, 1, 0);
INSERT INTO `user` VALUES (6, '12312322', '$2a$10$l6DtyCOfVdCjF2wgee6DPey05jXMjksca6/G91GkJ2pGNaIDHltEa', '2021-06-26 22:52:17', '2021-06-26 22:52:17', 0, 1, 0);
INSERT INTO `user` VALUES (7, 'zhangsan', '$2a$10$iG7/Iu9ubnoy8fahhVo6jusINWiOLl15JsLWqMmjxoJGDAaCJdrYK', '2021-06-26 22:53:31', '2021-06-26 22:53:31', 0, 1, 0);
INSERT INTO `user` VALUES (8, 'zhangsan11', '$2a$10$7hIsUAC11urxF4rP9z0RBeRDHmT7tPALULU5TrC3XEtmQC9kanob6', '2021-06-26 22:53:55', '2021-06-26 22:53:55', 0, 1, 0);
INSERT INTO `user` VALUES (9, 'tests123', '$2a$10$SX3lWq4AT0kz9/P5GUd2zOBCkVTHS3SM7l9xBwfanMn8qqApRDkny', '2021-06-26 23:17:39', '2021-06-26 23:17:39', 0, 1, 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
