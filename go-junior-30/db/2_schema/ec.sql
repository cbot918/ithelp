-- setup
CREATE TABLE IF NOT EXISTS users(
  id serial PRIMARY KEY,
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL,
  password varchar(255) NOT NULL,
  discount float NOT NULL
);

CREATE TABLE IF NOT EXISTS orders(
  id serial PRIMARY KEY,
  order_by int NOT NULL,
  items int[] NOT NULL,
  total_price int NOT NULL,
  discount float NOT NULL,
  FOREIGN KEY (order_by) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS items(
  id serial PRIMARY KEY,
  available int NOT NULL,
  name varchar(255) NOT NULL,
  vender varchar(255) NOT NULL,
  unit_price int NOT NULL
);

-- teardown
ALTER TABLE IF EXISTS "orders" DROP CONSTRAINT IF EXISTS "orders_order_by_fkey";

DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS items;
DROP TABLE IF EXISTS users;