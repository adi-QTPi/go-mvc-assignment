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
