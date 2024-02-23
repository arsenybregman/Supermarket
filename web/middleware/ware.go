package ware

import (
	"Supermarket/sql"
	"net/http"
)

func CheckAuth(storage *sql.Storage) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		})
	}
}
