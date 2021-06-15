
-- ----------------------------------------------------------------------------------------------------------
-- 文章
-- ----------------------------------------------------------------------------------------------------------


CREATE TABLE `article` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `type_id` int(11) DEFAULT '0' COMMENT '分类ID',
  `user_id` int(11) DEFAULT '0' COMMENT '用户ID',
  `label_id` int(11) DEFAULT '0' COMMENT '标签ID',
  `title` varchar(256) DEFAULT '' COMMENT '标题',
  `image` varchar(256) DEFAULT '' COMMENT '封面图片',
  `describe` varchar(256) DEFAULT '' COMMENT '描述',
  `content` text COMMENT '内容',
  `sort` int(11) DEFAULT '0' COMMENT '排序（从大到小）',
  `url` varchar(512) DEFAULT '' COMMENT '原文地址',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态 1：正常 2：禁用',
  `sign` tinyint(4) DEFAULT '2' COMMENT '置顶 1：是 2：否',
  `thumbs` int(11) DEFAULT '0' COMMENT '点赞',
  `hiss` int(11) DEFAULT '0' COMMENT '嘘声',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_type_id` (`type_id`),
  KEY `idx_user_id` (`user_id`),
  KEY `idx_label_id` (`label_id`),
  KEY `idx_sort` (`sort`),
  KEY `idx_status` (`status`),
  KEY `idx_sign` (`sign`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';

