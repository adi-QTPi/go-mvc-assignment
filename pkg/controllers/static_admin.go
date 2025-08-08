package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/gorilla/mux"
)

type AdminStaticController struct{}

func NewAdminStaticController() *AdminStaticController {
	return &AdminStaticController{}
}

func (asc *AdminStaticController) FetchAdminOrderDashboardByDate(w http.ResponseWriter, r *http.Request) {
	xUser := util.ExtractUserFromContext(r)

	params := mux.Vars(r)
	reqDate := params["date"]

	if reqDate == "" {
		today := time.Now()
		reqDate = today.Format("2006-01-02") //dont change date
	}

	var orderSlice []models.Order
	orderSlice, err := models.FetchAllOrderDetailsByDate(reqDate, xUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching orders: %v", err), http.StatusInternalServerError)
		return
	}

	util.EncodeAndSendOrderWithStatus(w, orderSlice, http.StatusOK)
}
