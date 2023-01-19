create table `userinfo`(
    `id` integer primary key autoincrement,
    `username` varchar(50) not null,
    `password` varchar(128) not null,
    `create_time` date not null
);