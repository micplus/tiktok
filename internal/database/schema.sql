USE tiktok;

CREATE TABLE videos (
    id bigint AUTO_INCREMENT NOT NULL,
    title varchar(64) NOT NULL DEFAULT "Title",
    play_url text NOT NULL,
    cover_url text NOT NULL,
    created_at bigint NOT NULL,
    modified_at bigint NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX videosuserid ON videos(`user_id`);

CREATE TABLE users (
    id bigint AUTO_INCREMENT NOT NULL,
    name varchar(32),
    created_at bigint NOT NULL,
    modified_at bigint NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE user_logins (
    id bigint AUTO_INCREMENT NOT NULL,
    username varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    salt varchar(8) NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE UNIQUE INDEX userloginsusername ON user_logins(`username`);

CREATE TABLE user_favorites (
    id bigint AUTO_INCREMENT NOT NULL,
    user_id bigint NOT NULL,
    video_id bigint NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX userfavoritesuserid ON user_favorites(`user_id`);

CREATE TABLE user_follows (
    id bigint NOT NULL AUTO_INCREMENT,
    user_id bigint NOT NULL,
    follow_id bigint NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX userfollowsuserid ON user_follows(`user_id`);
CREATE INDEX userfollowsfollowid ON user_follows(`follow_id`);

CREATE TABLE comments (
    id bigint NOT NULL AUTO_INCREMENT,
    content text NOT NULL,
    create_date varchar(10),
    created_at bigint NOT NULL,
    modified_at bigint NOT NULL,
    video_id bigint NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY(`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE INDEX commentsvideoid ON comments(`video_id`);
