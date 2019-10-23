package logout

import (
	"net/http"
	"sample/app/shared/handler"

	"github.com/gorilla/sessions"
)

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	UserName string
}

func (h *HTTPHandler) HandleLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/login", 302)
}

// NewHTTPLogoutHandler responses new HTTPLogoutHandler instance.
func NewLogoutHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
