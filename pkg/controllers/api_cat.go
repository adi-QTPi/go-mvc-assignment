package controllers

import (
	"fmt"
	"net/http"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type CatApiController struct{}

func NewCatApiController() *CatApiController {
	return &CatApiController{}
}

func (cc *CatApiController) GetCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all categories controller called")

	categories, err := models.GetAllCategories()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching categories: %v", err), http.StatusInternalServerError)
		return
	}

	util.EncodeAndSendCategoriesWithStatus(w, categories, http.StatusOK)
}
