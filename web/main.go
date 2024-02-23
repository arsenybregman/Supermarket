package main

import (
	"log"
	"net/http"
	"os"

	"Supermarket/handler"
	ware "Supermarket/middleware"

	"Supermarket/sql"

	gorillaH "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	var err = godotenv.Load() // load env
	if err != nil {
		log.Fatal(err)
	}

	var storage, errS = sql.NewStorage() // load db
	if errS != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	router := mux.NewRouter()
	dir := http.Dir("./assets")
	fs := http.StripPrefix("/assets/", http.FileServer(dir))
	router.PathPrefix("/assets/").Handler(fs) // static load

	auth := ware.CheckAuth(storage) // middleware for auth check

	router.HandleFunc("/", handler.IndexHandler(storage)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/sub", handler.SubHandler(storage)).Methods(http.MethodPost)
	router.HandleFunc("/prices", handler.PricesHandler(storage)).Methods(http.MethodGet)
	router.Handle("/profile", auth(handler.ProfileHandler(storage))).Methods(http.MethodGet)
	router.HandleFunc("/signup", handler.SignUpHandler(storage)).Methods(http.MethodPost, http.MethodGet)

	log.Println("Server Satrt on " + os.Getenv("HOST"))
	defer log.Println("Stop Server")
	log.Fatal(http.ListenAndServe(os.Getenv("HOST"), gorillaH.LoggingHandler(os.Stdout, router)))

}
