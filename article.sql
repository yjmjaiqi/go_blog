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

 Date: 23/11/2022 11:58:52
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for article
-- ----------------------------
DROP TABLE IF EXISTS `article`;
CREATE TABLE `article`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `blogname` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `content` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `comment` longtext CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `click` int(11) NULL DEFAULT 0,
  `blogtime` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `altertime` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '',
  `userid` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `userid`(`userid`) USING BTREE,
  CONSTRAINT `article_ibfk_1` FOREIGN KEY (`userid`) REFERENCES `user` (`id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of article
-- ----------------------------
INSERT INTO `article` VALUES (1, '张三', '隐形的翅膀', '作词 : 王雅君\r\n作曲 : 王雅君\r\n每一次都在徘徊孤单中坚强，每一次就算很受伤也不闪泪光，我知道我一直有双隐形的翅膀，带我飞飞过绝望，不去想他们拥有美丽的太阳，我看见每天的夕阳也会有变化，我知道我一直有双隐形的翅膀，带我飞给我希望，我终于看到所有梦想都开花，追逐的年轻 歌声多嘹亮，我终于翱翔 用心凝望不害怕，哪里会有风就飞多远吧，不去想他们拥有美丽的太阳，我看见每天的夕阳也会有变化，我知道我一直有双隐形的翅膀，带我飞 给我希望，我终于看到所有梦想都开花，追逐的年轻 歌声多嘹亮，我终于翱翔用心凝望不害怕，哪里会有风就飞多远吧，隐形的翅膀 让梦恒久比天长\r\n留一个愿望 让自己想象', '张三:  我的博客记录点滴;张三:  小李子;张三:  怀念;', 5, '2021-12-24 12:51:01', '', 1);
INSERT INTO `article` VALUES (2, '张三', '人间值得', '作词 : 林徽因\r\n作曲 : 马帅兵\r\n编曲: 马天瑞\r\n    你是人间值得,笑响点亮了四面风,你是早天里的云烟,黄昏吹着风的软,你是人间值得,细雨点洒在花丛间,你是一树树的花开,是燕子在梁间呢喃,那轻轻的,那娉婷的,鲜艳百花的冠冕,你是天真,你是庄严,你是夜夜的月圆。\r\n\r\n    你是人间值得\r\n细雨点洒在花丛间,你是一树树的花开,是燕子在梁间呢喃,那轻轻的,那娉婷的,鲜艳百花的冠冕,你是天真,你是庄严,你是夜夜的月圆,那轻轻的,那娉婷的,鲜艳百花的冠冕,你是心爱,你是温暖,你是夜夜的月圆,你是夜夜的月圆。', '张三:  笋子哥哥好棒张三:  童年爱听的歌，每一次;张三:  听海哭的声音张三:  我的博客记录点滴;李四:  我的四月天;', 16, '2021-12-24 12:57:34', '2021-12-24 20:42:04', 1);
INSERT INTO `article` VALUES (3, '李四', '像鱼', '作词 : 周有才\r\n作曲 : 周有才\r\n这是一首简单的歌，没有什么独特，试着代入我的心事，它那么幼稚，像个顽皮的孩子，多么可笑的心事，只剩我还在坚持，谁能看透我的眼睛，让我能够不再失明，我要记住你的样子，像鱼记住水的拥抱，像云在天空中停靠，夜晚的来到，也不会忘了阳光的温暖，我要忘了你的样子，像鱼忘了海的味道，放下所有梦和烦恼，却放不下回忆的乞讨，多么可笑的心事，只剩我还在坚持，谁能看透我的眼睛，让我能够不再失明，记住你的样子，像鱼记住水的拥抱，像云在天空中停靠，夜晚的来到，也不会忘了阳光的温暖，我要忘了你的样子，像鱼忘了海的味道，放下所有梦和烦恼，却放不下回忆的乞讨，只剩自己就好', '李四:  超爱听呢;张三:  恋爱的味道;先锋:  jijio;', 2, '2021-12-24 16:38:16', '', 2);
INSERT INTO `article` VALUES (5, '张三', '翅膀', '会飞', '张三:  翅膀会飞吗？;', 1, '2021-12-24 22:41:30', '', 1);
INSERT INTO `article` VALUES (6, '张三', '隐身', '你看不找', '张三:  笋子哥哥好棒;', 0, '2021-12-24 22:42:41', '', 1);
INSERT INTO `article` VALUES (7, '张三', '人间四月天', '在最好的年华，陪最美的你', '张三:  说得好;', 1, '2021-12-25 00:21:35', '', 1);
INSERT INTO `article` VALUES (8, '张三', '幸福来敲门', '你会开心吗', '', 2, '2021-12-25 00:22:08', '', 1);

SET FOREIGN_KEY_CHECKS = 1;
