CREATE TABLE `accounts` (`id` text,`created_by` text,`created_at` datetime,`domain` text,`domain_category` text,`is_domain_primary_account` numeric,`network_identifier` text,`network_net` text,`network_dns` text,`network_serial` integer,`dns_settings_disabled_management_groups` text,`settings_peer_login_expiration_enabled` numeric,`settings_peer_login_expiration` integer,`settings_regular_users_view_blocked` numeric,`settings_groups_propagation_enabled` numeric,`settings_jwt_groups_enabled` numeric,`settings_jwt_groups_claim_name` text,`settings_jwt_allow_groups` text,`settings_extra_peer_approval_enabled` numeric,`settings_extra_integrated_validator_groups` text,PRIMARY KEY (`id`));
CREATE TABLE `groups` (`id` text,`account_id` text,`name` text,`issued` text,`peers` text,`integration_ref_id` integer,`integration_ref_integration_type` text,PRIMARY KEY (`id`),CONSTRAINT `fk_accounts_groups_g` FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`));
CREATE TABLE `setup_keys` (`id` text,`account_id` text,`key` text,`key_secret` text,`name` text,`type` text,`created_at` datetime,`expires_at` datetime,`updated_at` datetime,`revoked` numeric,`used_times` integer,`last_used` datetime,`auto_groups` text,`usage_limit` integer,`ephemeral` numeric,PRIMARY KEY (`id`),CONSTRAINT `fk_accounts_setup_keys_g` FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`));
CREATE TABLE `peers` (`id` text,`account_id` text,`key` text,`setup_key` text,`ip` text,`meta_hostname` text,`meta_go_os` text,`meta_kernel` text,`meta_core` text,`meta_platform` text,`meta_os` text,`meta_os_version` text,`meta_wt_version` text,`meta_ui_version` text,`meta_kernel_version` text,`meta_network_addresses` text,`meta_system_serial_number` text,`meta_system_product_name` text,`meta_system_manufacturer` text,`meta_environment` text,`meta_files` text,`name` text,`dns_label` text,`peer_status_last_seen` datetime,`peer_status_connected` numeric,`peer_status_login_expired` numeric,`peer_status_requires_approval` numeric,`user_id` text,`ssh_key` text,`ssh_enabled` numeric,`login_expiration_enabled` numeric,`last_login` datetime,`created_at` datetime,`ephemeral` numeric,`location_connection_ip` text,`location_country_code` text,`location_city_name` text,`location_geo_name_id` integer,PRIMARY KEY (`id`),CONSTRAINT `fk_accounts_peers_g` FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`));

INSERT INTO accounts VALUES('testAccountId','','2024-10-02 16:01:38.000000000+00:00','test.com','private',1,'testNetworkIdentifier','{"IP":"100.64.0.0","Mask":"//8AAA=="}','',0,'[]',0,86400000000000,0,0,0,'',NULL,NULL,NULL);
INSERT INTO "groups" VALUES('testGroupId','testAccountId','testGroupName','api','[]',0,'');
INSERT INTO "groups" VALUES('newGroupId','testAccountId','newGroupName','api','[]',0,'');
INSERT INTO setup_keys VALUES('testKeyId','testAccountId','testKey','testK****','existingKey','one-off','2021-08-19 20:46:20.000000000+00:00','2321-09-18 20:46:20.000000000+00:00','2021-08-19 20:46:20.000000000+00:000',0,0,'0001-01-01 00:00:00+00:00','["testGroupId"]',1,0);
INSERT INTO setup_keys VALUES('revokedKeyId','testAccountId','revokedKey','testK****','existingKey','reusable','2021-08-19 20:46:20.000000000+00:00','2321-09-18 20:46:20.000000000+00:00','2021-08-19 20:46:20.000000000+00:00',1,0,'0001-01-01 00:00:00+00:00','["testGroupId"]',3,0);
INSERT INTO setup_keys VALUES('expiredKeyId','testAccountId','expiredKey','testK****','existingKey','reusable','2021-08-19 20:46:20.000000000+00:00','1921-09-18 20:46:20.000000000+00:00','2021-08-19 20:46:20.000000000+00:00',0,1,'0001-01-01 00:00:00+00:00','["testGroupId"]',5,1);
INSERT INTO peers VALUES('testPeerId','testAccountId','5rvhvriKJZ3S9oxYToVj5TzDM9u9y8cxg7htIMWlYAg=','72546A29-6BC8-4311-BCFC-9CDBF33F1A48','"100.64.114.31"','f2a34f6a4731','linux','Linux','11','unknown','Debian GNU/Linux','','0.12.0','','',NULL,'','','','{"Cloud":"","Platform":""}',NULL,'f2a34f6a4731','f2a34f6a4731','2023-03-02 09:21:02.189035775+01:00',0,0,0,'','ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAILzUUSYG/LGnV8zarb2SGN+tib/PZ+M7cL4WtTzUrTpk',0,1,'2023-03-01 19:48:19.817799698+01:00','2024-10-02 17:00:32.527947+02:00',0,'""','','',0);


CREATE TABLE `users` (`id` text,`account_id` text,`role` text,`is_service_user` numeric,`non_deletable` numeric,`service_user_name` text,`auto_groups` text,`blocked` numeric,`last_login` datetime,`created_at` datetime,`issued` text DEFAULT "api",`integration_ref_id` integer,`integration_ref_integration_type` text,PRIMARY KEY (`id`),CONSTRAINT `fk_accounts_users_g` FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`));

INSERT INTO users VALUES('testUserId','testAccountId','user',0,0,'','[]',0,'0001-01-01 00:00:00+00:00','2024-10-02 16:01:38.000000000+00:00','api',0,'');
INSERT INTO users VALUES('testAdminId','testAccountId','admin',0,0,'','[]',0,'0001-01-01 00:00:00+00:00','2024-10-02 16:01:38.000000000+00:00','api',0,'');
INSERT INTO users VALUES('testOwnerId','testAccountId','owner',0,0,'','[]',0,'0001-01-01 00:00:00+00:00','2024-10-02 16:01:38.000000000+00:00','api',0,'');
INSERT INTO users VALUES('testServiceUserId','testAccountId','user',1,0,'','[]',0,'0001-01-01 00:00:00+00:00','2024-10-02 16:01:38.000000000+00:00','api',0,'');
INSERT INTO users VALUES('testServiceAdminId','testAccountId','admin',1,0,'','[]',0,'0001-01-01 00:00:00+00:00','2024-10-02 16:01:38.000000000+00:00','api',0,'');
INSERT INTO users VALUES('blockedUserId','testAccountId','admin',0,0,'','[]',1,'0001-01-01 00:00:00+00:00','2024-10-02 16:01:38.000000000+00:00','api',0,'');
INSERT INTO users VALUES('otherUserId','otherAccountId','admin',0,0,'','[]',0,'0001-01-01 00:00:00+00:00','2024-10-02 16:01:38.000000000+00:00','api',0,'');
