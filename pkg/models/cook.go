package models

import "fmt"

func StatusUpdateByCook(cookId string, orderId int64, itemId int64, isComplete string) error {
	sqlQuery := "UPDATE item_order SET is_complete = ?, cook_id = ? WHERE order_id = ? AND item_id = ?;"

	_, err := DB.Exec(sqlQuery, isComplete, cookId, orderId, itemId)
	if err != nil {
		return fmt.Errorf("error changing kitchen item status , %v", err)
	}

	return nil
}
