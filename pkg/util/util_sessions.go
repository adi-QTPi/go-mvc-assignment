package util

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"

	"github.com/gorilla/sessions"
)

var SessionStore *sessions.CookieStore
var sessionName = "foodopia-session"

type M map[string]interface{}

func InitiateStructSession() {
	SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSIONS_SECRET")))
	gob.Register(&Popup{})
	gob.Register(&M{})
	gob.Register(&models.User{})
}

func InsertPopupInFlash(w http.ResponseWriter, r *http.Request, object Popup) error {
	session, err := SessionStore.Get(r, sessionName)
	if err != nil {
		return fmt.Errorf("error getting the session from session store : %v", err)
	}

	session.AddFlash(object)

	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("error saving the flash session : %v", err)
	}

	return nil
}

func ExtractPopupFromFlash(w http.ResponseWriter, r *http.Request) (Popup, error) {
	var nilPopup = Popup{}
	var popup = &Popup{}

	session, err := SessionStore.Get(r, sessionName)
	if err != nil {
		return nilPopup, fmt.Errorf("error getting data : %v", err)
	}

	if flashes := session.Flashes(); len(flashes) > 0 {
		var ok bool
		popup, ok = flashes[0].(*Popup)
		if !ok {
			return nilPopup, fmt.Errorf("error in deserialisation flash : %v", err)
		}
	}

	err = session.Save(r, w)
	if err != nil {
		return nilPopup, fmt.Errorf("error saving the session : %v", err)
	}

	return *popup, nil
}

func InsertUserInSession(w http.ResponseWriter, r *http.Request, user models.User) error {
	session, err := SessionStore.Get(r, sessionName)
	if err != nil {
		return fmt.Errorf("error getting session: %v", err)
	}

	session.Values["XUser"] = user

	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}
