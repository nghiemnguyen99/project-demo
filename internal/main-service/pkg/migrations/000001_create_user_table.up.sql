CREATE TABLE `users` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `first_name`  varchar(255) NOT NULL,
  `last_name`  varchar(255) NOT NULL,
  `mssv`  varchar(10) NOT NULL,
  `subject_id` int NOT NULL,
  PRIMARY KEY (`id`)
);
