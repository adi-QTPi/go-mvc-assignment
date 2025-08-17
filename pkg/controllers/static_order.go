package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
	"github.com/adi-QTPi/go-mvc-assignment/template_helpers"
	"github.com/gorilla/mux"
)

type StaticOrderCotroller struct{}

func NewStaticOrderController() *StaticOrderCotroller {
	return &StaticOrderCotroller{}
}

func (cuc *StaticOrderCotroller) RenderCustOrderPage(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	reqDate := queryParams.Get("date")

	if reqDate == "" {
		today := time.Now()
		reqDate = today.Format("2006-01-02") //dont change date pls
	}

	xUser := util.ExtractUserFromContext(r)

	var orderSlice []models.Order
	orderSlice, err := models.FetchAllOrderDetailsByDate(reqDate, xUser)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching orders: %v", err), http.StatusInternalServerError)
		return
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

	err = template_helpers.Tmpl.ExecuteTemplate(w, "order.html", toPage)
	if err != nil {
		fmt.Printf("error rendering order page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func (cuc *StaticOrderCotroller) RenderOrderById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	orderId := params["id"]

	xUser := util.ExtractUserFromContext(r)

	orderMetaData, err := models.FetchOrderByOrderId(orderId)
	billData, err := models.FetchBillDetailsByOrderId(orderId)

	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching bill data : %v", err), http.StatusInternalServerError)
		return
	}

	popup, err := util.ExtractPopupFromFlash(w, r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error getting stuff from sessions and db : %v", err), http.StatusInternalServerError)
		return
	}

	toPage := util.DataToPage{
		Popup:         popup,
		XUser:         xUser,
		OrderMetaData: orderMetaData,
		BillData:      billData,
	}

	err = template_helpers.Tmpl.ExecuteTemplate(w, "bill_by_order_id.html", toPage)
	if err != nil {
		fmt.Printf("error rendering bill page : %v\n", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}
