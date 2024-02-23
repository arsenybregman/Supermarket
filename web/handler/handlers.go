package handler

import (
	"Supermarket/sql"
	"fmt"

	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type Message struct {
	Email string
	Name  string
	Text  string
}

type GoodsAnswer struct {
	Prices []sql.Goods
}

type User struct {
	Name         string `validate:"required"`
	Surname      string `validate:"required"`
	Email string `validate:"email,required"`
	Password     string `validate:"required,min=8,eqfield=ConfPassword"`
	ConfPassword string `validate:"required,min=8"`
}

// main page
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

// sub page
func SubHandler(storage *sql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		storage.CreateSub(r.FormValue("email"))
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}
}

// show all prices
func PricesHandler(storage *sql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		v, err := storage.GetGoods()

		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		tmpl, _ := template.ParseFiles("templates/main.html")
		tmpl.Execute(w, v)
	}
}

// profile
func ProfileHandler(storage *sql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/profile.html")
		tmpl.Execute(w, "")
	}
}

func SignUpHandler(storage *sql.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "templates/reg.html")
		case http.MethodPost:
			r.ParseForm()
			validate := validator.New()
			user := User{
				Name:     r.FormValue("name"),
				Surname:  r.FormValue("surname"),
				Email: r.FormValue("email"),
				Password: r.FormValue("password"),
				ConfPassword: r.FormValue("confirm-password"),
			}
			fmt.Println(user)
			err := validate.Struct(&user)
			if err != nil {
				http.NotFound(w, r)
				return
			}
			storage.CreateUser(user.Name, user.Surname, user.Email, user.Password)
			http.Redirect(w, r, "/profile", http.StatusFound)
		}
	}
}
