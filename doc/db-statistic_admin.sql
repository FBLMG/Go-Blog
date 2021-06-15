
-- ----------------------------------------------------------------------------------------------------------
-- 文章统计
-- ----------------------------------------------------------------------------------------------------------
CREATE TABLE `statistic_admin` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` int(11) DEFAULT '0' COMMENT '管理员ID',
  `ip` varchar(256) DEFAULT '' COMMENT 'IP地址',
  `address` varchar(256) DEFAULT '' COMMENT '城市',
  `create_date` date COMMENT '创建日期',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_ip` (`ip`),
  KEY `idx_address` (`address`),
  KEY `idx_admin_id` (`admin_id`),
  KEY `idx_create_date` (`create_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员登陆-统计记录表';
