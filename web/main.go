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
	"github.com/gorilla/sessions"
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
	var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET")))

	router := mux.NewRouter()
	dir := http.Dir("./assets")
	fs := http.StripPrefix("/assets/", http.FileServer(dir))
	router.PathPrefix("/assets/").Handler(fs) // static load

	authWare := ware.CheckAuth(storage, store) // middleware for auth check

	router.HandleFunc("/", handler.IndexHandler(storage)).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/sub", handler.SubHandler(storage)).Methods(http.MethodPost)
	router.HandleFunc("/prices", handler.PricesHandler(storage)).Methods(http.MethodGet)
	router.Handle("/profile", authWare(handler.ProfileHandler(storage))).Methods(http.MethodGet)
	
	auth := router.PathPrefix("/auth/").Subrouter()
	auth.HandleFunc("/signup", handler.SignUpHandler(storage, store)).Methods(http.MethodPost, http.MethodGet)
	auth.HandleFunc("/signin", handler.SignInHandler(storage, store)).Methods(http.MethodPost, http.MethodGet)

	log.Println("Server Satrt on " + os.Getenv("HOST"))
	defer log.Println("Stop Server")
	log.Fatal(http.ListenAndServe(os.Getenv("HOST"), gorillaH.LoggingHandler(os.Stdout, router)))

}
