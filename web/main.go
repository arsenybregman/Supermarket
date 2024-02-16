package main

import (
	"log"
	"net/http"
	"os"

	"Supermarket/handler"

	gorillaH "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"Supermarket/sql"
)

func main() {
	var storage, err = sql.NewStorage()
	if  err != nil {
		log.Fatalf("Could not connect to database: %s", err)
	}

	router := mux.NewRouter()
	dir := http.Dir("./assets")
	fs := http.StripPrefix("/assets", http.FileServer(dir))
	router.PathPrefix("/assets/").Handler(fs) //static load

	router.HandleFunc("/", handler.IndexHandler(storage)).Methods(http.MethodGet, http.MethodPost)
	log.Println("Server Satrt")
	log.Fatal(http.ListenAndServe("127.0.0.1:8000", gorillaH.LoggingHandler(os.Stdout, router)))

}
