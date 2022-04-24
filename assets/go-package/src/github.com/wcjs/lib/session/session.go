package session

import (
	"io"
	"crypto/rand"
	"fmt"
	"net/http"
)


type SessionManager struct {

}

var sessionmanger = &SessionManager{}
func NewSession() *SessionManager{
	return sessionmanger
}

func GenerateSID() string{
	b := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return fmt.Sprintf("%x",b)
}

func (sm *SessionManager) GetSID(w http.ResponseWriter,r *http.Request) (string,bool) {
	cookie,err := r.Cookie("PHPSESSID")
	if err == nil {
		return cookie.Value,true
	} else {
		sid := GenerateSID()
		newCookie := &http.Cookie{
			Name:     "PHPSESSID",
			Value:    sid,
			Path:     "/",
			HttpOnly: false,
			MaxAge:   int(864000),
		}
		http.SetCookie(w,newCookie)
		return sid,false
	}
}