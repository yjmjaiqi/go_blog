/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80011
 Source Host           : localhost:3306
 Source Schema         : register

 Target Server Type    : MySQL
 Target Server Version : 80011
 File Encoding         : 65001

 Date: 23/11/2022 11:58:59
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `blogname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `comment` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `blogid` int(11) NULL DEFAULT NULL,
  `userid` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `blogid`(`blogid`) USING BTREE,
  INDEX `userid`(`userid`) USING BTREE,
  CONSTRAINT `comment_ibfk_1` FOREIGN KEY (`blogid`) REFERENCES `article` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `comment_ibfk_2` FOREIGN KEY (`userid`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 23 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comment
-- ----------------------------
INSERT INTO `comment` VALUES (1, '张三', '隐形的翅膀', '张三:  童年爱听的歌，每一次...', 1, 1);
INSERT INTO `comment` VALUES (2, '张三', '人间四月天', '张三:  笋子哥哥好棒', 2, 1);
INSERT INTO `comment` VALUES (3, '张三', '隐形的翅膀', '张三:  我的博客记录点滴', 1, 1);
INSERT INTO `comment` VALUES (4, '张三', '人间四月天', '张三:  童年爱听的歌，每一次', 2, 1);
INSERT INTO `comment` VALUES (5, '张三', '人间四月天', '张三:  你好', 2, 1);
INSERT INTO `comment` VALUES (6, '张三', '隐形的翅膀', '张三:  你好', 1, 1);
INSERT INTO `comment` VALUES (7, '张三', '人间四月天', '张三:  听海哭的声音', 2, 1);
INSERT INTO `comment` VALUES (8, '张三', '人间四月天', '张三:  我的博客记录点滴', 2, 1);
INSERT INTO `comment` VALUES (9, '张三', '隐形的翅膀', '张三:  我的博客记录点滴', 1, 1);
INSERT INTO `comment` VALUES (10, '李四', '像鱼', '李四:  超爱听呢', 3, 2);
INSERT INTO `comment` VALUES (12, '张三', '隐形的翅膀', '张三:  小李子', 1, 1);
INSERT INTO `comment` VALUES (14, '张三', '隐形的翅膀', '张三:  怀念', 1, 1);
INSERT INTO `comment` VALUES (15, '张三', '像鱼', '张三:  恋爱的味道', 3, 1);
INSERT INTO `comment` VALUES (16, '李四', '人间四月天', '李四:  我的四月天', 2, 2);
INSERT INTO `comment` VALUES (17, '李四', '人间四月天', '李四:  我的四月天', 2, 2);
INSERT INTO `comment` VALUES (19, '张三', '人间四月天', '张三:  说得好', 7, 1);
INSERT INTO `comment` VALUES (20, '张三', '翅膀', '张三:  翅膀会飞吗？', 5, 1);
INSERT INTO `comment` VALUES (21, '张三', '隐身', '张三:  笋子哥哥好棒', 6, 1);
INSERT INTO `comment` VALUES (22, '先锋', '像鱼', '先锋:  jijio', 3, 4);

SET FOREIGN_KEY_CHECKS = 1;
