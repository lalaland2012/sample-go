package hello

import (
	"net/http"
	"sample/app/shared/auth"
	"sample/app/shared/handler"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	Name string
}

// HelloWorld hello word page
func (h *HTTPHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	user := auth.User(w, r)
	err := h.ResponseHTML(w, r, templateData{
		Name: user["userName"],
	}, "layout/base", "hello/hello_world")
	if err != nil {
		_ = h.StatusServerError(w, r)
	}
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
