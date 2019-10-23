package hello

import (
	"net/http"
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
	err := h.ResponseHTML(w, r, "hello/hello_world", templateData{
		Name: "tttt",
	})
	if err != nil {
		_ = h.StatusServerError(w, r)
	}
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
