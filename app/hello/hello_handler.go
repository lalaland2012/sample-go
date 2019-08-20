package hello

import (
	"net/http"
	"sample/app/shared/handler"

	"github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Name string
	Pass string
}

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// HelloWorld hello word page
func (h *HTTPHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	userName, existUserName := session.Values["userName"].(string)
	pass, existPass := session.Values["pass"].(string)
	if existUserName && existPass {
		err := h.ResponseHTML(w, r, templateData{
			Name: userName,
			Pass: pass,
		}, "layout/base.tmpl", "hello/hello_world.tmpl")
		if err != nil {
			_ = h.StatusServerError(w, r)
		}
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
