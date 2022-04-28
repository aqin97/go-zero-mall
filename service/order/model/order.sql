CREATE TABLE `order` (
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    `uid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '用户id',
    `pid` bigint unsigned NOT NULL DEFAULT '0' COMMENT '产品id',
    `amount` int unsigned NOT NULL DEFAULT '0' COMMENT '订单金额',
    `status` int unsigned NOT NULL DEFAULT '0' COMMENT '订单状态',
    `create_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`),
    KEY `idx_uid` (`uid`),
    KEY `idx_pid` (`pid`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;