package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/gorilla/mux"
)

type UserApiController struct{}

func NewUserApiController() *UserApiController {
	return &UserApiController{}
}

func (uc *UserApiController) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching users: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (uc *UserApiController) GetUser(w http.ResponseWriter, r *http.Request) {

	queryParams := mux.Vars(r)
	id := queryParams["id"]

	user, err := models.GetUserById(id)
	if err != nil {
		if err.Error() == "user not found" {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		http.Error(w, fmt.Sprintf("Error fetching user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func (uc *UserApiController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	userID := queryParams["id"]

	err := models.DeleteUserById(userID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Deleting user: %v", err), http.StatusInternalServerError)
		return
	}

	var responseJson util.StandardResponseJson
	responseJson.Msg = fmt.Sprintf("Success : Deleted user %v", userID)

	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusOK)
}

func (uc *UserApiController) AddUser(w http.ResponseWriter, r *http.Request) {

	var responseJson util.StandardResponseJson

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	var newUser models.User
	newUser.UserName = r.Form.Get("user_name")
	newUser.Name = r.Form.Get("name")
	newUser.PwdHash = r.Form.Get("pwd_hash")
	newUser.Role = r.Form.Get("role")

	alreadyExists, err := models.GetUserByUsername(newUser.UserName)
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

	err = models.AddNewUser(newUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error Adding user: %v", err), http.StatusInternalServerError)
		return
	}
	responseJson.Msg = "successfully added user"

	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusCreated)
}
