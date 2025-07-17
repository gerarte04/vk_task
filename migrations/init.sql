CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    login           text NOT NULL UNIQUE,
    password_hash   bytea NOT NULL
);

CREATE INDEX ON users (login);

CREATE TABLE ads (
    id              uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    author_login    text,

    title           text NOT NULL CHECK (LENGTH(title) > 0),
    description     text,
    image_address   text,
    price           int CHECK (price > 0),

    FOREIGN KEY (author_login) REFERENCES users (login) ON DELETE SET NULL ON UPDATE CASCADE
);
