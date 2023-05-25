CREATE TABLE users
(
    id            serial       not null unique,
    group_id serial not null,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    PRIMARY KEY (id)
);

CREATE TABLE posts
(
    id serial not null unique,
    user_id int references users (id) on delete no action on update no action not null,
    title varchar(255) not null,
    description text not null
);

CREATE TABLE comments
(
    id serial not null unique,
    user_id int references users (id) on delete no action on update no action not null,
    user_name varchar(255) references users (username) on delete no action on update no action not null,
    post_id int references posts (id) on delete cascade on update no action not null,
    content text not null
);