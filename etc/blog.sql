CREATE TABLE IF NOT EXISTS `article`
(
    id         BIGINT AUTO_INCREMENT,
    title      VARCHAR(100) NOT NULL,
    summary    VARCHAR(300) NOT NULL,
    content_id BIGINT       NOT NULL,
    ctime      BIGINT       NOT NULL,
    utime      BIGINT       NOT NULL,
    CONSTRAINT article_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `content`
(
    id      BIGINT AUTO_INCREMENT,
    content TEXT   NOT NULL,
    ctime   BIGINT NOT NULL,
    utime   BIGINT NOT NULL,
    CONSTRAINT content_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `comment`
(
    id         BIGINT AUTO_INCREMENT,
    article_id BIGINT       NOT NULL,
    nickname   VARCHAR(50)  NOT NULL,
    content    VARCHAR(500) NOT NULL,
    ctime      BIGINT       NOT NULL,
    utime      BIGINT       NOT NULL,
    CONSTRAINT comment_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `tag`
(
    id    BIGINT AUTO_INCREMENT,
    tag   VARCHAR(50) NOT NULL,
    ctime BIGINT      NOT NULL,
    utime BIGINT      NOT NULL,
    CONSTRAINT tag_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `article_tag`
(
    id         BIGINT AUTO_INCREMENT,
    article_id BIGINT NOT NULL,
    tag_id     BIGINT NOT NULL,
    ctime      BIGINT NOT NULL,
    utime      BIGINT NOT NULL,
    CONSTRAINT article_tag_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `user`
(
    id       BIGINT AUTO_INCREMENT,
    username VARCHAR(50)  NOT NULL,
    password VARCHAR(100) NOT NULL,
    ctime    BIGINT       NOT NULL,
    utime    BIGINT       NOT NULL,
    CONSTRAINT user_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `role`
(
    id    BIGINT AUTO_INCREMENT,
    role  VARCHAR(50) NOT NULL,
    ctime BIGINT      NOT NULL,
    utime BIGINT      NOT NULL,
    CONSTRAINT role_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `user_role`
(
    id      BIGINT AUTO_INCREMENT,
    user_id BIGINT NOT NULL,
    role_id BIGINT NOT NULL,
    ctime   BIGINT NOT NULL,
    utime   BIGINT NOT NULL,
    CONSTRAINT user_role_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `user_article`
(
    id         BIGINT AUTO_INCREMENT,
    user_id    BIGINT NOT NULL,
    article_id BIGINT NOT NULL,
    ctime      BIGINT NOT NULL,
    utime      BIGINT NOT NULL,
    CONSTRAINT user_article_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;
