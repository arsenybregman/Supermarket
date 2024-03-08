package handler

import (
	"Supermarket/sql"

	"html/template"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

type Service struct {
	Storage     *sql.Storage
	CookieStore *sessions.CookieStore
}

type Message struct {
	Email string
	Name  string
	Text  string
}

type GoodsAnswer struct {
	Prices []sql.Goods
}

// main page
func (s Service) IndexHandler() http.HandlerFunc {
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
			s.Storage.CreateForm(form.Email, form.Name, form.Text)

			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

	}
}

// sub page
func (s Service) SubHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		s.Storage.CreateSub(r.FormValue("email"))
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

// show all prices
func (s Service) PricesHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/main.html")
	}
}

// profile
func (s Service) ProfileHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/profile2.html")
		tmpl.Execute(w, "")
	}
}

func (s Service) SignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tmpl, _ := template.ParseFiles("templates/reg.html")
			tmpl.Execute(w, "")
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
				tmpl, _ := template.ParseFiles("templates/reg.html")
				tmpl.Execute(w, "Ошибка ввода данных")
				return
			}

			err = s.Storage.CreateUser(user)
			if err != nil {
				tmpl, _ := template.ParseFiles("templates/reg.html")
				tmpl.Execute(w, "Данный пользователь существует")
				return
			}

			var session, _ = s.CookieStore.Get(r, "auth")
			session.Values["check"] = true
			session.Values["email"] = r.FormValue("email")
			session.Options.MaxAge = 86400 * 7 // sec
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
		default:
			http.Error(w, "Bad request", http.StatusBadRequest)
		}
	}
}

func (s Service) SignInHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			tmpl, _ := template.ParseFiles("templates/login.html")
			tmpl.Execute(w, "")
		case http.MethodPost:
			r.ParseForm()
			validate := validator.New()
			user := sql.UserLogin{
				Email:    r.FormValue("email"),
				Password: r.FormValue("password"),
			}
			err := validate.Struct(&user)
			if err != nil {
				tmpl, _ := template.ParseFiles("templates/login.html")
				tmpl.Execute(w, "Ошибка ввода данных")
				return
			}
			check, _ := s.Storage.CheckAuthUser(user)
			if !check {
				tmpl, _ := template.ParseFiles("templates/login.html")
				tmpl.Execute(w, "Данного пользователя не существует")
				return
			}
			var session, _ = s.CookieStore.Get(r, "auth")
			session.Values["check"] = true
			session.Values["email"] = r.FormValue("email")
			session.Options.MaxAge = 86400 * 7 // sec
			err = session.Save(r, w)
			if err != nil {
				http.Error(w, "Server error", http.StatusInternalServerError)
				return
			}

			http.Redirect(w, r, "/", http.StatusSeeOther)
			return

		}
	}
}

func (s Service) GoodHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		good := mux.Vars(r)
		row, err := s.Storage.GetGood(good["id"])
		if err != nil {
			http.Error(w, "Server error", http.StatusInternalServerError)
			return
		}
		tmpl, _ := template.ParseFiles("templates/good.html")
		tmpl.Execute(w, row)
	}
}

func (s Service) CartHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:

		case http.MethodPost:
		}
	}
}
