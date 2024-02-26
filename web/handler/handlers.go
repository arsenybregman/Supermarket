package handler

import (
	"Supermarket/sql"

	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
)

type Message struct {
	Email string
	Name  string
	Text  string
}

type GoodsAnswer struct {
	Prices []sql.Goods
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

func SignUpHandler(storage *sql.Storage, sessionsStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "templates/reg.html")
		case http.MethodPost:
			r.ParseForm()
			validate := validator.New()
			user := sql.User{
				Name:         r.FormValue("name"),
				Surname:      r.FormValue("surname"),
				Email:        r.FormValue("email"),
				Password:     r.FormValue("password"),
				ConfPassword: r.FormValue("confirm-password"),
			}
			err := validate.Struct(&user)
			if err != nil {
				http.Error(w, "Failed to create user.", http.StatusBadRequest)
				return
			}

			err = storage.CreateUser(user)
			if err != nil {
				http.Error(w, "Failed to create user.", http.StatusBadRequest)
				return
			}

			var session, _ = sessionsStore.Get(r, "auth")
			session.Values["check"] = true
			session.Options.MaxAge = 86400 * 7 // sec
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/profile", http.StatusFound)
		default:
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	}
}

func SignInHandler(storage *sql.Storage, sessionsStore *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			http.ServeFile(w, r, "templates/login.html")
		case http.MethodPost:
			r.ParseForm()
			validate := validator.New()
			user := sql.UserLogin{
				Email:    r.FormValue("email"),
				Password: r.FormValue("password"),
			}
			err := validate.Struct(&user)
			if err != nil {
				http.ServeFile(w, r, "templates/login.html")
			}
			check, _ := storage.CheckAuthUser(user)
			if !check {
				http.Error(w, "Failed to find user.", http.StatusBadRequest)
				return
			}
			var session, _ = sessionsStore.Get(r, "auth")
			session.Values["check"] = true
			session.Options.MaxAge = 86400 * 7 // sec
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/profile", http.StatusFound)
		}
	}
}
