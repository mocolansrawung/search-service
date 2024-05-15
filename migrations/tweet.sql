CREATE TABLE `tweet` (
    `id` char(36) NOT NULL,
    `user_id` bigint NOT NULL,
    `product_id` char(36) DEFAULT NULL,
    `model_id` bigint DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
  INDEX `idx_user_id` (`user_id`),
  INDEX `idx_product_id` (`product_id`),
  INDEX `idx_model_id` (`model_id`),
  INDEX `idx_user_product` (`user_id`,`product_id`)
  INDEX `idx_user_model` (`user_id`,`model_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;