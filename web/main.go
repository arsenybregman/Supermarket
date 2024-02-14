package main

import (
	"log"
	"net/http"
	"os"

	"Supermarket/handler"

	gorillaH "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	dir := http.Dir("./assets")
	fs := http.StripPrefix("/assets", http.FileServer(dir))
	router.PathPrefix("/assets/").Handler(fs) //static load

	router.HandleFunc("/", handler.IndexHandler).Methods(http.MethodGet, http.MethodPost)
	log.Println("Server Satrt")
	log.Fatal(http.ListenAndServe(":80", gorillaH.LoggingHandler(os.Stdout, router)))

}
