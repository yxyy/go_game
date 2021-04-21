/*
Navicat MySQL Data Transfer

Source Server         : 127.0.01
Source Server Version : 50553
Source Host           : localhost:3306
Source Database       : picture_crawler

Target Server Type    : MYSQL
Target Server Version : 50553
File Encoding         : 65001

Date: 2020-10-27 16:08:12
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for yxyy_category
-- ----------------------------
DROP TABLE IF EXISTS `yxyy_category`;
CREATE TABLE `yxyy_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(50) NOT NULL,
  `mark` varchar(20) NOT NULL,
  `pid` int(11) NOT NULL DEFAULT '0',
  `url` varchar(50) DEFAULT NULL,
  `description` varchar(50) DEFAULT NULL,
  `sort` int(4) DEFAULT NULL,
  `author` varchar(30) DEFAULT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `mark` (`url`) USING BTREE
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of yxyy_category
-- ----------------------------
INSERT INTO `yxyy_category` VALUES ('1', '彼岸图网', 'netbian', '0', 'netbian', '彼岸图网', '0', 'lhc', '1603419312', '1603436277');
INSERT INTO `yxyy_category` VALUES ('2', '4K风景', '4kfengjing', '1', 'http://pic.netbian.com/4kfengjing/', '4K风景图片', '0', 'lhc', '1603422485', '1603436338');
INSERT INTO `yxyy_category` VALUES ('3', '4K美女', '4kmeinv', '1', 'http://pic.netbian.com/4kmeinv/', '4K美女', '0', 'lhc', '1603422579', '1603436355');
INSERT INTO `yxyy_category` VALUES ('4', '4K游戏', '4kyouxi', '1', 'http://pic.netbian.com/4kyouxi/', '4K游戏', '0', 'lhc', '1603422606', '1603436370');
INSERT INTO `yxyy_category` VALUES ('5', '4K动漫', '4kdongman', '1', 'http://pic.netbian.com/4kdongman/', '4K动漫', '0', 'lhc', '1603422638', '1603436379');
INSERT INTO `yxyy_category` VALUES ('6', '4K影视', '4kyingshi', '1', 'http://pic.netbian.com/4kyingshi/', '4K影视', '0', 'lhc', '1603422832', '1603436389');
INSERT INTO `yxyy_category` VALUES ('7', '4K明星', '4kmingxing', '1', 'http://pic.netbian.com/4kmingxing/', '4K明星', '0', 'lhc', '1603422937', '1603436410');
INSERT INTO `yxyy_category` VALUES ('8', '4K动物', '4kdongwu', '1', 'http://pic.netbian.com/4kdongwu/', '4K动物', '0', 'lhc', '1603422976', '1603436421');
INSERT INTO `yxyy_category` VALUES ('9', '4K人物', '4krenwu', '1', 'http://pic.netbian.com/4krenwu/', '4K人物', '0', 'lhc', '1603423004', '1603436433');

-- ----------------------------
-- Table structure for yxyy_crawler_menu
-- ----------------------------
DROP TABLE IF EXISTS `yxyy_crawler_menu`;
CREATE TABLE `yxyy_crawler_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) NOT NULL,
  `mark` varchar(30) NOT NULL,
  `crawler_url` varchar(100) NOT NULL,
  `cron_url` varchar(100) NOT NULL,
  `status` int(4) DEFAULT NULL,
  `cron_time` varchar(20) NOT NULL,
  `cron_fomat_time` varchar(20) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) DEFAULT NULL,
  `last_cron` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of yxyy_crawler_menu
-- ----------------------------
INSERT INTO `yxyy_crawler_menu` VALUES ('1', '彼岸图网', 'netbian', 'http://pic.netbian.com/', 'netbian/cron', '1', '23', '', '1603438867', '1603442017', null);

-- ----------------------------
-- Table structure for yxyy_menu
-- ----------------------------
DROP TABLE IF EXISTS `yxyy_menu`;
CREATE TABLE `yxyy_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `mark` varchar(20) DEFAULT NULL,
  `type` varchar(16) DEFAULT NULL,
  `author` varchar(20) DEFAULT NULL,
  `url` varchar(50) NOT NULL,
  `parent` int(4) NOT NULL,
  `icon` varchar(20) DEFAULT NULL,
  `weight` int(11) NOT NULL,
  `status` tinyint(1) NOT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of yxyy_menu
-- ----------------------------
INSERT INTO `yxyy_menu` VALUES ('1', '系统设置', 'sys', '', '', '#', '0', 'fa-cogs', '100', '0', '1603180616', '1603180616');
INSERT INTO `yxyy_menu` VALUES ('2', '菜单管理', 'menu', '', '', 'menu/index', '1', '', '0', '0', '1603181566', '1603181566');
INSERT INTO `yxyy_menu` VALUES ('3', '图片管理', 'picture', '', '', '#', '0', 'fa-bar-chart-o', '10', '0', '1603181603', '1603181603');
INSERT INTO `yxyy_menu` VALUES ('4', '账号管理', 'account', '', '', '#', '0', 'fa-user', '80', '0', '1603246516', '1603246516');
INSERT INTO `yxyy_menu` VALUES ('5', '管理员', 'admin', '', '', 'user/index', '4', '', '0', '0', '1603246573', '1603246573');
INSERT INTO `yxyy_menu` VALUES ('6', '用户组', 'user_group', '', '', 'user_group/index', '4', '', '0', '0', '1603246617', '1603246617');
INSERT INTO `yxyy_menu` VALUES ('7', '概述', 'welcome', '', '', 'welcome', '0', 'fa-bar-chart-o', '0', '0', '1603339075', '1603339075');
INSERT INTO `yxyy_menu` VALUES ('8', '图片信息', 'img_info', '', '', 'picture/info', '7', '', '0', '0', '1603352174', '1603352174');
INSERT INTO `yxyy_menu` VALUES ('9', '爬虫管理', 'crawler', '', '', '#', '0', 'fa-gamepad', '70', '0', '1603353999', '1603353999');
INSERT INTO `yxyy_menu` VALUES ('10', '彼岸图网', 'netbian', '', '', 'netbian/index', '9', '', '0', '0', '1603354137', '1603354137');
INSERT INTO `yxyy_menu` VALUES ('11', '图片分类', 'category', '', '', 'category/index', '3', '', '0', '0', '1603355809', '1603355809');
INSERT INTO `yxyy_menu` VALUES ('12', '图片列表', 'gallery', '', 'lhc', 'gallery/index', '3', '', '0', '0', '1603700741', '1603700741');

-- ----------------------------
-- Table structure for yxyy_user
-- ----------------------------
DROP TABLE IF EXISTS `yxyy_user`;
CREATE TABLE `yxyy_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL,
  `nickname` varchar(50) NOT NULL,
  `password` varchar(100) NOT NULL,
  `status` int(4) NOT NULL,
  `salt` varchar(100) NOT NULL,
  `group_id` int(11) NOT NULL,
  `phone` int(11) DEFAULT NULL,
  `expiration` int(11) DEFAULT NULL,
  `login_num` int(11) DEFAULT NULL,
  `last_login_ip` varchar(20) DEFAULT NULL,
  `last_login_time` int(11) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of yxyy_user
-- ----------------------------
INSERT INTO `yxyy_user` VALUES ('1', 'lhc', 'lhc', '64cc2a635cc6a2bbae5cafab765f9106', '0', '4aeJ6SroXP', '1', '0', '1634774400', '3', '127.0.0.1', '1603351300', '1603251395', '1603351300');

-- ----------------------------
-- Table structure for yxyy_user_group
-- ----------------------------
DROP TABLE IF EXISTS `yxyy_user_group`;
CREATE TABLE `yxyy_user_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `mark` varchar(20) NOT NULL,
  `description` varchar(50) DEFAULT NULL,
  `create_time` int(11) NOT NULL,
  `update_time` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of yxyy_user_group
-- ----------------------------
INSERT INTO `yxyy_user_group` VALUES ('1', '超级管理员', 'root', '超级管理员', '1603166010', '1603166010');
INSERT INTO `yxyy_user_group` VALUES ('2', '一级管理员', 'first_level', '一级管理员', '1603248677', '1603248677');
