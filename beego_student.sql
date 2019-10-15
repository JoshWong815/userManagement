/*
 Navicat Premium Data Transfer

 Source Server         : localhost_3306
 Source Server Type    : MySQL
 Source Server Version : 80012
 Source Host           : localhost:3306
 Source Schema         : beego_student

 Target Server Type    : MySQL
 Target Server Version : 80012
 File Encoding         : 65001

 Date: 04/10/2019 13:04:11
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(50) NOT NULL,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `password` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL,
  `sex` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `age` int(11) NOT NULL,
  `usertype` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '馄钝教授', 'stan', '男', 21, '1');
INSERT INTO `user` VALUES (2, 'kenny', 'kyle', '男', 23, '1');
INSERT INTO `user` VALUES (3, '啊啊啊等等', 'xxx', 'a', 12, '1');
INSERT INTO `user` VALUES (4, 'cartman', '123456', '男', 33, '1');
INSERT INTO `user` VALUES (5, '测试测试', 'asd', '男', 10, '1');
INSERT INTO `user` VALUES (6, 'admin', '123456', '男', 15, '1');
INSERT INTO `user` VALUES (7, 'professorchaos', '123456', '男', 17, '1');
INSERT INTO `user` VALUES (8, 'kyle', '123456', '男', 33, '1');
INSERT INTO `user` VALUES (9, 'clyde', '123456', '男', 16, '3');
INSERT INTO `user` VALUES (10, '小黄油', '123456222', '男', 11, '1');

SET FOREIGN_KEY_CHECKS = 1;
