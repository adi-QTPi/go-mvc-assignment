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
	OrderId      int64         `json:"order_id"`
	OrderAt      time.Time     `json:"order_at"`
	TableNo      sql.NullInt64 `json:"table_no"`
	CustomerId   string        `json:"customer_id"`
	CustomerName string        `json:"customer_name"`
	Status       string        `json:"status"`
	TotalPrice   int64         `json:"total_price"`
}

type ItemOrder struct {
	OrderId     int64          `json:"order_id"`
	ItemId      int64          `json:"item_id"`
	Quantity    int            `json:"quantity"`
	Instruction sql.NullString `json:"instruction"`
	IsComplete  string         `json:"is_complete"`
	CookId      sql.NullString `json:"cook_id"`
}

type KitchenOrder struct {
	OrderId     int64          `json:"order_id"`
	ItemId      int64          `json:"item_id"`
	ItemName    string         `json:"item_name"`
	Quantity    int            `json:"quantity"`
	Instruction sql.NullString `json:"instruction"`
	IsComplete  string         `json:"is_complete"`
	CookId      sql.NullString `json:"cook_id"`
	TableNo     sql.NullInt64  `json:"table_no"`
	OrderAt     time.Time      `json:"order_at"`
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

func FetchKitchenOrderForToday() ([]KitchenOrder, error) {
	var kitchenData []KitchenOrder

	sqlQuery := "SELECT io.order_id, io.item_id,i.item_name, io.quantity, io.instruction, io.is_complete, io.cook_id, o.table_no, o.order_at FROM item_order io JOIN item i ON io.item_id = i.item_id JOIN `order` o ON io.order_id = o.order_id WHERE DATE(o.order_at) = CURDATE() ORDER BY io.order_id DESC;"

	rows, err := DB.Query(sqlQuery)
	if err != nil {
		return kitchenData, fmt.Errorf("error fetching item order ingfo, %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var order KitchenOrder

		err := rows.Scan(&order.OrderId, &order.ItemId, &order.ItemName, &order.Quantity, &order.Instruction, &order.IsComplete, &order.CookId, &order.TableNo, &order.OrderAt)
		if err != nil {
			return nil, fmt.Errorf("error scanning the rows, %v", err)
		}

		kitchenData = append(kitchenData, order)
	}

	return kitchenData, nil
}

func FetchAllOrderDetailsByDate(dateStr string, xUser User) ([]Order, error) {
	var orderSlice []Order
	if xUser.Role == "admin" {
		sqlQuery := "SELECT o.order_id, u.user_id, u.name AS customer_name, o.table_no, o.order_at, o.status, o.total_price FROM `order` AS o JOIN `user` AS u ON o.customer_id = u.user_id WHERE DATE(o.order_at) = ? ORDER BY o.order_at DESC;"

		rows, err := DB.Query(sqlQuery, dateStr)
		if err != nil {
			return nil, fmt.Errorf("error scanning the rows, %v", err)
		}
		defer rows.Close()
		for rows.Next() {
			var o Order
			err := rows.Scan(&o.OrderId, &o.CustomerId, &o.CustomerName, &o.TableNo, &o.OrderAt, &o.Status, &o.TotalPrice)
			if err != nil {
				return nil, fmt.Errorf("error scanning the rows, %v", err)
			}
			orderSlice = append(orderSlice, o)
		}
	} else if xUser.Role == "customer" {
		sqlQuery := "SELECT o.order_id, u.user_id, u.name AS customer_name, o.table_no, o.order_at, o.status, o.total_price FROM `order` AS o JOIN `user` AS u ON o.customer_id = u.user_id WHERE o.customer_id = ? AND DATE(o.order_at) = ? ORDER BY o.order_at DESC;"

		rows, err := DB.Query(sqlQuery, xUser.UserId, dateStr)
		if err != nil {
			return nil, fmt.Errorf("error fetching orders, %v", err)
		}
		defer rows.Close()
		for rows.Next() {
			var o Order
			err := rows.Scan(&o.OrderId, &o.CustomerId, &o.CustomerName, &o.TableNo, &o.OrderAt, &o.Status, &o.TotalPrice)
			if err != nil {
				return nil, fmt.Errorf("error scanning the rows, %v", err)
			}
			orderSlice = append(orderSlice, o)
		}
	}

	return orderSlice, nil
}
