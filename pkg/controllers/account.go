package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type AccountController struct{}

func NewAccountController() *AccountController {
	return &AccountController{}
}

func (ac *AccountController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	var newUser models.User
	newUser.UserName = r.Form.Get("user_name")
	newUser.Name = r.Form.Get("name")
	pwdHash := r.Form.Get("pwd")
	newUser.Role = "customer"

	alreadyExists, err, _ := models.GetUserByUsername(newUser.UserName)
	if alreadyExists {
		var popup = util.Popup{
			Msg:     "This Username is Already Taken ... Be more creative!",
			IsError: true,
		}

		requestFrom := r.Referer()
		util.InsertPopupInFlash(w, r, popup)
		util.RedirectToSite(w, r, requestFrom)
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking for existing user: %v", err), http.StatusInternalServerError)
		return
	}

	err = models.AddNewUser(newUser, pwdHash)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Adding user: %v", err), http.StatusInternalServerError)
		return
	}

	var popup = util.Popup{
		Msg:     "Account Created Successfully",
		IsError: false,
	}

	err = util.InsertPopupInFlash(w, r, popup)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error inserting popup before redirect (in signup page): %v", err), http.StatusInternalServerError)
	}

	util.RedirectToSite(w, r, "/login")

}

func (ac *AccountController) CreateNewUserByAdmin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	var newUser models.User
	newUser.UserName = r.Form.Get("user_name")
	newUser.Name = r.Form.Get("name")
	pwdHash := r.Form.Get("pwd")
	newUser.Role = r.Form.Get("role")

	alreadyExists, err, _ := models.GetUserByUsername(newUser.UserName)
	if alreadyExists {
		var popup = util.Popup{
			Msg:     "Username taken, try a different username for yourself",
			IsError: true,
		}
		err := util.InsertPopupInFlash(w, r, popup)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error in inserting popup: %v", err), http.StatusInternalServerError)
		}
		util.RedirectToSite(w, r, "/signup")
		return
	}
	if err != nil {
		http.Error(w, fmt.Sprintf("Error checking for existing user: %v", err), http.StatusInternalServerError)
		return
	}

	err = models.AddNewUser(newUser, pwdHash)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Adding user: %v", err), http.StatusInternalServerError)
		return
	}

	var popup = util.Popup{
		Msg:     "Account Created Successfully",
		IsError: false,
	}

	util.InsertPopupInFlash(w, r, popup)
	util.RedirectToSite(w, r, "/signup")
}

func (ac *AccountController) LogUserIn(w http.ResponseWriter, r *http.Request) {
	userId := util.ExtractFromContext(r, "user_id")

	signedJwtTokenString, err := util.GetSignedJwtOfUser(w, userId)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error in making signed JWT: %v", err), http.StatusInternalServerError)
	}

	jwtCookie := http.Cookie{
		Name:     "jwt_token",
		Value:    signedJwtTokenString,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
	}

	http.SetCookie(w, &jwtCookie)

	util.RedirectToSite(w, r, "/static")
}

func (ac *AccountController) LogUserOut(w http.ResponseWriter, r *http.Request) {
	jwtCookie := http.Cookie{
		Name:     "jwt_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		Expires:  time.Now().Add(24 * time.Hour),
		MaxAge:   -1,
	}

	http.SetCookie(w, &jwtCookie)

	util.RedirectToSite(w, r, "/")
}

func (ac *AccountController) ShowProfile(w http.ResponseWriter, r *http.Request) {

	xUser := util.ExtractUserFromContext(r)

	util.EncodeAndSendUsersWithStatus(w, http.StatusOK, xUser)
}
