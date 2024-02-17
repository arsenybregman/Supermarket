package handler

import (
	"Supermarket/sql"

	"html/template"
	"net/http"

)

type Message struct {
	Email string
	Name  string
	Text  string
}

func IndexHandler(storage *sql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tmpl, _ := template.ParseFiles("templates/index.html")
			tmpl.Execute(w, "")
		case http.MethodPost:
			r.ParseForm()
			form := Message{
				Email: r.FormValue("email"),
				Name:  r.FormValue("name"),
				Text:  r.FormValue("textarea"),
			}

			storage.CreateForm(form.Email, form.Name, form.Text)

			http.Redirect(w, r, "/", http.StatusMovedPermanently)
		}

	}
}

func SubHandler(storage *sql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		storage.CreateSub(r.FormValue("email"))
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}
