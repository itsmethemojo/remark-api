CREATE DATABASE IF NOT EXISTS `remark`;

CREATE TABLE IF NOT EXISTS `remark`.`users` (
 `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'auto incrementing user_id of each user, unique index',
 `user_name` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT 'user''s name, unique',
 `user_password_hash` varchar(255) COLLATE utf8_unicode_ci NOT NULL COMMENT 'user''s password in salted and hashed format',
 `user_email` varchar(64) COLLATE utf8_unicode_ci NOT NULL COMMENT 'user''s email, unique',
 `user_active` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'user''s activation status',
 `user_activation_hash` varchar(40) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'user''s email verification hash string',
 `user_password_reset_hash` char(40) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'user''s password reset code',
 `user_password_reset_timestamp` bigint(20) DEFAULT NULL COMMENT 'timestamp of the password reset request',
 `user_rememberme_token` varchar(64) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'user''s remember-me cookie token',
 `user_failed_logins` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'user''s failed login attemps',
 `user_last_failed_login` int(10) DEFAULT NULL COMMENT 'unix timestamp of last failed login attempt',
 `user_registration_datetime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
 `user_registration_ip` varchar(39) COLLATE utf8_unicode_ci NOT NULL DEFAULT '0.0.0.0',
 PRIMARY KEY (`user_id`),
 UNIQUE KEY `user_name` (`user_name`),
 UNIQUE KEY `user_email` (`user_email`)
) ENGINE=MyISAM AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci COMMENT='user data';

CREATE TABLE IF NOT EXISTS `remark`.`bookmark` (
  `id` int(11) NOT NULL auto_increment,
  `user_id` int(11) NOT NULL,
  `url` varchar(400) character set latin1 NOT NULL,
  `domain` varchar(30) character set latin1 NOT NULL,
  `extension` varchar(10) character set latin1 NOT NULL,
  `title` varchar(200) character set latin1 NOT NULL,
  `customtitle` varchar(200) character set latin1 NOT NULL,
  `bookmarkcount` int(11) NOT NULL,
  `clickcount` int(11) NOT NULL,
  `tagcount` int(11) NOT NULL,
  `mediatype` int(11) NOT NULL,
  `created` bigint(20) NOT NULL,
  `updated` bigint(20) NOT NULL,
  `clicked` bigint(20) NOT NULL,
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=431 ;

CREATE TABLE IF NOT EXISTS `remark`.`bookmarktime` (
  `id` bigint(20) NOT NULL auto_increment,
  `user_id` int(11) NOT NULL,
  `bookmark_id` int(11) NOT NULL,
  `created` bigint(20) NOT NULL,
  UNIQUE KEY `id` (`id`),
  KEY `bookmark_id` (`bookmark_id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 AUTO_INCREMENT=523 ;

CREATE TABLE IF NOT EXISTS `remark`.`clicktime` (
  `id` bigint(20) NOT NULL default '0',
  `user_id` int(11) NOT NULL,
  `bookmark_id` int(11) NOT NULL,
  `created` bigint(20) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE IF NOT EXISTS `remark`.`tagtime` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `bookmark_id` int(11) NOT NULL,
  `name` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


