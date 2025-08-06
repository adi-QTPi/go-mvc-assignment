package controllers

import (
	"fmt"
	"net/http"
)

type UserApiController struct{}

func NewUserApiController() *UserApiController {
	return &UserApiController{}
}

func (uc *UserApiController) GetAllUsers(w http.ResponseWriter, r *http.Request) {

	_, err := fmt.Fprintf(w, "data for all users...")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching users: %v", err), http.StatusInternalServerError)
		return
	}

}

func (uc *UserApiController) GetUserById(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)

	_, err := fmt.Fprintf(w, "asked for the user with id :")
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching this user: %v", err), http.StatusInternalServerError)
		return
	}
}
