create table if not exists tiktok.t_comments
(
    id           int auto_increment
    primary key,
    user_id      int                                not null,
    video_id     int                                not null,
    content      varchar(100)                       not null,
    gmt_created  datetime default CURRENT_TIMESTAMP not null,
    gmt_modified datetime default CURRENT_TIMESTAMP not null
    );

create index t_comments_video_id_index
    on tiktok.t_comments (video_id);

create table if not exists tiktok.t_follow_lists
(
    id           int auto_increment
    primary key,
    user_id      int                                not null,
    follower_id  int                                not null,
    status       tinyint(1)                         not null,
    gmt_created  datetime default CURRENT_TIMESTAMP not null,
    gmt_modified datetime default CURRENT_TIMESTAMP not null,
    constraint t_follow_lists_pk
    unique (user_id, follower_id)
    );

create table if not exists tiktok.t_liked_videos
(
    id           int auto_increment
    primary key,
    user_id      int                                not null,
    video_id     int                                not null,
    gmt_created  datetime default CURRENT_TIMESTAMP not null,
    gmt_modified datetime default CURRENT_TIMESTAMP not null,
    constraint t_liked_videos_pk
    unique (user_id, video_id)
    );

create table if not exists tiktok.t_messages
(
    id           int auto_increment
    primary key,
    content      varchar(2048)                      not null,
    user_id      int                                not null,
    to_user_id   int                                not null,
    send_time    datetime default CURRENT_TIMESTAMP not null,
    gmt_created  datetime default CURRENT_TIMESTAMP not null,
    gmt_modified datetime default CURRENT_TIMESTAMP not null
    );

create index t_messages_user_id_to_user_id_send_time_index
    on tiktok.t_messages (user_id, to_user_id, send_time);

create table if not exists tiktok.t_user_infos
(
    id             int auto_increment
    primary key,
    username       varchar(32)                        not null,
    password       char(60)                           not null,
    follow_count   int      default 0                 not null,
    follower_count int      default 0                 not null,
    gmt_created    datetime default CURRENT_TIMESTAMP not null,
    gmt_modified   datetime default CURRENT_TIMESTAMP not null,
    constraint t_user_infos_pk
    unique (username)
    );

create table if not exists tiktok.t_video_infos
(
    id             int auto_increment
    primary key,
    author_id      int                                not null,
    publish_time   datetime default CURRENT_TIMESTAMP not null,
    cover_url      varchar(255)                       not null,
    play_url       varchar(255)                       not null,
    favorite_count int      default 0                 not null,
    comment_count  int      default 0                 not null,
    title          varchar(32)                        null,
    gmt_created    datetime default CURRENT_TIMESTAMP not null,
    gmt_modified   datetime default CURRENT_TIMESTAMP not null
    );

