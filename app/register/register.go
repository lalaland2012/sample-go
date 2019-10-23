package register

import (
	"net/http"
	"sample/app/shared/handler"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	RegisterError string
	Noti          string
}

var registerError = ""

var noti = ""

func (h *HTTPHandler) Register(w http.ResponseWriter, r *http.Request) {
	err := h.ResponseHTML(w, r, "register/register", templateData{
		RegisterError: registerError,
		Noti:          noti,
	})
	if err != nil {
		_ = h.StatusServerError(w, r)
	}

}

func (h *HTTPHandler) HandleRegister(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	pass := r.FormValue("pass")
	name := r.FormValue("name")
	var pas string
	pas, _ = handler.HashPassword(pass)

	db, err := sql.Open("mysql", "root:@tcp(192.168.200.181:3306)/golang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()

	// Execute the query
	rows, err := db.Query("SELECT email FROM user")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))

	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for _, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			if email == value {
				registerError = "Email has existed!"
			} else {
				registerError = ""
			}
		}
	}

	if registerError != "" {
		http.Redirect(w, r, "/register", 302)
	} else {

		if err = rows.Err(); err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Prepare statement for inserting data
		stmtIns, err := db.Prepare("INSERT INTO user (email,password,name) VALUES( ?, ?, ? )") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

		// Insert data to Database
		_, err = stmtIns.Exec(email, pas, name) // Insert information of user
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		noti = "Register Successfully! You can login now."
		http.Redirect(w, r, "/register", 302)

	}
}

// NewHTTPLoginHandler responses new HTTPLoginHandler instance.
func NewRegisterHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
