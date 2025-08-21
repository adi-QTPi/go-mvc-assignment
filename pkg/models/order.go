package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/cache"
)

type ItemInCart struct {
	ItemId      int64  `json:"item_id"`
	Quantity    int64  `json:"quantity"`
	Price       int64  `json:"price"`
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

type ItemOrderDescriptive struct {
	OrderId     int64          `json:"order_id"`
	ItemId      int64          `json:"item_id"`
	ItemName    string         `json:"item_name"`
	ItemPrice   int            `json:"price"`
	Quantity    int            `json:"quantity"`
	Instruction sql.NullString `json:"instruction"`
	IsComplete  string         `json:"is_complete"`
	CookId      sql.NullString `json:"cook_id"`
	CookName    sql.NullString `json:"cook_name"`
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

func PlaceNewOrder(o Order, tx *sql.Tx) (int64, error) {
	sqlQuery := "INSERT INTO `order` (order_at, table_no, customer_id, status, total_price) VALUES (?, ?, ?, ?, ?);"

	result, err := tx.Exec(sqlQuery, o.OrderAt, o.TableNo, o.CustomerId, o.Status, o.TotalPrice)
	if err != nil {
		return 0, fmt.Errorf("Error in creating order, %v", err)
	}
	orderId, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error in fetching order id, %v", err)
	}

	cacheField := fmt.Sprintf("orders%s", time.Now().Format("2006-01-02"))
	cache.AppCache.Delete(cacheField)
	return orderId, nil
}

func OccupyTable(tableNo int64, tx *sql.Tx) error {
	sqlQuery := "UPDATE `table` SET is_empty = 0 WHERE table_id = ?"

	_, err := tx.Exec(sqlQuery, tableNo)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error occupying table : %v", err)
	}

	return err
}

func VacateTable() error {
	sqlQuery := "UPDATE `table` SET is_empty = 1 WHERE table_id IN (SELECT DISTINCT table_no FROM `order` o1 WHERE table_no IS NOT NULL AND NOT EXISTS (SELECT 1 FROM `order` o2 WHERE o2.table_no = o1.table_no AND o2.status != 'paid'));"

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	return nil
}

func EntriesInItemOrder(orderSlice []ItemInCart, newOrder Order, tx *sql.Tx) error {

	orderId := newOrder.OrderId

	var args []any
	var qMarks string

	for k, v := range orderSlice {
		qMarks += "(?,?,?,?)"
		if k < (len(orderSlice) - 1) {
			qMarks += ","
		}
		args = append(args, orderId, v.ItemId, v.Quantity, v.Instruction)
	}

	sqlQuery := "INSERT INTO item_order (order_id, item_id, quantity, instruction) VALUES "

	_, err := tx.Exec(sqlQuery+qMarks, args...)
	if err != nil {
		return fmt.Errorf("error inserting items in item-order : %v", err)
	}

	return nil
}

func FetchKitchenOrderForToday() ([]KitchenOrder, error) {
	var kitchenData []KitchenOrder

	d := time.Now().Local()
	todayDate := d.Format("2006-01-02")

	sqlQuery := "SELECT io.order_id, io.item_id,i.item_name, io.quantity, io.instruction, io.is_complete, io.cook_id, o.table_no, o.order_at FROM item_order io JOIN item i ON io.item_id = i.item_id JOIN `order` o ON io.order_id = o.order_id WHERE DATE(o.order_at) = ? ORDER BY io.order_id DESC;"

	rows, err := DB.Query(sqlQuery, todayDate)
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
		cacheField := fmt.Sprintf("orders%s", dateStr)
		orders, ok := cache.AppCache.Get(cacheField)
		if ok {
			orderSlice = orders.([]Order)
			return orderSlice, nil
		}

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

		cache.AppCache.Set(cacheField, orderSlice, 24*time.Hour)

		cache.AppCache.Set(cacheField, orderSlice, 24*time.Hour)

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

func MakePayment(orderId string, customerReview string) error {

	order, err := FetchOrderByOrderId(orderId)
	if err != nil {
		return fmt.Errorf("error in scanning for order struct : , %v", err)
	}

	sqlQuery := "INSERT INTO paid_orders (order_id, customer_review, total_amount) VALUES (?, ?, ?);"

	_, err = DB.Exec(sqlQuery, order.OrderId, customerReview, order.TotalPrice)
	if err != nil {
		return fmt.Errorf("error inseting in paid_orders : , %v", err)
	}

	sqlQuery = "UPDATE `order` SET status = 'paid' WHERE order_id = ?;"

	_, err = DB.Exec(sqlQuery, order.OrderId)
	if err != nil {
		return fmt.Errorf("error in changing order status : , %v", err)
	}

	cacheField := fmt.Sprintf("orders%s", order.OrderAt.Format("2006-01-02"))
	cache.AppCache.Delete(cacheField)

	return nil
}

func FetchOrderByOrderId(orderId string) (Order, error) {
	var order Order

	sqlQuery := "SELECT o.order_id, u.user_id, u.name AS customer_name, o.table_no, o.order_at, o.status, o.total_price FROM `order` AS o JOIN `user` AS u ON o.customer_id = u.user_id WHERE o.order_id = ?"

	row := DB.QueryRow(sqlQuery, orderId)
	err := row.Scan(&order.OrderId, &order.CustomerId, &order.CustomerName, &order.TableNo, &order.OrderAt, &order.Status, &order.TotalPrice)
	if err != nil {
		return order, fmt.Errorf("error in scanning for order struct : , %v", err)
	}

	return order, nil
}

func FetchBillDetailsByOrderId(orderId string) ([]ItemOrderDescriptive, error) {
	sqlQuery := "SELECT o.order_id, io.item_id, i.item_name, i.price, io.quantity, io.instruction, io.is_complete, io.cook_id, cook.name AS cook_name FROM `order` o INNER JOIN item_order io ON o.order_id = io.order_id INNER JOIN item i ON io.item_id = i.item_id LEFT JOIN user cook ON io.cook_id = cook.user_id WHERE o.order_id = ?;"

	rows, err := DB.Query(sqlQuery, orderId)
	if err != nil {
		return nil, fmt.Errorf("error in fetching order meta data and itemorderdescriptives : %v", err)
	}
	defer rows.Close()

	var billData []ItemOrderDescriptive
	for rows.Next() {
		var i ItemOrderDescriptive
		err := rows.Scan(&i.OrderId, &i.ItemId, &i.ItemName, &i.ItemPrice, &i.Quantity, &i.Instruction, &i.IsComplete, &i.CookId, &i.CookName)
		if err != nil {
			return nil, fmt.Errorf("error scanning the rows, %v", err)
		}
		billData = append(billData, i)
	}

	return billData, nil
}
