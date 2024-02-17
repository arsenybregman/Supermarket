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
	var err = godotenv.Load()
	if  err != nil {
		log.Fatal(err)
	}
	
	var storage, errS = sql.NewStorage()
	if errS != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	router := mux.NewRouter()
	dir := http.Dir("./assets")
	fs := http.StripPrefix("/assets", http.FileServer(dir))
	router.PathPrefix("/assets/").Handler(fs) //static load

	router.HandleFunc("/", handler.IndexHandler(storage)).Methods(http.MethodGet, http.MethodPost)
	log.Println("Server Satrt")
	log.Fatal(http.ListenAndServe(os.Getenv("HOST"), gorillaH.LoggingHandler(os.Stdout, router)))

}
