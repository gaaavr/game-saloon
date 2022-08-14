-- +goose Up
-- +goose StatementBegin
CREATE TABLE users
(
    id serial not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null,
    role varchar(255) not null,
    ppm int,
    money int,
    dead boolean,
    last_drink timestamp
);
CREATE TABLE drinks
(
    id serial unique not null,
    name varchar(255) not null,
    price int,
    alcohol int
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
DROP TABLE drinks;
-- +goose StatementEnd
