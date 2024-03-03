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
	"github.com/gorilla/sessions"
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
	var store = sessions.NewCookieStore([]byte(os.Getenv("SECRET")))

	var service = handler.Service{
		Storage:     storage,
		CookieStore: store,
	}

	router := mux.NewRouter().StrictSlash(true)
	dir := http.Dir("./assets")
	fs := http.StripPrefix("/assets/", http.FileServer(dir))
	router.PathPrefix("/assets/").Handler(fs) // static load

	authWare := ware.CheckAuth(service.Storage, service.CookieStore) // middleware for auth check

	router.HandleFunc("/", service.IndexHandler()).Methods(http.MethodGet, http.MethodPost)
	router.HandleFunc("/sub", service.SubHandler()).Methods(http.MethodPost)
	router.Handle("/prices", authWare(service.PricesHandler())).Methods(http.MethodGet)
	router.Handle("/prices/{id}", authWare(service.GoodHandler())).Methods(http.MethodGet)
	router.Handle("/profile", authWare(service.ProfileHandler())).Methods(http.MethodGet)
	//router.Handle("/cart", authWare(service.CartHandler()))

	auth := router.PathPrefix("/auth/").Subrouter()
	auth.HandleFunc("/signup", service.SignUpHandler()).Methods(http.MethodPost, http.MethodGet)
	auth.HandleFunc("/signin", service.SignInHandler()).Methods(http.MethodPost, http.MethodGet)

	//api := router.PathPrefix("/api/").Subrouter()

	log.Println("Server Satrt on " + os.Getenv("HOST"))
	defer log.Println("Stop Server")
	log.Fatal(http.ListenAndServe(os.Getenv("HOST"), gorillaH.LoggingHandler(os.Stdout, router)))

}
