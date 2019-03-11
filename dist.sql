/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : dist

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2019-03-05 10:10:26
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for passdist
-- ----------------------------
DROP TABLE IF EXISTS `passdist`;
CREATE TABLE `passdist` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `password` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of passdist
-- ----------------------------
INSERT INTO `passdist` VALUES ('1', '1');
INSERT INTO `passdist` VALUES ('2', '2');
INSERT INTO `passdist` VALUES ('3', '3');
INSERT INTO `passdist` VALUES ('4', '4');
INSERT INTO `passdist` VALUES ('5', '5');
INSERT INTO `passdist` VALUES ('6', '6');
INSERT INTO `passdist` VALUES ('7', '7');
INSERT INTO `passdist` VALUES ('8', '8');
INSERT INTO `passdist` VALUES ('9', '9');
INSERT INTO `passdist` VALUES ('10', '123456');
INSERT INTO `passdist` VALUES ('11', '1111');
INSERT INTO `passdist` VALUES ('12', '111111');
INSERT INTO `passdist` VALUES ('13', '123');

-- ----------------------------
-- Table structure for userdist
-- ----------------------------
DROP TABLE IF EXISTS `userdist`;
CREATE TABLE `userdist` (
  `id` int(10) unsigned zerofill NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=12 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of userdist
-- ----------------------------
INSERT INTO `userdist` VALUES ('0000000005', 'admin');
INSERT INTO `userdist` VALUES ('0000000006', 'root');
INSERT INTO `userdist` VALUES ('0000000007', '30wish');
INSERT INTO `userdist` VALUES ('0000000008', '123');
INSERT INTO `userdist` VALUES ('0000000009', '1');
INSERT INTO `userdist` VALUES ('0000000010', '111');
INSERT INTO `userdist` VALUES ('0000000011', 'test');
