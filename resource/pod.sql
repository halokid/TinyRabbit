
SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for pod
-- ----------------------------
DROP TABLE IF EXISTS `pod`;
CREATE TABLE `pod` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `did` int(11) NOT NULL,
  `name` varchar(256) NOT NULL,
  `eid` int(11) NOT NULL,
  `ename` varchar(256) NOT NULL,
  `addr` varchar(256) NOT NULL,
  `env` varchar(64) NOT NULL,
  `cid` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=153 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of pod
-- ----------------------------
INSERT INTO `pod` VALUES ('46', '32', 'xx-svc', '10', 'test-endpoint', '8.8.8.8:9100', 'gwx', 'cd4e215e458c3299cdeb4f26b1ff4be4cd277d88cf712aebab940f236842621e');
