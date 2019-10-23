package login

import (
	"database/sql"
	"fmt"
	"net/http"
	"sample/app/shared/handler"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.ApplicationHTTPHandler
}

type templateData struct {
	LoginError string
	UserName   string
}

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

var loginError = ""

func (h *HTTPHandler) Login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	userName, _ := session.Values["userName"].(string)
	if userName == "" {
		err := h.ResponseHTML(w, r, "login/login", templateData{
			LoginError: loginError,
			UserName:   userName,
		})
		if err != nil {
			_ = h.StatusServerError(w, r)
		}
	} else {
		http.Redirect(w, r, "/profile", 302)
	}

}

var infor string

func CheckUser(user string, pass string) bool {
	var check bool
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

	var password string
	
	password = handler.GetPassword(user)
	fmt.Println(password)
	fmt.Println(pass)
	match := handler.CheckPasswordHash(pass, password)
	fmt.Print(match)
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
			if user == value && match == true {
				check = true
				break
			} else {
				check = false
			}
		}
		if check == true {
			break
		}
	}
	if check == true {
		return true
	} else {
		return false
	}

}

func (h *HTTPHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
	loginError = ""
	userName := r.FormValue("user-name")
	password := r.FormValue("password")
	if !(len(userName) <= 0) || !(len(password) <= 0) {
		userIsValid := CheckUser(userName,password)
		if userIsValid {
			session, _ := store.Get(r, "user")
			session.Values["userName"] = userName
			session.Values["password"] = password
			session.Save(r, w)
		} else {
			loginError = "Incorrect name and password"
		}
	} else {
		loginError = "Please, input name and password!"
	}

	if loginError == "" {
		http.Redirect(w, r, "/profile", 302)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

// NewHTTPLoginHandler responses new HTTPLoginHandler instance.
func NewLoginHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
