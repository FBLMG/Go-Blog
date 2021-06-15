
-- ----------------------------------------------------------------------------------------------------------
-- 文章标签
-- ----------------------------------------------------------------------------------------------------------

CREATE TABLE `article_label` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `title` varchar(225) DEFAULT '' COMMENT '标题',
  `describe` varchar(256) DEFAULT '' COMMENT '描述',
  `sort` int(11) DEFAULT '0' COMMENT '排序（从小到大）',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态 1：正常 2：禁用',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE `idx_un_title` (`title`),
  KEY `idx_title` (`title`),
  KEY `idx_sort` (`sort`),
  KEY `idx_status` (`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章标签';

