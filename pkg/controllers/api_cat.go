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

	// err = util.InsertCategoriesInSession(w, r, categories)

	// util.RedirectToSite(w, r, "/api/item")
	util.EncodeAndSendCategoriesWithStatus(w, categories, http.StatusOK)
}

func (cc *CatApiController) AddCategory(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing the form: %v", err), http.StatusInternalServerError)
		return
	}

	var newCat models.Category

	newCat.CategoryName = r.Form.Get("category_name")

	descStr := r.Form.Get("category_description")
	if descStr == "" {
		newCat.CategoryDescription = sql.NullString{
			Valid: false,
		}
	}
	newCat.CategoryDescription = sql.NullString{
		String: descStr,
		Valid:  true,
	}

	validDemand, err := models.AddCategory(newCat)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding category: %v", err), http.StatusInternalServerError)
		return
	}

	if !validDemand {
		// responseJson.Msg = "Uable to add new Category"
		// responseJson.ErrDescription = "Category already exists"
		// util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusBadRequest)
		popup := util.Popup{
			Msg:     "This Category already Exists",
			IsError: true,
		}

		err := util.InsertPopupInFlash(w, r, popup)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error adding popup: %v", err), http.StatusInternalServerError)
		}

		requestFrom := r.Referer()

		util.RedirectToSite(w, r, requestFrom)

		return
	}

	popup := util.Popup{
		Msg:     "Category Added Successfully",
		IsError: false,
	}

	err = util.InsertPopupInFlash(w, r, popup)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding popup: %v", err), http.StatusInternalServerError)
	}

	requestFrom := r.Referer()

	util.RedirectToSite(w, r, requestFrom)
	// responseJson.Msg = fmt.Sprintf("New category added succesfully : %v", newCat.CategoryName)
	// util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusCreated)
}
