
-- ----------------------------------------------------------------------------------------------------------
-- 友情链接
-- ----------------------------------------------------------------------------------------------------------


CREATE TABLE `friendship` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(256) DEFAULT '' COMMENT '网站名',
  `url` varchar(256) DEFAULT '' COMMENT '网站地址',
  `describe` varchar(256) DEFAULT '' COMMENT '描述',
  `sort` int(11) DEFAULT '0' COMMENT '排序（从大到小）',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态 1：正常 2：禁用',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_title` (`title`),
  KEY `idx_sort` (`sort`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='友情链接';

