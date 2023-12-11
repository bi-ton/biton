CREATE TABLE IF NOT EXISTS main.products (
    id    serial PRIMARY KEY,
    name  text NOT NULL,
    price real NOT NULL,
    done  bool NOT NULL
);
