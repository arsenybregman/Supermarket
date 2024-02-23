package sql

import (
	"Supermarket/internal"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

type Goods struct {
	Title       string
	Price       int
	Description string
	Image       []byte
}


type DDL []string

func NewStorage() (*Storage, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	CreateTables(db)
	return &Storage{db}, nil
}

func CreateTables(db *sql.DB) {

	ddl := DDL{`CREATE TABLE IF NOT EXISTS form (
		name    TEXT NOT NULL,
		email   TEXT NOT NULL,
		message TEXT NOT NULL
	);`,
		`CREATE TABLE IF NOT EXISTS products (
		id          INTEGER PRIMARY KEY UNIQUE NOT NULL,
		title       TEXT    NOT NULL UNIQUE,
		price       NUMERIC NOT NULL,
		description TEXT    NOT NULL,
		image       BLOB
	);`,
		`CREATE TABLE IF NOT EXISTS sub (
		email TEXT PRIMARY KEY NOT NULL UNIQUE
	);`,
		`CREATE TABLE IF NOT EXISTS users (
		id       INTEGER PRIMARY KEY NOT NULL UNIQUE,
		name     TEXT    NOT NULL,
		surname  TEXT    NOT NULL,
		email    TEXT    UNIQUE NOT NULL,
		password TEXT    NOT NULL
	);`,
		`CREATE TABLE IF NOT EXISTS orders (
		order_id INTEGER PRIMARY KEY NOT NULL,
		user_id  INTEGER REFERENCES users (id) ON DELETE CASCADE NOT NULL,
		date     DATE    NOT NULL
	);`,
		`CREATE TABLE IF NOT EXISTS orders_products (
		order_id    INTEGER REFERENCES orders (order_id) ON DELETE CASCADE NOT NULL,
		products_id INTEGER REFERENCES products (id) ON DELETE CASCADE NOT NULL,
		quantity    INTEGER NOT NULL
	);`}

	for i := 0; i < len(ddl); i++ {
		pc, errP := db.Prepare(ddl[i]) // prepare statement but do not execute it yet
		if errP != nil {
			log.Fatal("Error preparing ddl: ", errP)
		}
		_, err := pc.Exec()
		if err != nil {
			log.Fatal("Error creating table: ", err)
		}
	}
}

func (db *Storage) CreateForm(email, name, message string) error {
	data, err := db.db.Prepare(
		`INSERT INTO form (email, name, message) VALUES(?, ?, ?) `)
	if err != nil {
		return err
	}
	_, err = data.Exec(email, name, message)
	if err != nil {
		return err
	}
	return nil
}

func (db *Storage) CreateSub(email string) error {
	data, err := db.db.Prepare(`INSERT INTO sub (email) VALUES(?)`)
	if err != nil {
		return err
	}
	_, err = data.Exec(email)
	if err != nil {
		return err
	}
	return nil
}

func (db *Storage) GetGoods() ([]Goods, error) {
	var res []Goods
	rows, err := db.db.Query("SELECT title, price, description, image FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Обрабатываем каждую запись
	for rows.Next() {
		good := Goods{}
		err := rows.Scan(&good.Title, &good.Price, &good.Description, &good.Image)
		if err == nil {
			res = append(res, good)
		}
	}
	return res, nil
}

func (db *Storage) CreateUser(name, surname, email, password string) error {
	data, err := db.db.Prepare("INSERT INTO users (name, surname, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err1 := data.Exec(name, surname, email, internal.Hash([]byte(password)))
	if err1 != nil {
		return err1
	}
	return nil
}
