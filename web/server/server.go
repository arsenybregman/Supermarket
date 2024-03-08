package server

import (
	"Supermarket/sql"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"log"
	"os"
)
type Service struct {
	Storage     *sql.Storage
	CookieStore *sessions.CookieStore
}

func NewService() Service{
	var err = godotenv.Load() // load env
	if err != nil {
		log.Fatal(err)
	}
	var storage, errS = sql.NewStorage() // load db
	if errS != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}
	var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET")))

	return Service{
		Storage:     storage,
		CookieStore:  store,
	}
}