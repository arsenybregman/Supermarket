package main

import (
	"log"
	"net/http"
	"os"

	"Supermarket/handler"

	"Supermarket/sql"

	gorillaH "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	var err = godotenv.Load() // load env
	if  err != nil {
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

	router.HandleFunc("/", handler.IndexHandler(storage)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/sub", handler.SubHandler(storage)).Methods(http.MethodPost)
	router.HandleFunc("/prices", handler.PricesHandler(storage)).Methods(http.MethodGet)
	router.HandleFunc("/profile", handler.ProfileHandler(storage)).Methods(http.MethodGet)

	log.Println("Server Satrt")
	defer log.Println("Stop Server")
	log.Fatal(http.ListenAndServe(os.Getenv("HOST"), gorillaH.LoggingHandler(os.Stdout, router)))

}
