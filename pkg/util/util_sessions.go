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

func registerGobTypes(values ...any) {
	for _, v := range values {
		gob.Register(v)
	}
}

func InitiateStructSession() {
	SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSIONS_SECRET")))

	registerGobTypes(
		&Popup{},
		&M{},
		&models.User{},

		models.DisplayItem{},
		// &models.DisplayItem{}, giving error (duplication ???why)
		[]models.DisplayItem{},
		[]*models.DisplayItem{},

		models.Category{},
		// &models.Category{}, giving error (duplication ???why)
		[]models.Category{},
		[]*models.Category{},

		models.Order{},
		[]models.Order{},
		[]*models.Order{},

		models.ItemOrderDescriptive{},
		[]models.ItemOrderDescriptive{},
		[]*models.ItemOrderDescriptive{},

		models.KitchenOrder{},
		[]models.KitchenOrder{},
	)
}

func InsertPopupInFlash(w http.ResponseWriter, r *http.Request, object Popup) error {
	session, _ := SessionStore.Get(r, sessionName)

	session.AddFlash(object)

	err := session.Save(r, w)
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

func InsertItemsInSession(w http.ResponseWriter, r *http.Request, itemSlice []models.DisplayItem) error {
	session, err := SessionStore.Get(r, sessionName)
	if err != nil {
		return fmt.Errorf("error getting session: %v", err)
	}

	session.Values["Menu"] = itemSlice

	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}

func ExtractItemsFromSession(r *http.Request) ([]models.DisplayItem, error) {
	session, err := SessionStore.Get(r, sessionName)
	if err != nil {
		return nil, fmt.Errorf("error getting session: %v", err)
	}

	items, ok := session.Values["Menu"].([]models.DisplayItem)
	if !ok {
		return nil, fmt.Errorf("no valid item slice found in session")
	}

	return items, nil
}

func InsertCategoriesInSession(w http.ResponseWriter, r *http.Request, itemSlice []models.Category) error {
	session, err := SessionStore.Get(r, sessionName)
	if err != nil {
		return fmt.Errorf("error getting session: %v", err)
	}

	session.Values["CategoryList"] = itemSlice

	err = session.Save(r, w)
	if err != nil {
		return fmt.Errorf("error saving session: %v", err)
	}

	return nil
}

func ExtractCategoriesFromSession(r *http.Request) ([]models.Category, error) {
	session, err := SessionStore.Get(r, sessionName)
	if err != nil {
		return nil, fmt.Errorf("error getting session: %v", err)
	}

	items, ok := session.Values["CategoryList"].([]models.Category)
	if !ok {
		return nil, fmt.Errorf("no valid item slice found in session")
	}

	return items, nil
}
