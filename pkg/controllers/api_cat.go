package controllers

import (
	"database/sql"
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
	categories, err := models.GetAllCategories()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching categories: %v", err), http.StatusInternalServerError)
		return
	}

	util.EncodeAndSendCategoriesWithStatus(w, categories, http.StatusOK)
}

func (cc *CatApiController) AddCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing the form: %v", err), http.StatusInternalServerError)
		return
	}

	var newCat models.Category

	newCat.CatName = r.Form.Get("cat_name")

	descStr := r.Form.Get("cat_description")
	if descStr == "" {
		newCat.CatDescription = sql.NullString{
			Valid: false,
		}
	}
	newCat.CatDescription = sql.NullString{
		String: descStr,
		Valid:  true,
	}

	validDemand, err := models.AddCategory(newCat)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding category: %v", err), http.StatusInternalServerError)
		return
	}
	var responseJson util.StandardResponseJson
	if !validDemand {
		responseJson.Msg = "Uable to add new Category"
		responseJson.ErrDescription = "Category already exists"
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusBadRequest)
		return
	}

	responseJson.Msg = fmt.Sprintf("New category added succesfully : %v", newCat.CatName)
	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusCreated)
}
