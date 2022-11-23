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

 Date: 23/11/2022 11:59:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '张三', '123456');
INSERT INTO `user` VALUES (2, '李四', '123456');
INSERT INTO `user` VALUES (3, '王五', '123456');
INSERT INTO `user` VALUES (4, '先锋', '123456');
INSERT INTO `user` VALUES (5, 'xxzx', '4QrcOUm6Wau+VuBX8g+IPg==');
INSERT INTO `user` VALUES (7, '2334310165@qq.com', 'FKNCN0qLz34rsrwsdVJOEg==');
INSERT INTO `user` VALUES (8, '2334310165@qq.com', 'FKNCN0qLz34rsrwsdVJOEg==');
INSERT INTO `user` VALUES (9, '2334310165@qq.com', 'FKNCN0qLz34rsrwsdVJOEg==');
INSERT INTO `user` VALUES (10, '2334310165@qq.com', 'FKNCN0qLz34rsrwsdVJOEg==');

SET FOREIGN_KEY_CHECKS = 1;
