package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/cache"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

type OrderApiController struct{}

func NewOrderApiController() *OrderApiController {
	return &OrderApiController{}
}

func (oc *OrderApiController) PlaceOrder(w http.ResponseWriter, r *http.Request) {
	xUser := util.ExtractUserFromContext(r)
	tableNoStr := util.ExtractFromContext(r, "tableNo")

	orderSlice := util.ExtractCartFromContext(r)

	if len(orderSlice) == 0 {
		fmt.Println("the order is empty !!")
		popup := util.Popup{
			Msg:     "Order Not Placed ! Empty Cart",
			IsError: true,
		}
		util.InsertPopupInFlash(w, r, popup)
		var responseJson util.StandardResponseJson
		responseJson.Msg = "Order Not Placed !!! no item present"
		util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusCreated)
		return
	}

	totalOrderPrice := 0
	for _, v := range orderSlice {
		totalOrderPrice += int(v.Price * v.Quantity)
	}

	var newOrder models.Order
	newOrder.CustomerId = xUser.UserId
	newOrder.OrderAt = time.Now()
	newOrder.TableNo, _ = util.StringToSqlNullInt64(tableNoStr)
	newOrder.TotalPrice = int64(totalOrderPrice)
	newOrder.Status = "received"

	orderId, err := models.PlaceNewOrder(newOrder)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error placing new order: %v", err), http.StatusInternalServerError)
		return
	}
	newOrder.OrderId = orderId
	err = models.OccupyTable(newOrder.TableNo.Int64)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	err = models.EntriesInItemOrder(orderSlice, newOrder)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusInternalServerError)
		return
	}

	cacheField := fmt.Sprintf("orders%s", time.Now().Format("2006-01-02"))
	cache.AppCache.Delete(cacheField)

	popup := util.Popup{
		Msg:     fmt.Sprintf("Successfully placed Order #%v", newOrder.OrderId),
		IsError: false,
	}

	util.InsertPopupInFlash(w, r, popup)

	var responseJson util.StandardResponseJson
	responseJson.Msg = fmt.Sprintf("Order Placed !!! orderId = %v", newOrder.OrderId)
	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusCreated)
}

func (oc *OrderApiController) OrderPayment(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing form data , %v", err), http.StatusInternalServerError)
		return
	}
	orderId := r.Form.Get("order_id")
	customerReview := r.Form.Get("customer_review")

	err = models.MakePayment(orderId, customerReview)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing the payment , %v", err), http.StatusInternalServerError)
		return
	}

	err = models.VacateTable()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error vacating table , %v", err), http.StatusInternalServerError)
		return
	}

	cache.AppCache.Delete("orders")

	popup := util.Popup{
		Msg: "Payment successful... happyCustomer++",
	}
	util.InsertPopupInFlash(w, r, popup)
	util.RedirectToSite(w, r, "/static/order")
}
