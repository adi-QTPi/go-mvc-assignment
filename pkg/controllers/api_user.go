package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/gorilla/mux"
)

type UserApiController struct{}

func NewUserApiController() *UserApiController {
	return &UserApiController{}
}

func (uc *UserApiController) GetUsers(w http.ResponseWriter, r *http.Request) {

	users, err2 := models.GetAllUsers()
	if err2 != nil {
		http.Error(w, fmt.Sprintf("Error fetching users: %v", err2), http.StatusInternalServerError)
		return
	}

	for _, v := range users {
		fmt.Println(v)
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
