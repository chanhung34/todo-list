/*
 Navicat Premium Data Transfer

 Source Server         : local
 Source Server Type    : MySQL
 Source Server Version : 80016
 Source Host           : localhost:3306
 Source Schema         : db

 Target Server Type    : MySQL
 Target Server Version : 80016
 File Encoding         : 65001

 Date: 23/05/2019 22:32:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for todo_items
-- ----------------------------
DROP TABLE IF EXISTS `todo_items`;
CREATE TABLE `todo_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `status` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `update_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of todo_items
-- ----------------------------
BEGIN;
INSERT INTO `todo_items` VALUES (1, 1, 'tokugawa', 'todo', '2019-05-23 14:03:39', '2019-05-23 15:12:39');
INSERT INTO `todo_items` VALUES (2, 1, 'takeda', 'done', '2019-05-23 14:03:43', '2019-05-23 15:13:35');
INSERT INTO `todo_items` VALUES (4, 1, 'ahihihihih', 'todo', '2019-05-23 14:03:50', '2019-05-23 14:03:50');
INSERT INTO `todo_items` VALUES (5, 1, 'yameteeeee', 'todo', '2019-05-23 14:03:56', '2019-05-23 14:03:56');
INSERT INTO `todo_items` VALUES (6, 1, 'yatao te', 'todo', '2019-05-23 14:04:03', '2019-05-23 14:04:03');
INSERT INTO `todo_items` VALUES (7, 1, 'hihsadasi', 'todo', '2019-05-23 14:30:47', '2019-05-23 14:30:47');
INSERT INTO `todo_items` VALUES (8, 1, 'test', 'todo', '2019-05-23 14:30:55', '2019-05-23 14:30:55');
INSERT INTO `todo_items` VALUES (9, 1, 'waiting', 'todo', '2019-05-23 14:31:02', '2019-05-23 14:31:02');
INSERT INTO `todo_items` VALUES (10, 1, 'master golang', 'todo', '2019-05-23 14:31:16', '2019-05-23 14:31:16');
COMMIT;

-- ----------------------------
-- Table structure for user_accounts
-- ----------------------------
DROP TABLE IF EXISTS `user_accounts`;
CREATE TABLE `user_accounts` (
  `id` int(50) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of user_accounts
-- ----------------------------
BEGIN;
INSERT INTO `user_accounts` VALUES (1, 'hungvtc', '123abc', '2019-05-22 01:17:43', '2019-05-22 01:17:43');
INSERT INTO `user_accounts` VALUES (4, 'hungvtc1', '123abctx context.Context, ctx context.Context, c', '2019-05-22 05:20:05', '2019-05-22 05:20:05');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
