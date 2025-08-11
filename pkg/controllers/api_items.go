package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/gorilla/mux"
)

type ItemApiController struct{}

func NewItemApiController() *ItemApiController {
	return &ItemApiController{}
}

func (ic *ItemApiController) AddItem(w http.ResponseWriter, r *http.Request) {
	const maxMemory = 32 << 20

	var err error
	if r.Referer() == "/static/menu" {
		err = r.ParseMultipartForm(maxMemory)
	} else {
		err = r.ParseForm()
	}
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	var newItem models.Item

	newItem.ItemName = r.Form.Get("item_name")

	CookTimeMinStr := r.Form.Get("cook_time_min")
	if CookTimeMinStr != "" {
		newItem.CookTimeMin, err = util.StringToSqlNullInt64(CookTimeMinStr)
		if err != nil {
			http.Error(w, "Invalid input for cook time ", http.StatusBadRequest)
			return
		}
	}

	PriceStr := r.Form.Get("price")
	newItem.Price, err = strconv.ParseInt(PriceStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid input for price ", http.StatusBadRequest)
		return
	}

	// newItem.DisplayPic = r.Form.Get("display_pic")

	CatIdStr := r.Form.Get("cat_id")
	newItem.CatId, err = strconv.ParseInt(CatIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid input for cat_id ", http.StatusBadRequest)
		return
	}

	SubCatIdStr := r.Form.Get("subcat_id")
	if SubCatIdStr != "" {
		newItem.SubCatId, err = util.StringToSqlNullInt64(SubCatIdStr)
		if err != nil {
			http.Error(w, "Invalid input for subcat id", http.StatusBadRequest)
			return
		}
	}

	err = models.AddItem(newItem)
	if err != nil {
		fmt.Println("error in adding item : ", err)
		http.Error(w, "Unable to add new item ", http.StatusInternalServerError)
		return
	}

	popup := util.Popup{
		Msg:     "Item Added Successfully",
		IsError: false,
	}

	err = util.InsertPopupInFlash(w, r, popup)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error adding popup: %v", err), http.StatusInternalServerError)
	}

	requestFrom := r.Referer()

	util.RedirectToSite(w, r, requestFrom)

	// var responseJson util.StandardResponseJson
	// responseJson.Msg = "New item created successfully"

	// util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusOK)
}

func (ic *ItemApiController) GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := models.GetAllItems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching items: %v", err), http.StatusInternalServerError)
		return
	}

	// err = util.InsertItemsInSession(w, r, items)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Error puttin menu in sessions: %v", err), http.StatusInternalServerError)
	// }

	// util.RedirectToSite(w, r, "/static/menu")
	util.EncodeAndSendItemWithStatus(w, items, http.StatusOK)
}

func (ic *ItemApiController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	ItemIdStr := queryParams["item_id"]

	err := models.DeleteItemById(ItemIdStr)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unable to Delete item", http.StatusBadRequest)
		return
	}

	var responseJson util.StandardResponseJson
	responseJson.Msg = "Item deleted successfully"
	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusOK)
}
