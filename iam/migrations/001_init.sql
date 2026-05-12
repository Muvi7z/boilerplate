-- +goose Up
CREATE TABLE IF NOT EXISTS users
(
    uuid     text primary key,
    email    text,
    login    text,
    password text
);

CREATE TABLE IF NOT EXISTS notification_methods
(
    id            serial primary key,
    user_uuid     text not null,
    provider_name text,
    target        text,
    FOREIGN KEY (user_uuid) REFERENCES users (uuid)
);

-- +goose Down
drop table if exists users;
drop table if exists notification_methods;