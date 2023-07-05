CREATE TABLE IF NOT EXISTS `invoices` (
  `id` uuid NOT NULL,
  `number` varchar NOT NULL,
  `status` varchar NOT NULL,
  `description` varchar,
  `amount` BIGINT  DEFAULT 0 NOT NULL,
  `currency_id` INT,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`currency_id`) REFERENCES currency (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
