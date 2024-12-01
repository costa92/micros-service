-- 创建订单表
CREATE TABLE IF NOT EXISTS `orders` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '自增主键ID',
    `order_id` varchar(64) NOT NULL COMMENT '订单ID',
    `user_id` bigint(20) NOT NULL COMMENT '用户ID',
    `payment_amount` decimal(10,2) NOT NULL DEFAULT '0.00' COMMENT '支付金额',
    `order_status` varchar(32) NOT NULL COMMENT '订单状态：1-未支付 2-已支付 3-已发货 4-已收货 5-已完成 6-已取消',
    `created_at` bigint(20) NOT NULL COMMENT '创建时间',
    `updated_at` bigint(20) NOT NULL COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_order_id` (`order_id`),
    KEY `idx_user_id` (`user_id`),
    KEY `idx_created_at` (`created_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='订单表';

-- 添加订单状态的检查约束（MySQL 8.0+支持）
-- ALTER TABLE `orders` 
-- ADD CONSTRAINT `chk_order_status` 
-- CHECK (`order_status` IN ('1', '2', '3', '4', '5', '6'));

-- 初始化订单状态数据（可选）
INSERT INTO `orders` (`order_id`, `user_id`, `payment_amount`, `order_status`, `created_at`, `updated_at`) 
VALUES 
('TEST_ORDER_001', 1, 99.99, '1', UNIX_TIMESTAMP(), UNIX_TIMESTAMP()); 