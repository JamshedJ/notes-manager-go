CREATE TABLE users
(
    id            serial       not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE tasks
(
    id          serial       not null unique,
    user_id     int          references users (id) on delete cascade  not null,
    title       varchar(255) not null,
    description varchar(255),
    done        boolean      not null default false
);