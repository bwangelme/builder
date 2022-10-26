create table record
(
    `id`       int unsigned not null primary key,
    `status`   tinyint unsigned not null,
    `giturl`   varchar(100) not null,
    `revision` varchar(100) not null,
    `img`      varchar(100) not null,
    `tag`      varchar(100) not null
);
