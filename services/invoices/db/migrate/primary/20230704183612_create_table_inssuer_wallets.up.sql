CREATE TABLE IF NOT EXISTS `issuer_wallets` (
  `id` uuid NOT NULL,
  `issuer_id` uuid NOT NULL,
  `balance` BIGINT DEFAULT 0,
  `currency_id` INT,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `issuer_wallets_issuer_id_IDX` (`issuer_id`),
  CONSTRAINT uc_issuer_wallets_issuer_id UNIQUE (`issuer_id`),
  FOREIGN KEY (`issuer_id`) REFERENCES issuers (`id`) ON DELETE CASCADE,
  FOREIGN KEY (`currency_id`) REFERENCES currency (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
