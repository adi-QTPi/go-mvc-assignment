package controllers

import (
	"fmt"
	"net/http"
	"time"

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

	totalOrderPrice := 0
	for _, v := range orderSlice {
		totalOrderPrice += int(v.TotalPrice)
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

	var responseJson util.StandardResponseJson
	responseJson.Msg = fmt.Sprintf("Ordder Placed !!! orderId = %v", newOrder.OrderId)
	util.EncodeAndSendResponseWithStatus(w, responseJson, http.StatusCreated)
}
