CREATE TABLE IF NOT EXISTS `investors` (
  `id` uuid NOT NULL,
  `name` VARCHAR NOT NULL,
  `vat_number` VARCHAR NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  INDEX `investors_name_IDX` (`name`),
  INDEX `investors_vat_number_IDX` (`vat_number`),
  CONSTRAINT uc_investors_vat_number UNIQUE (`vat_number`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
