package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
)

type AdminStaticController struct{}

func NewAdminStaticController() *AdminStaticController {
	return &AdminStaticController{}
}

func (asc *AdminStaticController) FetchAdminOrderDashboardByDate(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()
	reqDate := queryParams.Get("date")

	if reqDate == "" {
		today := time.Now()
		reqDate = today.Format("2006-01-02") //dont change date pls
	}

	cacheField := fmt.Sprintf("orders%s", reqDate)

	xUser := util.ExtractUserFromContext(r)

	var orderSlice []models.Order

	orders, ok := util.AppCache.Get(cacheField)
	if !ok {
		var err error
		orderSlice, err = models.FetchAllOrderDetailsByDate(reqDate, xUser)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error fetching orders: %v", err), http.StatusInternalServerError)
			return
		}
		cacheField = fmt.Sprintf("orders%s", reqDate)
		util.AppCache.Set(cacheField, orderSlice, 24*time.Hour)
	}
	if ok {
		orderSlice = orders.([]models.Order)
	}

	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting stuff from sessions and db : %v", err), http.StatusInternalServerError)
		return
	}

	toPage := util.DataToPage{
		Popup:     popup,
		XUser:     xUser,
		OrderData: orderSlice,
		ReqDate:   reqDate,
	}

	var responseJson util.StandardResponseJson
	err = template_helpers.Tmpl.ExecuteTemplate(w, "admin.html", toPage)
	if err != nil {
		responseJson.Msg = "Can't show this page"
		responseJson.ErrDescription = fmt.Sprintf("Error in executing admin.html : %v", err)
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusInternalServerError)
	}
}
