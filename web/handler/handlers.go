package handler

import (
	"Supermarket/sql"
	sqlGo "database/sql"
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
		session, _ := s.CookieStore.Get(r, "auth")
		emailCookie, ok := session.Values["email"].(string)
		if !ok {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		user, err := s.Storage.GetUser(emailCookie)
		if err != nil {
			if err == sqlGo.ErrNoRows {
				http.Error(w, "Такого пользователя не существует", http.StatusInternalServerError)
				return
			}
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		tmpl, _ := template.ParseFiles("templates/profile2.html")
		tmpl.Execute(w, user)
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

			err = s.SetAuthSession(w, r)
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
			
			err = s.SetAuthSession(w, r)
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
			tmpl, _ := template.ParseFiles("templates/basket.html")
			tmpl.Execute(w, "")
		case http.MethodPost:
		}
	}
}
