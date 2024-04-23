package handler

import (
	"Supermarket/internal/api"
	"encoding/json"
	"strconv"

	"net/http"
)

func (s Service) ApiPrices() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			res := make(map[string]interface{})
			query, err := s.Storage.GetGoods()
			
			if err != nil {
				json.NewEncoder(w).Encode(api.ErrorMessage("Internal server error: " + err.Error()))
				return
			}

			for _, v := range query {
				res[strconv.Itoa(v.Id)] = v
			}
			RespondJSON(w, res)
		}
	}
}
