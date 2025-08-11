package util

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)

var SessionStore *sessions.CookieStore

type M map[string]interface{}

func InitiateStructSession() {
	// SessionStore = sessions.NewCookieStore([]byte("dsfhsjhfsdfj;ksd;fkjj;lsfuiewghiufhiuwefhuhiuhisdfn"))
	SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSIONS_SECRET")))
	gob.Register(&Popup{})
	gob.Register(&M{})
}

func InsertPopupInFlash(w http.ResponseWriter, r *http.Request, object any) error {
	session, err := SessionStore.Get(r, "toNextPage")
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
	var popup = &Popup{}

	session, err := SessionStore.Get(r, "toNextPage")
	if err != nil {
		return *popup, fmt.Errorf("error getting data : %v", err)
	}

	if flashes := session.Flashes(); len(flashes) > 0 {
		var ok bool
		popup, ok = flashes[0].(*Popup)
		if !ok {
			return *popup, fmt.Errorf("error in deserialisation flash : %v", err)
		}
	}

	err = session.Save(r, w)
	if err != nil {
		return *popup, fmt.Errorf("error saving the session : %v", err)
	}

	return *popup, nil
}
