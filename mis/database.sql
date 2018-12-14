DROP TABLE IF EXISTS `record`;
CREATE TABLE `record` (
  `id` INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  `title` VARCHAR(60),
  `address` VARCHAR(120),
  `lat` VARCHAR(10),
  `lng` VARCHAR(10),
  `agent` VARCHAR(200),
  `house` VARCHAR(1000),
  `remark` VARCHAR(1000)
);
