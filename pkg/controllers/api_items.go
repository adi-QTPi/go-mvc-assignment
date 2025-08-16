package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

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
		popup := util.Popup{
			Msg:     "Invalid input for price",
			IsError: true,
		}
		util.InsertPopupInFlash(w, r, popup)
		reqFrom := r.Referer()
		util.RedirectToSite(w, r, reqFrom)
		return
	}

	newItem.DisplayPic.String, err = util.ManageImageUpload(r)
	if err != nil {
		newItem.DisplayPic.Valid = false
		http.Error(w, "error in image upload ", http.StatusBadRequest)
		return
	}
	newItem.DisplayPic.Valid = true

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
	util.AppCache.Delete("menu")

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
}

func (ic *ItemApiController) GetItems(w http.ResponseWriter, r *http.Request) {

	if items, ok := util.AppCache.Get("menu"); ok {
		util.EncodeAndSendItemWithStatus(w, items.([]models.DisplayItem), http.StatusOK)
	}

	items, err := models.GetAllItems()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching items: %v", err), http.StatusInternalServerError)
		return
	}
	util.AppCache.Set("menu", items, 24*time.Hour)
	util.EncodeAndSendItemWithStatus(w, items, http.StatusOK)
}

func (ic *ItemApiController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	queryParams := mux.Vars(r)
	ItemIdStr := queryParams["item_id"]

	err := models.DeleteItemById(ItemIdStr)
	if err != nil {
		http.Error(w, "Unable to Delete item", http.StatusBadRequest)
		return
	}

	util.AppCache.Delete("menu")

	popup := util.Popup{
		Msg:     "Item Deleted Successfully",
		IsError: false,
	}
	err = util.InsertPopupInFlash(w, r, popup)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to insert popup in flash : %v", err), http.StatusBadRequest)
		return
	}

	requestFrom := r.Referer()
	util.RedirectToSite(w, r, requestFrom)
}
