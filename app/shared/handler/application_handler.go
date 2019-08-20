package handler

import (
	"html/template"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

var (
	templateDir    = os.Getenv("WORK_DIR") + "/template"
	commonDir      = templateDir + "/common/"
	templateHeader = "header.html"
	templateFooter = "footer.html"
	mainTemplate   = "base"
)

// ApplicationHTTPHandler base handler struct.
type ApplicationHTTPHandler struct {
	BaseHTTPHandler
}

// DataTemplate uses all mapping template data.
type DataTemplate struct {
	Common CommonTemplate
	Data   interface{}
}

// CommonTemplate uses common mapping template data.
type CommonTemplate struct {
	ErrorCode int
	ErrorInfo string
}

// ApplicationHTTPHandlerInterface is interface.
type ApplicationHTTPHandlerInterface interface {
	ResponseHTML(http.ResponseWriter, string, interface{})
	ResponseErrorHTML(http.ResponseWriter, int, interface{})
}

// ResponseHTML responses status code 200 and html.
func (h *ApplicationHTTPHandler) ResponseHTML(w http.ResponseWriter, r *http.Request, data interface{}, templateFiles ...string) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	commonTemplate := CommonTemplate{}

	d := &DataTemplate{
		Common: commonTemplate,
		Data:   data,
	}
	var tmpls []string
	for _, templateFile := range templateFiles {
		tmpls = append(tmpls, templateDir+"/"+templateFile+".tmpl")
	}

	return executeTemplate(w, d, tmpls...)
}

// ResponseErrorHTML calls utils.ResponseHTML.
func (h *ApplicationHTTPHandler) ResponseErrorHTML(w http.ResponseWriter, r *http.Request, code int, errorInfo string) error {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	commonTemplate := CommonTemplate{
		ErrorCode: code,
		ErrorInfo: errorInfo,
	}

	// set parameters.
	d := &DataTemplate{
		Common: commonTemplate,
	}
	// set common template.
	return executeTemplate(w, d, templateDir+"/error/error.html")
}

// executeTemplate executes template.
func executeTemplate(w http.ResponseWriter, data interface{}, templateFile ...string) error {
	tmpl := template.Must(template.ParseFiles(
		templateFile...,
	))
	err := tmpl.ExecuteTemplate(w, mainTemplate, data)
	if err != nil {
		panic(err)
	}
	return err
}

// makeSlice is template custom function.
func makeSlice(args ...interface{}) []interface{} {
	return args
}

// NewApplicationHTTPHandler returns ApplicationHTTPHandler instance.
func NewApplicationHTTPHandler(logger *logrus.Logger) *ApplicationHTTPHandler {
	return &ApplicationHTTPHandler{BaseHTTPHandler: BaseHTTPHandler{Logger: logger}}
}

// StatusServerError responses status code 500 and html.
func (h *ApplicationHTTPHandler) StatusServerError(w http.ResponseWriter, r *http.Request) error {
	// status code 500
	return h.ResponseErrorHTML(w, r, http.StatusInternalServerError, "500 Internal Server Error.")
}
