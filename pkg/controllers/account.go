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

	var responseJson util.StandardResponseJson

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
		responseJson.Msg = "Unable to add user"
		responseJson.ErrDescription = "user with this username exists"

		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusConflict)
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
	responseJson.Msg = "successfully added user"

	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusCreated)
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

	var responseJson util.StandardResponseJson
	responseJson.Msg = "Logged in Successfully"
	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusOK)
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

	var responseJson util.StandardResponseJson
	responseJson.Msg = "Successfully logged Out."
	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusOK)
}

func (ac *AccountController) ShowProfile(w http.ResponseWriter, r *http.Request) {

	xUser := util.ExtractUserFromContext(r)

	util.EncodeAndSendUsersWithStatus(w, http.StatusOK, xUser)
}
