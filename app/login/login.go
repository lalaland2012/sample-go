package login

import (
	"net/http"
	"sample/app/shared/handler"
	"sample/app/shared/helper"

	"github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Msg string
}

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// Login page
func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	sesValidate, _ := store.Get(r, "validate")
	_, existUserName := session.Values["userName"].(string)
	_, existPass := session.Values["pass"].(string)
	if existUserName && existPass {
		http.Redirect(w, r, "/hello", 302)
	}
	msg := ""
	if flashes := sesValidate.Flashes("Error"); len(flashes) > 0 {
		sesValidate.Save(r, w)
		msg = flashes[0].(string)
	}
	err := h.ResponseHTML(w, r, templateData{
		Msg: msg,
	}, "layout/base", "login/login", "layout/style", "layout/js")
	if err != nil {
		_ = h.StatusServerError(w, r)
	}
}

// LoginHandler post Login
func (h *HTTPHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	sesValidate, _ := store.Get(r, "validate")
	userName := r.FormValue("name")
	pass := r.FormValue("pass")
	redirectTarget := "/login"
	if !helper.IsEmpty(userName) && !helper.IsEmpty(pass) {
		userIsValid := CheckUser(userName, pass)
		if userIsValid {
			session, _ := store.Get(r, "user")
			session.Values["userName"] = userName
			session.Values["pass"] = pass
			session.Save(r, w)
			redirectTarget = "/hello"
		} else {
			sesValidate.AddFlash("Incorrect name and password", "Error")
			sesValidate.Save(r, w)
			redirectTarget = "/login"
		}
	} else {
		sesValidate.AddFlash("Please, input name and password!", "Error")
		sesValidate.Save(r, w)
	}
	http.Redirect(w, r, redirectTarget, 302)

}

//LogoutHandler handler logout
func (h *HTTPHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

// CheckUser check user is valid
func CheckUser(userName, pass string) bool {
	uName, pwd, isValid := "manhnd", "123456", false
	if uName == userName && pwd == pass {
		isValid = true
	} else {
		isValid = false
	}
	return isValid
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
