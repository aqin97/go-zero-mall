CREATE TABLE `pay` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `oid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '订单id',
    `amount` int unsigned NOT NULL DEFAULT '0' COMMENT '产品金额',
    `source` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '支付方式',
    `status` tinyint unsigned NOT NULL DEFAULT '0' COMMENT '支付状态',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `idx_uid` (`uid`),
    KEY `idx_oid` (`oid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;