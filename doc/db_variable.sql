
-- ----------------------------------------------------------------------------------------------------------
-- 变量表
-- ----------------------------------------------------------------------------------------------------------


CREATE TABLE `variable` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(128) DEFAULT '' COMMENT '名称',
  `desc` varchar(256) DEFAULT '' COMMENT '描述 ',
  `value` TEXT COMMENT '值',
  `create_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_at` int(11) DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_name` (`name`(10))
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='变量表';
