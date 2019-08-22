package login

import (
	"net/http"
	"sample/app/shared/auth"
	"sample/app/shared/handler"
	"sample/app/shared/session"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Msg string
}

// Login page
func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	sesUser := session.GetStore(w, r, "user")
	userName, _ := sesUser.Values["userName"].(string)
	pass, _ := sesUser.Values["pass"].(string)
	if auth.Attempt(userName, pass) {
		http.Redirect(w, r, "/hello", 302)
		return
	}
	msg := session.GetMsgFlash(w, r, "Error")
	err := h.ResponseHTML(w, r, templateData{
		Msg: msg,
	}, "layout/base", "login/login", "layout/style", "layout/js")
	if err != nil {
		_ = h.StatusServerError(w, r)
	}
}

// LoginHandler post Login
func (h *HTTPHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("name")
	pass := r.FormValue("pass")
	redirectTarget := "/login"
	if auth.Login(userName, pass, w, r) {
		redirectTarget = "/hello"
	}
	http.Redirect(w, r, redirectTarget, 302)
}

//LogoutHandler handler logout
func (h *HTTPHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	auth.Logout(w, r)
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
