CREATE TABLE cats(
    id        text NOT NULL,
    name      text NOT NULL,
    age       int  NOT NULL
);

CREATE TABLE users(
    id serial not null unique,
    username varchar(255) not null unique,
    password varchar(255) not null
)