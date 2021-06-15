
-- ----------------------------------------------------------------------------------------------------------
-- 文章评论
-- ----------------------------------------------------------------------------------------------------------


CREATE TABLE `article_comment` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `article_id` int(11) DEFAULT '0' COMMENT '文章ID',
  `comment_id` int(11) DEFAULT '0' COMMENT '评论ID',
  `comment_user_nickname`  varchar(256) DEFAULT '' COMMENT '被评论用户昵称',
  `user_nickname` varchar(256) DEFAULT '' COMMENT '用户昵称',
  `user_email` varchar(256) DEFAULT '' COMMENT '用户邮箱',
  `content` text COMMENT '评论内容',
  `admin_status` tinyint(4) DEFAULT '2' COMMENT '是否管理员 1：是 2：否',
  `status` tinyint(4) DEFAULT '2' COMMENT '状态 1：正常 2：禁用',
  `sign` tinyint(4) DEFAULT '2' COMMENT '置顶 1：是 2：否',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_article_id` (`article_id`),
  KEY `idx_comment_id` (`comment_id`),
  KEY `idx_status` (`status`),
  KEY `idx_sign` (`sign`),
  KEY `idx_admin_status` (`admin_status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章评论表';

