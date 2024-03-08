package handler

import (
	"Supermarket/internal/api"
	"encoding/json"
	"strconv"

	"net/http"
)

func (s Service) ApiPricesForJS() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			res := make(map[string]interface{})
			query, err := s.Storage.GetGoodsApi()

			if err != nil {
				json.NewEncoder(w).Encode(api.ErrorMessage("Internal server error: " + err.Error()))
				return
			}

			for _, v := range query.Products {
				res["product"+strconv.Itoa(v.Id)] = v
			}
			RespondJSON(w, res)
		}
	}
}
