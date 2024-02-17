package sql

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Storage struct {
	db *sql.DB
}

func NewStorage() (*Storage, error) {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		return nil, err
	}
	CreateTables(db)
	return &Storage{db}, nil
}

func CreateTables(db *sql.DB) {
	data, err := db.Prepare(`CREATE TABLE IF NOT EXISTS form (
		name    TEXT NOT NULL,
		email   TEXT NOT NULL,
		message TEXT NOT NULL
	);`)

	if err != nil {
		log.Fatal(err)
	}
	data.Exec()
}

func (db *Storage) CreateForm(email, name, message string) error {
	data, err := db.db.Prepare(
		`INSERT INTO form (email, name, message) VALUES(?, ?, ?) `)
	if err != nil {
		return err
	}
	_, err = data.Exec(email, name, message)
	if err != nil{
		return err
	}
	return nil
}

func (db *Storage) CreateSub(email string) error {
	data, err := db.db.Prepare(`INSERT INTO sub (email) VALUES(?)`)
	if err != nil{
		return err
	}
	_, err = data.Exec(email)
	if err != nil{
		return err
	}
	return nil
}