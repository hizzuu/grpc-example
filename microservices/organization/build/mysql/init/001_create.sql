CREATE TABLE `organizations` (
  `id`         INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name`       VARCHAR(255) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL
);

CREATE TABLE `relations` (
  `id`              INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `role`            VARCHAR(255) NOT NULL,
  `user_id`         INT(11) NOT NULL,
  `organization_id` INT(11) NOT NULL,
  `created_at`      DATETIME NOT NULL,
  `updated_at`      DATETIME NOT NULL
);
