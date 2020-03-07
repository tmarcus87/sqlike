USE `library`;

CREATE TABLE `library`.`book`
(
    `id`        BIGINT(20) NOT NULL AUTO_INCREMENT,
    `title`     VARCHAR(300),
    `author_id` BIGINT(20) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;

CREATE TABLE `library`.`author`
(
    `id`   BIGINT(20) NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(300),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB;
