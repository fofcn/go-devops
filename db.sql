create table `userinfo`(
    `id` integer primary key autoincrement,
    `username` varchar(50) not null,
    `password` varchar(128) not null,
    `create_time` datetime not null
);

create table `media_file`(
    `id` integer primary key autoincrement,
    `file_name` varchar(128) not null, 
    `store_path` varchar(128) not null,
    `file_create_time` datetime not null,
    `media_type` tinyint not null,
    `create_time` datetime not null
);