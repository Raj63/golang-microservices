CREATE TABLE IF NOT EXISTS `investor_wallets` (
  `id` uuid NOT NULL,
  `investor_id` uuid NOT NULL,
  `balance` BIGINT DEFAULT 0,
  `currency_id` INT,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `investor_wallets_investor_id_IDX` (`investor_id`),
  CONSTRAINT uc_investor_wallets_investor_id UNIQUE (`investor_id`),
  FOREIGN KEY (`investor_id`) REFERENCES investors (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`currency_id`) REFERENCES currency (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
