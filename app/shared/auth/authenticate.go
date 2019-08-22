package auth

import (
	"net/http"
	"sample/app/shared/helper"
	"sample/app/shared/session"
)

// Login Handler login
func Login(userName, pass string, w http.ResponseWriter, r *http.Request) bool {
	loginSuccess := false
	msgError := ""
	if !helper.IsEmpty(userName) && !helper.IsEmpty(pass) {
		userIsValid := Attempt(userName, pass)
		if userIsValid {
			session := session.GetStore(w, r, "user")
			session.Values["userName"] = userName
			session.Values["pass"] = pass
			session.Save(r, w)
			loginSuccess = true
		} else {
			msgError = "Incorrect name and password"
		}
	} else {
		msgError = "Please, input name and password!"
	}
	if !helper.IsEmpty(msgError) {
		sesValidate := session.GetStore(w, r, "message")
		sesValidate.AddFlash(msgError, "Error")
		sesValidate.Save(r, w)
	}
	return loginSuccess
}

// Attempt check user is valid
func Attempt(userName, pass string) bool {
	uName, pwd, isValid := "manhnd", "123456", false
	if uName == userName && pwd == pass {
		isValid = true
	} else {
		isValid = false
	}
	return isValid
}

// Logout handler logout
func Logout(w http.ResponseWriter, r *http.Request) {
	session.RemoveSession(w, r, "user")
}

// User get info user login
func User(w http.ResponseWriter, r *http.Request) map[string]string {
	session := session.GetStore(w, r, "user")
	user := make(map[string]string)
	userName, exsitUserName := session.Values["userName"].(string)
	pass, existPass := session.Values["pass"].(string)
	if exsitUserName {
		user["userName"] = userName
	}
	if existPass {
		user["pass"] = pass
	}

	return user
}
