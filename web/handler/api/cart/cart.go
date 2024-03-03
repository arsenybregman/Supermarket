package cart

import (
	"Supermarket/sql"
	"encoding/json"
	"net/http"
)

func ApiCart(store *sql.Storage) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
			case http.MethodGet:
				
			case http.MethodPost:
				var cart sql.Goods
				err := json.NewDecoder(r.Body).Decode(&cart)
				if err != nil{
					http.Error(w, "Server error", http.StatusInternalServerError)
				}
		}
	}
}