USE blog_service;

CREATE TABLE IF NOT EXISTS `blog_auth` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `app_key` varchar(20) DEFAULT '' COMMENT 'Key',
  `app_secret` varchar(50) DEFAULT '' COMMENT 'Secret',  
  `created_on` int(10) unsigned DEFAULT '0' COMMENT 'Creation time',
  `created_by` varchar(100) DEFAULT '' COMMENT 'Creator',
  `modified_on` int(10) unsigned DEFAULT '0' COMMENT 'Modification time',
  `modified_by` varchar(100) DEFAULT '' COMMENT 'Modifier',
  `deleted_on` int(10) unsigned DEFAULT '0' COMMENT 'Deletion time',
  `is_del` tinyint(3) unsigned DEFAULT '0' COMMENT 'Is deleted: 0 means not deleted, 1 means deleted',

  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Authentication management';

INSERT INTO `blog_service`.`blog_auth`(`id`, `app_key`, `app_secret`, `created_on`, `created_by`, `modified_on`, `modified_by`, `deleted_on`, `is_del`) VALUES (1, 'ebbi', 'gin-blog', 0, 'ebbi', 0, '', 0, 0);
