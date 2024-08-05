/*
 Navicat Premium Data Transfer

 Source Server         : 192.168.1.20
 Source Server Type    : MySQL
 Source Server Version : 50728
 Source Host           : 192.168.1.20:3306
 Source Schema         : sinfotek

 Target Server Type    : MySQL
 Target Server Version : 50728
 File Encoding         : 65001

 Date: 02/06/2023 14:33:35
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for alarmsetting
-- ----------------------------
DROP TABLE IF EXISTS `alarmsetting`;
CREATE TABLE `alarmsetting`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `hostid` bigint(20) NULL DEFAULT NULL,
  `alarmtype` int(255) NULL DEFAULT NULL,
  `alarmstatus` int(255) NULL DEFAULT NULL,
  `alarmhostonwer` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `fk_hostid`(`hostid`) USING BTREE,
  CONSTRAINT `fk_hostid` FOREIGN KEY (`hostid`) REFERENCES `hostlist` (`hostid`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE = InnoDB AUTO_INCREMENT = 33 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of alarmsetting
-- ----------------------------
INSERT INTO `alarmsetting` VALUES (6, 10086, 4011, 1, 'op');

-- ----------------------------
-- Table structure for alarmstatistics
-- ----------------------------
DROP TABLE IF EXISTS `alarmstatistics`;
CREATE TABLE `alarmstatistics`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `alarmid` bigint(20) NULL DEFAULT NULL COMMENT '报警id',
  `hostid` bigint(20) NOT NULL COMMENT '主机id',
  `alarmstatus` int(10) NULL DEFAULT NULL COMMENT '报警状态',
  `alarmtype` int(20) NULL DEFAULT NULL COMMENT '报警类型',
  `alarminfo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '报警信息',
  `alarmnote` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '报警备注',
  `alarmstarttime` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '报警开始时间',
  `alarmstoptime` timestamp(0) NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '报警结束时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_id`(`id`) USING BTREE,
  UNIQUE INDEX `idx_alarmid`(`alarmid`) USING BTREE,
  INDEX `fx_alarmstatistics`(`hostid`) USING BTREE,
  CONSTRAINT `fx_alarmstatistics` FOREIGN KEY (`hostid`) REFERENCES `hostlist` (`hostid`) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for alarmtype
-- ----------------------------
DROP TABLE IF EXISTS `alarmtype`;
CREATE TABLE `alarmtype`  (
  `alarmtypeid` int(11) NOT NULL,
  `alarmtypename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `emergencylevel` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`alarmtypeid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of alarmtype
-- ----------------------------
INSERT INTO `alarmtype` VALUES (1000, '应用服务问题', '警告');
INSERT INTO `alarmtype` VALUES (1001, '系统问题', '一般');
INSERT INTO `alarmtype` VALUES (1004, '网络问题', '严重');
INSERT INTO `alarmtype` VALUES (1006, '硬件问题', '故障');

-- ----------------------------
-- Table structure for filedata
-- ----------------------------
DROP TABLE IF EXISTS `filedata`;
CREATE TABLE `filedata`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `fileid` bigint(20) NULL DEFAULT NULL,
  `filename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `fileoption` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `fileinfo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `optiontime` timestamp(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for filelog
-- ----------------------------
DROP TABLE IF EXISTS `filelog`;
CREATE TABLE `filelog`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `filename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `fileid` bigint(20) NULL DEFAULT NULL,
  `uploadtime` timestamp(0) NULL DEFAULT NULL,
  `filedir` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `filesize` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 28 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for hostdata
-- ----------------------------
DROP TABLE IF EXISTS `hostdata`;
CREATE TABLE `hostdata`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cpupart` int(10) NULL DEFAULT NULL,
  `rampart` int(10) NULL DEFAULT NULL,
  `insertdatatime` timestamp(0) NULL DEFAULT NULL,
  `hostid` bigint(20) NULL DEFAULT NULL,
  `uns` int(255) NULL DEFAULT NULL,
  `dns` int(255) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for hostlist
-- ----------------------------
DROP TABLE IF EXISTS `hostlist`;
CREATE TABLE `hostlist`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '主机序号',
  `hostid` bigint(20) NOT NULL,
  `hostname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主机名',
  `systemtype` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '系统类型',
  `hoststatus` int(10) NOT NULL DEFAULT 1 COMMENT '系统状态',
  `hostip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '主机IP',
  `hostlocation` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '主机位置',
  `hostowner` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '主机负责人',
  `hostaddtime` timestamp(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '主机添加时间',
  `hostnote` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '主机备注',
  `hostsysteminfo` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '主机信息',
  `hostuptime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '主机运行时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_hostip`(`hostip`) USING BTREE,
  UNIQUE INDEX `idx_hostname`(`hostname`) USING BTREE,
  INDEX `hostid`(`hostid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 64 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of hostlist
-- ----------------------------
INSERT INTO `hostlist` VALUES (1, 10086, 'sinfotek', 'windows', 1, '192.168.0.117', '北京', 'op1', '2023-05-05 14:19:21.000000', '测试12', '\r\n\r\nCaption           : Intel64 Family 6 Model 94 Stepping 3\r\nDeviceID          : CPU0\r\nManufacturer      : GenuineIntel\r\nMaxClockSpeed     : 3501\r\nName              : Intel(R) Core(TM) i5-6600K CPU @ 3.50GHz\r\nSocketDesignation : U3E1\r\n\r\n\r\n\r\n\r\n\r\nSMBIOSBIOSVersion : F2\r\nManufacturer      : American Megatrends Inc.\r\nName              : BIOS Date: 11/22/16 14:48:07 Ver: 05.0000C\r\nSerialNumber      : Default string\r\nVersion           : ALASKA - 1072009\r\n\r\n\r\n\r\n\r\n\r\n__GENUS              : 2\r\n__CLASS              : Win32_PhysicalMemory\r\n__SUPERCLASS         : CIM_PhysicalMemory\r\n__DYNASTY            : CIM_ManagedSystemElement\r\n__RELPATH            : Win32_PhysicalMemory.Tag=\"Physical Memory 1\"\r\n__PROPERTY_COUNT     : 36\r\n__DERIVATION         : {CIM_PhysicalMemory, CIM_Chip, CIM_PhysicalComponent, CIM_PhysicalElement...}\r\n__SERVER             : DESKTOP-48RD9C2\r\n__NAMESPACE          : root\\cimv2\r\n__PATH               : \\\\DESKTOP-48RD9C2\\root\\cimv2:Win32_PhysicalMemory.Tag=\"Physical Memory 1\"\r\nAttributes           : 1\r\nBankLabel            : BANK 1\r\nCapacity             : 8589934592\r\nCaption              : 物理内存\r\nConfiguredClockSpeed : 2133\r\nConfiguredVoltage    : 1200\r\nCreationClassName    : Win32_PhysicalMemory\r\nDataWidth            : 64\r\nDescription          : 物理内存\r\nDeviceLocator        : ChannelA-DIMM1\r\nFormFactor           : 8\r\nHotSwappable         : \r\nInstallDate          : \r\nInterleaveDataDepth  : 2\r\nInterleavePosition   : 1\r\nManufacturer         : Kingston\r\nMaxVoltage           : 1200\r\nMemoryType           : 0\r\nMinVoltage           : 1200\r\nModel                : \r\nName                 : 物理内存\r\nOtherIdentifyingInfo : \r\nPartNumber           : 9905678-012.A00G    \r\nPositionInRow        : \r\nPoweredOn            : \r\nRemovable            : \r\nReplaceable          : \r\nSerialNumber         : 04142A56\r\nSKU                  : \r\nSMBIOSMemoryType     : 26\r\nSpeed                : 2400\r\nStatus               : \r\nTag                  : Physical Memory 1\r\nTotalWidth           : 64\r\nTypeDetail           : 128\r\nVersion              : \r\nPSComputerName       : DESKTOP-48RD9C2\r\n\r\n__GENUS              : 2\r\n__CLASS              : Win32_PhysicalMemory\r\n__SUPERCLASS         : CIM_PhysicalMemory\r\n__DYNASTY            : CIM_ManagedSystemElement\r\n__RELPATH            : Win32_PhysicalMemory.Tag=\"Physical Memory 3\"\r\n__PROPERTY_COUNT     : 36\r\n__DERIVATION         : {CIM_PhysicalMemory, CIM_Chip, CIM_PhysicalComponent, CIM_PhysicalElement...}\r\n__SERVER             : DESKTOP-48RD9C2\r\n__NAMESPACE          : root\\cimv2\r\n__PATH               : \\\\DESKTOP-48RD9C2\\root\\cimv2:Win32_PhysicalMemory.Tag=\"Physical Memory 3\"\r\nAttributes           : 1\r\nBankLabel            : BANK 3\r\nCapacity             : 8589934592\r\nCaption              : 物理内存\r\nConfiguredClockSpeed : 2133\r\nConfiguredVoltage    : 1200\r\nCreationClassName    : Win32_PhysicalMemory\r\nDataWidth            : 64\r\nDescription          : 物理内存\r\nDeviceLocator        : ChannelB-DIMM1\r\nFormFactor           : 8\r\nHotSwappable         : \r\nInstallDate          : \r\nInterleaveDataDepth  : 2\r\nInterleavePosition   : 2\r\nManufacturer         : Kingston\r\nMaxVoltage           : 1200\r\nMemoryType           : 0\r\nMinVoltage           : 1200\r\nModel                : \r\nName                 : 物理内存\r\nOtherIdentifyingInfo : \r\nPartNumber           : 9905678-012.A00G    \r\nPositionInRow        : \r\nPoweredOn            : \r\nRemovable            : \r\nReplaceable          : \r\nSerialNumber         : 1E147F37\r\nSKU                  : \r\nSMBIOSMemoryType     : 26\r\nSpeed                : 2400\r\nStatus               : \r\nTag                  : Physical Memory 3\r\nTotalWidth           : 64\r\nTypeDetail           : 128\r\nVersion              : \r\nPSComputerName       : DESKTOP-48RD9C2\r\n\r\n\r\n\r\n\r\n\r\nManufacturer : Gigabyte Technology Co., Ltd.\r\nModel        : \r\nName         : 基板\r\nSerialNumber : Default string\r\nSKU          : \r\nProduct      : B250-HD3-CF\r\n\r\n\r\n\r\n\r\n\r\nSystemDirectory : C:\\WINDOWS\\system32\r\nOrganization    : P R C\r\nBuildNumber     : 19045\r\nRegisteredUser  : China\r\nSerialNumber    : 00391-80000-00001-AA710\r\nVersion         : 10.0.19045\r\n\r\n\r\n\r\n\r\n\r\nDomain              : WORKGROUP\r\nManufacturer        : Gigabyte Technology Co., Ltd.\r\nModel               : B250-HD3\r\nName                : DESKTOP-48RD9C2\r\nPrimaryOwnerName    : China\r\nTotalPhysicalMemory : 17135796224\r\n\r\n\r\n\r\n', '3天3夜');

-- ----------------------------
-- Table structure for jobdata
-- ----------------------------
DROP TABLE IF EXISTS `jobdata`;
CREATE TABLE `jobdata`  (
  `jobname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `jobstarttime` timestamp(0) NULL DEFAULT NULL,
  `jobstoptime` timestamp(0) NULL DEFAULT NULL,
  `jobinfo` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `jobrunning` bigint(20) NULL DEFAULT NULL,
  `joberr` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jobdata
-- ----------------------------
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:54:30', '2023-05-29 16:54:30', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:00', '2023-05-29 16:55:01', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:25', '2023-05-29 16:55:25', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:30', '2023-05-29 16:55:30', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:35', '2023-05-29 16:55:35', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:40', '2023-05-29 16:55:40', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:45', '2023-05-29 16:55:45', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:50', '2023-05-29 16:55:50', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:55:55', '2023-05-29 16:55:55', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:56:00', '2023-05-29 16:56:00', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:56:05', '2023-05-29 16:56:05', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:56:10', '2023-05-29 16:56:10', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:56:20', '2023-05-29 16:56:20', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:56:30', '2023-05-29 16:56:30', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:56:40', '2023-05-29 16:56:40', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:56:50', '2023-05-29 16:56:50', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:57:00', '2023-05-29 16:57:00', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:57:10', '2023-05-29 16:57:10', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:57:20', '2023-05-29 16:57:20', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:57:30', '2023-05-29 16:57:30', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:57:40', '2023-05-29 16:57:40', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:57:50', '2023-05-29 16:57:50', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:58:00', '2023-05-29 16:58:00', 'hello\n', 0, '');
INSERT INTO `jobdata` VALUES ('job1', '2023-05-29 16:58:10', '2023-05-29 16:58:10', 'hello\n', 0, '');

-- ----------------------------
-- Table structure for joblist
-- ----------------------------
DROP TABLE IF EXISTS `joblist`;
CREATE TABLE `joblist`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `jobid` bigint(20) NULL DEFAULT NULL COMMENT '任务ID\r\n',
  `jobname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务名称',
  `jobshell` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '任务shell',
  `jobstarttime` timestamp(0) NOT NULL COMMENT '任务添加时间',
  `jobstatus` int(10) NULL DEFAULT NULL COMMENT '任务状态',
  `jobcronexpr` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT 'cron表达式',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_jobname`(`jobname`) USING BTREE,
  INDEX `jobid`(`jobid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 36 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for notification
-- ----------------------------
DROP TABLE IF EXISTS `notification`;
CREATE TABLE `notification`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `cpuoption` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `memoryoption` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `systemdiskoption` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `thresholdstatus` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `workapiurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `workatuser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `dingapiurl` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `dingatuser` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of notification
-- ----------------------------
INSERT INTO `notification` VALUES (1, '80', '70', '60', '1', 'https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=72375e4f-7198-4002-8583-dcf905c6374c', '15901368529', 'https://oapi.dingtalk.com/robot/send?access_token=6695aa858422e9c30acda84a21d601c34e34c930d0797c8ff4fb31c82d0b3405', '15901368529');

-- ----------------------------
-- Table structure for systemlog
-- ----------------------------
DROP TABLE IF EXISTS `systemlog`;
CREATE TABLE `systemlog`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '序号',
  `systemlogid` bigint(20) NULL DEFAULT NULL COMMENT '系统事件id',
  `systemloghostname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '系统事件主机',
  `systemlogtype` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '系统事件类型',
  `systemloginfo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '系统事件信息',
  `systemlognote` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '系统事件备注',
  `systemlogstarttime` timestamp(0) NOT NULL ON UPDATE CURRENT_TIMESTAMP(0) COMMENT '系统事件开始时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_id`(`id`) USING BTREE,
  UNIQUE INDEX `idx_alarmid`(`systemlogid`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL,
  `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `gender` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL DEFAULT '0',
  `create_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0),
  `update_time` timestamp(0) NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_user_id`(`user_id`) USING BTREE,
  UNIQUE INDEX `idx_username`(`username`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (2, 2994544563810471936, 'Mark', '73696e666f74656ba7df66ea7576ba297964104c965a0200', NULL, '0', '2023-04-17 16:51:00', '2023-04-17 16:51:00');

SET FOREIGN_KEY_CHECKS = 1;
