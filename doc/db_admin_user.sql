
-- ----------------------------------------------------------------------------------------------------------
-- 管理员表
-- ----------------------------------------------------------------------------------------------------------


CREATE TABLE `admin_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `nickname` varchar(64) DEFAULT '' COMMENT '昵称',
  `username` varchar(64) DEFAULT '' COMMENT '用户名',
  `password` varchar(512) DEFAULT '' COMMENT '密码',
  `token` varchar(64) DEFAULT '' COMMENT 'Token',
  `token_over_at` int(11) DEFAULT '0' COMMENT 'Token结束时间',
  `status` tinyint(4) DEFAULT '1' COMMENT '状态 1：正常 2：禁用',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_username` (`username`(10))
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员表';
