package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type CookApiController struct{}

func NewCookApiController() *CookApiController {
	return &CookApiController{}
}

func (coc *CookApiController) ChangeKitchenOrderStatus(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	orderIdStr := r.Form.Get("order_id")
	itemIdStr := r.Form.Get("item_id")
	isComplete := r.Form.Get("is_complete")

	orderId, err := strconv.ParseInt(orderIdStr, 10, 64)
	itemId, err := strconv.ParseInt(itemIdStr, 10, 64)
	if err != nil {
		http.Error(w, fmt.Sprintf("Invalid input for item_id or order_id : %v", err), http.StatusBadRequest)
		return
	}

	xUser := util.ExtractUserFromContext(r)
	cookId := xUser.UserId

	err = models.StatusUpdateByCook(cookId, orderId, itemId, isComplete)
	if err != nil {
		http.Error(w, fmt.Sprintf("error in model function : %v", err), http.StatusInternalServerError)
		return
	}

	err = models.SyncOrderStatus()
	if err != nil {
		http.Error(w, fmt.Sprintf("error in syncing order status (model function) : %v", err), http.StatusInternalServerError)
		return
	}

	util.RedirectToSite(w, r, "/static/cook")
}
