CREATE SCHEMA IF NOT EXISTS main;

CREATE TABLE IF NOT EXISTS main.products (
    id    serial PRIMARY KEY,
    name  text NOT NULL,
    price real NOT NULL,
    done  bool NOT NULL DEFAULT false
);
