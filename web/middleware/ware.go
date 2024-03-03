package ware

import (
	"Supermarket/sql"

	"net/http"

	"github.com/gorilla/sessions"
)

func CheckAuth(storage *sql.Storage, store *sessions.CookieStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			session, _ := store.Get(req, "auth")
			if session.Values["check"] == true {
				next.ServeHTTP(w, req)
				return
			} else {
				http.Redirect(w, req, "/auth/signup", http.StatusSeeOther)
				return
			}

		})
	}
}
