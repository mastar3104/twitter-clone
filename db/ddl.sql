CREATE TABLE user
(
    user_id   VARCHAR(63) NOT NULL PRIMARY KEY,
    user_name VARCHAR(63) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    deleted BOOLEAN NOT NULL DEFAULT false,
    created_at  DATETIME    NOT NULL,
    updated_at  DATETIME    NOT NULL
) ENGINE=InnoDB COMMENT 'ユーザ情報';

CREATE TABLE tweet
(
    tweet_id   VARCHAR(63) NOT NULL PRIMARY KEY,
    user_id VARCHAR(63) NOT NULL,
    content VARCHAR(140) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL ,
    created_at  DATETIME    NOT NULL,
    updated_at  DATETIME    NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
) ENGINE=InnoDB COMMENT 'ユーザの投稿情報';

CREATE TABLE follow
(
    user_id   VARCHAR(63) NOT NULL,
    follow_user_id VARCHAR(63) NOT NULL,
    created_at  DATETIME    NOT NULL,
    updated_at  DATETIME    NOT NULL,
    PRIMARY KEY (user_id, follow_user_id),
    FOREIGN KEY (follow_user_id) REFERENCES user(user_id),
    FOREIGN KEY (user_id) REFERENCES user(user_id)
) ENGINE=InnoDB COMMENT 'ユーザがフォローしているユーザとのリレーション';

-- 高速読み取りのためにキャッシュ化なども検討すべき
CREATE TABLE timeline
(
    tweet_id   VARCHAR(63) NOT NULL,
    user_id VARCHAR(63) NOT NULL,
    created_at  DATETIME    NOT NULL,
    updated_at  DATETIME    NOT NULL,
    PRIMARY KEY (user_id, tweet_id)
) ENGINE=InnoDB COMMENT 'ユーザのタイムラインに表示する投稿のID';


