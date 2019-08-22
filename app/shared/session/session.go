package session

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

// GetStore returns a session for the given name after adding it to the registry.
func GetStore(w http.ResponseWriter, r *http.Request, name string) *sessions.Session {
	session, err := store.Get(r, name)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return nil
	}
	return session
}

//GetMsgFlash return msg in session and delete msg this
func GetMsgFlash(w http.ResponseWriter, r *http.Request, name string) string {
	session := GetStore(w, r, "message")
	msg := ""
	if flashes := session.Flashes(name); len(flashes) > 0 {
		session.Save(r, w)
		msg = flashes[0].(string)
	}
	return msg
}

//RemoveSession remove session in store by name
func RemoveSession(w http.ResponseWriter, r *http.Request, name string) {
	session := GetStore(w, r, name)
	session.Options.MaxAge = -1
	session.Save(r, w)
}
