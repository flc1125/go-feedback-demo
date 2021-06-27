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

 Date: 27/06/2021 22:24:37
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
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_uid` (`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=24 DEFAULT CHARSET=utf8mb4 COMMENT='消息表';

-- ----------------------------
-- Records of message
-- ----------------------------
BEGIN;
INSERT INTO `message` VALUES (1, 1, '123123123', '2021-06-27 02:24:06');
INSERT INTO `message` VALUES (2, 1, 'asfasfasf ', '2021-06-27 02:24:11');
INSERT INTO `message` VALUES (3, 1, 'asdfasdfasf', '2021-06-27 02:24:12');
INSERT INTO `message` VALUES (4, 2, 'asdfasdf', '2021-06-27 02:24:13');
INSERT INTO `message` VALUES (5, 4, 'asdfasdf', '2021-06-27 02:24:19');
INSERT INTO `message` VALUES (6, 3, 'asdfasdf', '2021-06-27 02:24:15');
INSERT INTO `message` VALUES (7, 5, 'asdfasdf', '2021-06-27 02:24:23');
INSERT INTO `message` VALUES (8, 2, 'werqwer', '2021-06-25 02:24:28');
INSERT INTO `message` VALUES (9, 1, '按时发顺丰', '2021-06-27 02:24:25');
INSERT INTO `message` VALUES (10, 1, '大沙发斯蒂芬', '2021-04-27 02:24:27');
INSERT INTO `message` VALUES (11, 3, '啊啊阿斯蒂芬', '2021-06-27 02:24:29');
INSERT INTO `message` VALUES (12, 2, 'asdfasdf', '2021-06-27 02:24:32');
INSERT INTO `message` VALUES (13, 4, '阿斯顿撒', '2021-06-27 02:24:30');
INSERT INTO `message` VALUES (14, 2, '所有的业务逻辑实现均封装于service层中，不推荐实现于控制器api中。service层的包名只有一个，', '2021-06-27 05:43:47');
INSERT INTO `message` VALUES (15, 1, '上下文的变量必须在请求一开始便注入到请求流程中，以便于其他方法调用，因此我们使用中间件来实现。上下文的变量必须在请求一开始便注入到请求流程中，以便于其他方法调用，因此我们使用中间件来实现。', '2021-06-27 05:46:28');
INSERT INTO `message` VALUES (16, 1, '按服务而', '2021-06-27 08:32:49');
INSERT INTO `message` VALUES (17, 1, 'eqwerq ', '2021-06-27 16:40:12');
INSERT INTO `message` VALUES (18, 2, 'asdfasfasf', '2021-06-27 21:45:17');
INSERT INTO `message` VALUES (19, 4, 'aaaaadfasfdasfqwerqwr', '2021-06-27 21:47:28');
INSERT INTO `message` VALUES (20, 4, 'bbadfasfasf', '2021-06-27 21:47:34');
INSERT INTO `message` VALUES (21, 4, 'asdf', '2021-06-27 22:07:23');
INSERT INTO `message` VALUES (22, 10, 'asdfasf', '2021-06-27 22:14:50');
INSERT INTO `message` VALUES (23, 10, 'asfa', '2021-06-27 22:22:08');
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
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` VALUES (1, 'flc112', '$2a$10$el.0gbAN9S2aBBy2wy6Q9eGB/fBCRkRlgaEukjxNIKSBpeQuNMTXS', '2021-06-27 02:56:35', '2021-06-15 23:40:58', 1, 1, 0);
INSERT INTO `user` VALUES (2, 'asdf', '$2a$10$U4TJfB6vJHtUtZqsQNsvA.ggE9NWEdLH3JHD0Q36rtJr464uIoTmm', '2021-06-27 12:14:44', '2021-06-27 20:14:44', 0, 0, 0);
INSERT INTO `user` VALUES (3, '123123', '$2a$10$N50RQf.neDKiAJaB/Rh13.Ot1xA/MVkHkLHKBD2pbn6PZPyST9kTy', '2021-06-27 12:19:04', '2021-06-27 20:19:04', 0, 0, 0);
INSERT INTO `user` VALUES (4, 'aaaaa', '$2a$10$UG.i7FztETcL.azX8RUy4eN5QyE5H3BYPUGjlCZFSCyF9ug1hLtou', '2021-06-27 14:07:00', '2021-06-27 20:19:01', 1, 1, 0);
INSERT INTO `user` VALUES (5, 'flc1121111', '$2a$10$L5e4bz/xqEmBbsOIhs4lne/1sVge5e98e.wu6gMcDkbfZAHPQF4j6', '2021-06-27 12:15:43', '2021-06-27 20:15:43', 0, 0, 0);
INSERT INTO `user` VALUES (6, '12312322', '$2a$10$l6DtyCOfVdCjF2wgee6DPey05jXMjksca6/G91GkJ2pGNaIDHltEa', '2021-06-27 12:19:06', '2021-06-27 20:19:06', 0, 1, 1);
INSERT INTO `user` VALUES (7, 'zhangsan', '$2a$10$iG7/Iu9ubnoy8fahhVo6jusINWiOLl15JsLWqMmjxoJGDAaCJdrYK', '2021-06-26 22:53:31', '2021-06-26 22:53:31', 0, 1, 0);
INSERT INTO `user` VALUES (8, 'zhangsan11', '$2a$10$7hIsUAC11urxF4rP9z0RBeRDHmT7tPALULU5TrC3XEtmQC9kanob6', '2021-06-26 22:53:55', '2021-06-26 22:53:55', 0, 1, 0);
INSERT INTO `user` VALUES (9, 'tests123', '$2a$10$SX3lWq4AT0kz9/P5GUd2zOBCkVTHS3SM7l9xBwfanMn8qqApRDkny', '2021-06-26 23:17:39', '2021-06-26 23:17:39', 0, 1, 0);
INSERT INTO `user` VALUES (10, 'black', '$2a$10$jcweOEVEjuGUjjUfzlChGufGwwjHZca8rgqhU9p7xQuB823vwD0O.', '2021-06-27 14:22:12', '2021-06-27 22:22:12', 0, 1, 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
