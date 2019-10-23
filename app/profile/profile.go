package profile

import (
	"database/sql"
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
	Email            string
	Password         string
	Name             string
	Image            string
	Phone            string
	Age              string
	Messenger        string
	Location         string
	MessengerSuccess string
}

var (
	key   = []byte("secret-key")
	store = sessions.NewCookieStore(key)
)

var messengersuccess = ""
var messenger = ""
var email = ""
var password = ""
var name = ""
var location = ""
var image = "../../static/image/avatar_default"
var phone = ""
var age = ""

func (h *HTTPHandler) ShowProfile(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "user")
	userName, _ := session.Values["userName"].(string)
	db, err := sql.Open("mysql", "root:@tcp(192.168.200.181:3306)/golang")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	defer db.Close()
	email = userName
	password = handler.GetPassword(userName)
	name = handler.GetName(userName)
	image = handler.GetImage(userName)
	phone = handler.GetPhone(userName)
	age = handler.GetAge(userName)
	location = handler.GetLocation(userName)

	err = h.ResponseHTML(w, r, "profile/profile", templateData{
		Email:            email,
		Password:         password,
		Name:             name,
		Image:            image,
		Phone:            phone,
		Age:              age,
		Messenger:        messenger,
		Location:         location,
		MessengerSuccess: messengersuccess,
	})
	if err != nil {
		_ = h.StatusServerError(w, r)
	}

}

//HandleProfile function save data of user to database
func (h *HTTPHandler) HandleProfile(w http.ResponseWriter, r *http.Request) {
	messengersuccess = ""
	messenger = ""
	session, _ := store.Get(r, "user")
	userName, _ := session.Values["userName"].(string)
	email = userName
	newemail := r.FormValue("email")
	pas := r.FormValue("password")
	var pass string
	pass, _ = handler.HashPassword(pas)
	name := r.FormValue("name")
	age := r.FormValue("age")
	location := r.FormValue("location")
	phone := r.FormValue("phone")
	var oldimage string
	oldimage = handler.GetImage(userName)
	 
	image = handler.UploadFile(w, r) 
	if image != "nofile" {
		path := oldimage[6:len(oldimage)]+".png"
		handler.DeleteImage(path)
	} else {
		image = oldimage
	}
	
	if handler.CheckEmail(newemail) == true {
		messenger = "Email has existed!"
	}

	if newemail == email {
		messenger = ""
	}

	if messenger != "" {
		http.Redirect(w, r, "/profile", 302)
	} else {

		db, err := sql.Open("mysql", "root:@tcp(192.168.200.181:3306)/golang")
		if err != nil {
			panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		}
		defer db.Close()

		// if err = rows.Err(); err != nil {
		// 	panic(err.Error()) // proper error handling instead of panic in your app
		// }

		stmtIns, err := db.Prepare("UPDATE user SET email = ?, password = ?, name = ?, age = ?, location = ?, phone = ?, image = ? WHERE email = ?") // ? = placeholder
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

		// Insert data to Database
		_, err = stmtIns.Exec(newemail, pass, name, age, location, phone, image, email) // Insert information of user
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		messengersuccess = "Update Successfully!"
		http.Redirect(w, r, "/profile", 302)

	}
}

// NewHTTPLoginHandler responses new HTTPLoginHandler instance.
func NewProfileHTTPHandler(ah *handler.ApplicationHTTPHandler) *HTTPHandler {
	// item set.
	return &HTTPHandler{ApplicationHTTPHandler: *ah}
}
