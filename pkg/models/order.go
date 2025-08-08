package models

import (
	"database/sql"
	"fmt"
	"time"
)

type ItemInCart struct {
	ItemId      int64  `json:"item_id"`
	Quantity    int64  `json:"quantity"`
	TotalPrice  int64  `json:"total_price"`
	Instruction string `json:"instruction"`
}

type Order struct {
	OrderId    int64         `json:"order_id"`
	OrderAt    time.Time     `json:"order_at"`
	TableNo    sql.NullInt64 `json:"table_no"`
	CustomerId string        `json:"customer_id"`
	Status     string        `json:"status"`
	TotalPrice int64         `json:"total_price"`
}

type ItemOrder struct {
	OrderId     int64          `json:"order_id"`
	ItemId      int64          `json:"item_id"`
	Quantity    int            `json:"quantity"`
	Instruction sql.NullString `json:"instruction"`
	IsComplete  string         `json:"is_complete"`
	CookId      sql.NullString `json:"cook_id"`
}

func CheckAndAssignTable(xUserId string) (int64, error) {
	sqlQuery := "SELECT DISTINCT table_no FROM `order` WHERE customer_id = ? AND status != 'paid';"

	row := DB.QueryRow(sqlQuery, xUserId)

	var tableNo int64

	err := row.Scan(&tableNo)
	if err == sql.ErrNoRows {
		sqlQuery = "SELECT table_id FROM `table` WHERE is_empty = 1 LIMIT 1;"
		row = DB.QueryRow(sqlQuery)

		row.Scan(&tableNo)

		return tableNo, nil
	} else if err != nil {
		return 0, err
	}
	return tableNo, nil
}

func PlaceNewOrder(o Order) (int64, error) {
	sqlQuery := "INSERT INTO `order` (order_at, table_no, customer_id, status, total_price) VALUES (?, ?, ?, ?, ?)"

	result, err := DB.Exec(sqlQuery, o.OrderAt, o.TableNo, o.CustomerId, o.Status, o.TotalPrice)
	if err != nil {
		return 0, fmt.Errorf("Error in creating order, %v", err)
	}
	orderId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error in fetching order id, %v", err)
	}
	return orderId, nil
}

func OccupyTable(tableNo int64) error {
	sqlQuery := "UPDATE `table` SET is_empty = 0 WHERE table_id = ?"

	_, err := DB.Exec(sqlQuery, tableNo)
	if err != nil {
		return fmt.Errorf("error occupying table : %v", err)
	}

	return err
}

func EntriesInItemOrder(orderSlice []ItemInCart, newOrder Order) error {
	sqlQuery := "INSERT INTO item_order (order_id, item_id, quantity, instruction) VALUES (?, ?, ?, ?);"

	for _, item := range orderSlice {
		_, err := DB.Exec(sqlQuery, newOrder.OrderId, item.ItemId, item.Quantity, item.Instruction)
		if err != nil {
			return fmt.Errorf("error inserting item in item-order : %v", err)
		}
	}
	return nil
}
