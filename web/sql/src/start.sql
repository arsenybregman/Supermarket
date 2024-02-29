CREATE TABLE IF NOT EXISTS form (
		name    TEXT NOT NULL,
		email   TEXT NOT NULL,
		message TEXT NOT NULL
	);
CREATE TABLE IF NOT EXISTS products (
		id          INTEGER PRIMARY KEY UNIQUE NOT NULL,
		title       TEXT    NOT NULL UNIQUE,
		price       NUMERIC NOT NULL,
		description TEXT    NOT NULL,
		quantity    INTEGER,
		image       BLOB
	);
CREATE TABLE IF NOT EXISTS sub (
		email TEXT PRIMARY KEY NOT NULL UNIQUE
	);
CREATE TABLE IF NOT EXISTS users (
		id       INTEGER PRIMARY KEY NOT NULL UNIQUE,
		name     TEXT    NOT NULL,
		surname  TEXT    NOT NULL,
		email    TEXT    UNIQUE NOT NULL,
		password TEXT    NOT NULL
	);
CREATE TABLE IF NOT EXISTS orders (
		order_id INTEGER PRIMARY KEY NOT NULL,
		user_id  INTEGER REFERENCES users (id) ON DELETE CASCADE NOT NULL,
		date     DATE    NOT NULL
	);
CREATE TABLE IF NOT EXISTS orders_products (
		order_id    INTEGER REFERENCES orders (order_id) ON DELETE CASCADE NOT NULL,
		products_id INTEGER REFERENCES products (id) ON DELETE CASCADE NOT NULL,
		quantity    INTEGER NOT NULL
	);