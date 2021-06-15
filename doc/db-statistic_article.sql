
-- ----------------------------------------------------------------------------------------------------------
-- 文章统计
-- ----------------------------------------------------------------------------------------------------------
CREATE TABLE `statistic_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `article_id` int(11) DEFAULT '0' COMMENT '博客ID',
  `ip` varchar(256) DEFAULT '' COMMENT 'IP地址',
  `address` varchar(256) DEFAULT '' COMMENT '城市',
  `create_date` date COMMENT '创建日期',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_ip` (`ip`),
  KEY `idx_address` (`address`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_create_date` (`create_date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章统计记录表';
