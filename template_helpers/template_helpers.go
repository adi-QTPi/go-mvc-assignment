package template_helpers

import (
	"encoding/json"
	"html/template"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

var FuncMap = template.FuncMap{
	"ToJSON":                    ToJSON,
	"add":                       Add,
	"multiply":                  Multiply,
	"booleanUpdater":            CookPageHelper,
	"itemsTakenByCook":          ItemsTakenByCook,
	"itemsPending":              ItemsPending,
	"itemsCompletedByCookToday": ItemsCompletedByCookToday,
}

func ToJSON(v interface{}) (template.JS, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return template.JS(b), nil
}

func Multiply(a, b int) int {
	return a * b
}

func Add(a, b int) int {
	return a + b
}

func CookPageHelper(items []models.KitchenOrder) util.CookPageHelperStruct {
	var indicator util.CookPageHelperStruct
	for _, item := range items {
		if item.IsComplete == "pending" {
			indicator.SomePending = true
		}
		if item.IsComplete == "taken" {
			indicator.SomeTaken = true
		}
		if item.IsComplete == "complete" {
			indicator.SomeComplete = true
		}
	}
	return indicator
}

func ItemsTakenByCook(orders []models.KitchenOrder, cookID string) []models.KitchenOrder {
	var filteredOrders []models.KitchenOrder
	for _, order := range orders {
		if order.CookId.String == cookID && order.IsComplete == "taken" {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func ItemsPending(orders []models.KitchenOrder) []models.KitchenOrder {
	var filteredOrders []models.KitchenOrder
	for _, order := range orders {
		if order.IsComplete == "pending" {
			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}

func ItemsCompletedByCookToday(orders []models.KitchenOrder, cookID string) []models.KitchenOrder {
	var filteredOrders []models.KitchenOrder
	today := time.Now()

	for _, order := range orders {
		if order.CookId.String == cookID &&
			order.IsComplete == "complete" &&
			order.OrderAt.Year() == today.Year() &&
			order.OrderAt.Month() == today.Month() &&
			order.OrderAt.Day() == today.Day() {

			filteredOrders = append(filteredOrders, order)
		}
	}
	return filteredOrders
}
