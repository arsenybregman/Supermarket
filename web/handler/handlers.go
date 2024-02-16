package handler

import (
	"html/template"
	"net/http"
	"Supermarket/sql"
)

type Message struct {
	Email string
	Name  string
	Text  string
}

func IndexHandler(storage *sql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			tmpl, _ := template.ParseFiles("templates/index.html")
			tmpl.Execute(w, "")
		} else if r.Method == http.MethodPost {
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
