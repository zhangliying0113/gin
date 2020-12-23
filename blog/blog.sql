create database `blog`;

use `blog`;

DROP TABLE IF EXISTS `article`;
CREATE TABLE `article` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '文章ID',
  `category_id` BIGINT(20) NOT NULL COMMENT '分类ID',
  `content` LONGTEXT NOT NULL COMMENT '文章内容',
  `title` varchar(50) NOT NULL DEFAULT '' COMMENT '文章标题',
  `view_count` INT(11) NOT NULL DEFAULT 0 COMMENT '阅读次数',
  `comment_count` INT(11) NOT NULL DEFAULT 0 COMMENT '评论次数',
  `username` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '作者',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '状态 正常为1 删除为0',
  `summary` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '文章摘要',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='文章表';

DROP TABLE IF EXISTS `category`;
CREATE TABLE `category` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '分类ID',
  `category_name` VARCHAR(255) NOT NULL COMMENT '分类名字',
  `category_no` INT(10) NOT NULL COMMENT '分类排序',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类表';

DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '评论ID',
  `content` TEXT NOT NULL COMMENT '评论内容',
  `username` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '评论作者',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '评论发布时间',
  `status` TINYINT(1) NOT NULL DEFAULT 1 COMMENT '评论状态 正常为1 删除为0',
  `article_id` BIGINT(20) NOT NULL DEFAULT 0 COMMENT '文章ID',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论表';

DROP TABLE IF EXISTS `leave`;
CREATE TABLE `leave` (
  `id` BIGINT(20) NOT NULL AUTO_INCREMENT COMMENT '留言ID',
  `username` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '留言作者',
  `email` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '留言作者',
  `content` TEXT NOT NULL COMMENT '留言内容',
  `create_time` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '留言发布时间',
  `update_time` TIMESTAMP COMMENT '留言更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='留言表';