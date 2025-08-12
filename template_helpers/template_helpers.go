package template_helpers

import (
	"encoding/json"
	"html/template"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/models"
	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

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
