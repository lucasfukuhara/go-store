
CREATE DATABASE go_store;

CREATE TABLE IF NOT EXISTS products
(
    id serial NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    description varchar,
    price numeric,
    quantity integer
);