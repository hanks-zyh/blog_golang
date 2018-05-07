CREATE TABLE t_user(
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(20) NOT NULL,
  `pwd` VARCHAR (30) NOT NULL,
  `phone` VARCHAR (30),
  `authData` VARCHAR (200),
  `createdAt` BIGINT,
  `updatedAt` BIGINT,
  PRIMARY KEY (`id`)
);

CREATE TABLE t_article(
  `id` BIGINT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(20) NOT NULL,
  `content` VARCHAR (30) NOT NULL,
  `createdAt` BIGINT,
  `updatedAt` BIGINT,
  `uid` BIGINT,
  PRIMARY KEY (`id`)
);


SELECT * FROM t_user;

DELETE FROM t_user WHERE id>0;