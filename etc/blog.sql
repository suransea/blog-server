CREATE TABLE IF NOT EXISTS `article`
(
    id           INT AUTO_INCREMENT,
    title        VARCHAR(100) NOT NULL,
    summary      VARCHAR(300) NOT NULL,
    content_id   INT          NOT NULL,
    created_time DATETIME     NOT NULL,
    updated_time DATETIME     NOT NULL,
    CONSTRAINT article_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `content`
(
    id           INT AUTO_INCREMENT,
    content      TEXT     NOT NULL,
    created_time DATETIME NOT NULL,
    updated_time DATETIME NOT NULL,
    CONSTRAINT content_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `comment`
(
    id           INT AUTO_INCREMENT,
    article_id   INT          NOT NULL,
    nickname     VARCHAR(50)  NOT NULL,
    content      VARCHAR(500) NOT NULL,
    created_time DATETIME     NOT NULL,
    updated_time DATETIME     NOT NULL,
    CONSTRAINT comment_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `tag`
(
    id           INT AUTO_INCREMENT,
    tag          VARCHAR(50) NOT NULL,
    created_time DATETIME    NOT NULL,
    updated_time DATETIME    NOT NULL,
    CONSTRAINT tag_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `article_tag`
(
    id           INT AUTO_INCREMENT,
    article_id   INT      NOT NULL,
    tag_id       INT      NOT NULL,
    created_time DATETIME NOT NULL,
    updated_time DATETIME NOT NULL,
    CONSTRAINT article_tag_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `user`
(
    id           INT AUTO_INCREMENT,
    username     VARCHAR(50) NOT NULL,
    password     VARCHAR(100) NOT NULL,
    created_time DATETIME    NOT NULL,
    updated_time DATETIME    NOT NULL,
    CONSTRAINT user_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `role`
(
    id           INT AUTO_INCREMENT,
    role         VARCHAR(50) NOT NULL,
    created_time DATETIME    NOT NULL,
    updated_time DATETIME    NOT NULL,
    CONSTRAINT role_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `user_role`
(
    id           INT AUTO_INCREMENT,
    user_id      INT      NOT NULL,
    role_id      INT      NOT NULL,
    created_time DATETIME NOT NULL,
    updated_time DATETIME NOT NULL,
    CONSTRAINT user_role_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `user_article`
(
    id           INT AUTO_INCREMENT,
    user_id      INT      NOT NULL,
    article_id   INT      NOT NULL,
    created_time DATETIME NOT NULL,
    updated_time DATETIME NOT NULL,
    CONSTRAINT user_article_pk
        PRIMARY KEY (id)
) ENGINE = innodb
  DEFAULT CHARSET = utf8;
