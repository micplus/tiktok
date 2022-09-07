CREATE TABLE videos (
    id bigint NOT NULL AUTO_INCREAMENT,
    title varchar(64) NOT NULL,
    play_url text NOT NULL,
    cover_url text NOT NULL,
    created_at bigint NOT NULL,
    modified_at bigint NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY(`id`),
    INDEX(user_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE users (
    id bigint NOT NULL AUTO_INCREAMENT
    name varchar(32),
    created_at bigint NOT NULL,
    modified_at bigint NOT NULL,
    PRIMARY KEY(`id`),
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE user_logins (
    id bigint NOT NULL AUTO_INCREAMENT,
    username varchar(32) NOT NULL,
    password varchar(32) NOT NULL,
    salt varchar(8) NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY(`id`),
    UNIQUE INDEX(username)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE user_favorites (
    id bigint NOT NULL AUTO_INCREAMENT,
    user_id bigint NOT NULL,
    video_id bigint NOT NULL,
    PRIMARY KEY(`id`),
    INDEX(user_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE user_follows (
    id bigint NOT NULL AUTO_INCREAMENT,
    user_id bigint NOT NULL,
    follow_id bigint NOT NULL,
    PRIMARY KEY(`id`),
    INDEX(user_id),
    INDEX(follow_id),
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE comments (
    id bigint NOT NULL AUTO_INCREAMENT,
    content text NOT NULL,
    create_date char(5),
    created_at bigint NOT NULL,
    modified_at bigint NOT NULL,
    video_id bigint NOT NULL,
    user_id bigint NOT NULL,
    PRIMARY KEY(`id`),
    INDEX(video_id)
)ENGINE=InnoDB DEFAULT CHARSET=utf8;
