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
	router.HandleFunc("/", handler.IndexHandler).Methods(http.MethodGet, http.MethodPost)
	log.Fatal(http.ListenAndServe(":80", gorillaH.LoggingHandler(os.Stdout, router)))

}
