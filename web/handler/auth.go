package handler

import (
	"net/http"
)

func (s Service) SetAuthSession(w http.ResponseWriter, r *http.Request) error {
	var session, _ = s.CookieStore.Get(r, "auth")
	session.Values["check"] = true
	session.Values["email"] = r.FormValue("email")
	session.Options.MaxAge = 86400 * 7 // sec
	err := session.Save(r, w)
	return err
}
