package sql

import (
	"Supermarket/internal"
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DB struct
type Storage struct {
	db *sql.DB
}

type Goods struct {
	Id          int         `json:"id"`
	Title       string      `json:"title"`
	Price       float64     `json:"price"`
	Description string      `json:"description"`
	Quantity    interface{} `json:"quantity"`
}

type User struct {
	Name         string `validate:"required"`
	Surname      string `validate:"required"`
	Email        string `validate:"email,required"`
	Password     string `validate:"required,min=8,eqfield=ConfPassword"`
	ConfPassword string `validate:"required,min=8"`
}

type UserLogin struct {
	Email    string `validate:"email,required"`
	Password string `validate:"required,min=8"`
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
			quantity    INTEGER
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
	rows, err := db.db.Query("SELECT id, title, price, description, quantity FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Обрабатываем каждую запись
	for rows.Next() {
		good := Goods{}
		err := rows.Scan(&good.Id, &good.Title, &good.Price, &good.Description, &good.Quantity)
		if err == nil {
			res = append(res, good)
		}
	}
	return res, nil
}

func (db *Storage) GetGood(id string) (Goods, error) {
	var good Goods
	row := db.db.QueryRow("SELECT title, price, description, quantity FROM products WHERE id=?", id)
	err := row.Scan(&good.Title, &good.Price, &good.Description, &good.Quantity)
	if err != nil {
		return good, err
	}
	return good, nil
}

func (db *Storage) CreateUser(u User) error {
	data, err := db.db.Prepare("INSERT INTO users (name, surname, email, password) VALUES(?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err1 := data.Exec(u.Name, u.Surname, u.Email, internal.Hash([]byte(u.Password)))
	if err1 != nil {
		return err1
	}
	return nil
}

func (db *Storage) GetUser(email  string) (User, error) {
	var user User
	err := db.db.QueryRow("SELECT name, surname, email FROM users WHERE email = ?", email).Scan(&user.Name, &user.Surname, &user.Email)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (db *Storage) CheckAuthUser(u UserLogin) (bool, error) {
	var id int
	err := db.db.QueryRow("SELECT id FROM users WHERE email=? AND password=?", u.Email, internal.Hash([]byte(u.Password))).Scan(&id)
	if err != nil {
		return false, err
	}
	return true, nil
}
