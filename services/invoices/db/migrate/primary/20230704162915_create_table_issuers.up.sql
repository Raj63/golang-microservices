CREATE TABLE IF NOT EXISTS `issuers` (
  `id` uuid NOT NULL,
  `name` VARCHAR NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `investors_name_IDX` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
