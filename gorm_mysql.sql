/*
 Navicat Premium Data Transfer

 Source Server         : localhost
 Source Server Type    : MySQL
 Source Server Version : 80100 (8.1.0)
 Source Host           : localhost:3308
 Source Schema         : gorm_mysql

 Target Server Type    : MySQL
 Target Server Version : 80100 (8.1.0)
 File Encoding         : 65001

 Date: 01/08/2023 10:48:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `username` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL,
  `id` int NULL DEFAULT NULL,
  `age` int NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('xiaozhang', 1, 11);
INSERT INTO `user` VALUES ('wangwu', 3, 21);
INSERT INTO `user` VALUES ('liuzi', 4, 50);
INSERT INTO `user` VALUES ('xiaoming', 6, 99);
INSERT INTO `user` VALUES ('xiaozhang', 1, 11);
INSERT INTO `user` VALUES ('xiaozhang', 1, 11);

SET FOREIGN_KEY_CHECKS = 1;
