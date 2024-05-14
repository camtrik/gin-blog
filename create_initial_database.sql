-- Create database
CREATE DATABASE IF NOT EXISTS blog_service 
DEFAULT CHARACTER SET utf8mb4 
DEFAULT COLLATE utf8mb4_general_ci;

-- Use the created database
USE blog_service;

-- Create common fields table
CREATE TABLE IF NOT EXISTS `common_fields` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'Creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modification time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Deletion time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'Is deleted: 0 means not deleted, 1 means deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Common fields';


-- Create tag table
CREATE TABLE IF NOT EXISTS `blog_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT 'Tag name',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'Creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modification time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Deletion time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'Is deleted: 0 means not deleted, 1 means deleted',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'State: 0 means disabled, 1 means enabled',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Tag management';

-- Create article table
CREATE TABLE IF NOT EXISTS `blog_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT 'Article title',
  `desc` varchar(255) DEFAULT '' COMMENT 'Article description',
  `cover_image_url` varchar(255) DEFAULT '' COMMENT 'Cover image URL',
  `content` longtext COMMENT 'Article content',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'Creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modification time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Deletion time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'Is deleted: 0 means not deleted, 1 means deleted',
  `state` tinyint(3) unsigned DEFAULT '1' COMMENT 'State: 0 means disabled, 1 means enabled',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Article management';

-- Create article-tag association table
CREATE TABLE IF NOT EXISTS `blog_article_tag` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `article_id` int(11) NOT NULL COMMENT 'Article ID',
  `tag_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT 'Tag ID',
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'Creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modification time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Deletion time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'Is deleted: 0 means not deleted, 1 means deleted',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Article-Tag association';
