-- Drop tables if they exist
DROP TABLE IF EXISTS `acl_qr`;
DROP TABLE IF EXISTS `api_token`;
DROP TABLE IF EXISTS `sso_integration_mapping`;

-- Create sso_integration_mapping table
CREATE TABLE `sso_integration_mapping` (
  `id` int(11) NOT NULL auto_increment,
  `ssousername` varchar(255) NOT NULL,
  `aclusername` varchar(80) NOT NULL DEFAULT '',
  PRIMARY KEY  (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create api_token table
CREATE TABLE `api_token` (
  `id` int(11) NOT NULL auto_increment,
  `username` varchar(80) NOT NULL,
  `token` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `token` (`token`),
  KEY `username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create acl_qr table
CREATE TABLE `acl_qr` (
  `id` int(11) NOT NULL auto_increment,
  `uuid` varchar(255) NOT NULL,
  `token` varchar(255) NOT NULL,
  `device_name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `token` (`token`),
  CONSTRAINT `fk_acl_qr_token` FOREIGN KEY (`token`) REFERENCES `api_token` (`token`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
