package ware

import (
	"Supermarket/sql"
	"net/http"
	"github.com/gorilla/sessions"
)

func CheckAuth(storage *sql.Storage, store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			session, err := store.Get(req, "auth")
			if  err != nil {
				http.Redirect(w, req, "/auth/signin", http.StatusPermanentRedirect)
				return
			}
			if session.Values["check"] == true {
				next.ServeHTTP(w, req)
				return
			}
			http.Redirect(w, req, "/auth/signin", http.StatusPermanentRedirect)
		})
	}
}
