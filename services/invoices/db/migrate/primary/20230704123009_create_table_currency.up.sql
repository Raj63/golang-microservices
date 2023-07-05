CREATE TABLE IF NOT EXISTS `currency` (
  `id` INT auto_increment NOT NULL,
  `name`   VARCHAR,
  `code`   VARCHAR(3),
  `symbol` VARCHAR(5),
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
PRIMARY KEY (`id`)
)ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;