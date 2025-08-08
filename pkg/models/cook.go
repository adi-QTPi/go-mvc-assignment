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

func SyncOrderStatus() error {
	sqlQuery := "UPDATE `order` o SET status = 'cooking' WHERE EXISTS (SELECT 1 FROM item_order io WHERE io.order_id = o.order_id AND io.is_complete = 'taken'); UPDATE `order` o SET status = 'ready_to_serve' WHERE NOT EXISTS (SELECT 1 FROM item_order io WHERE io.order_id = o.order_id AND io.is_complete != 'complete') AND o.status != 'paid';"

	_, err := DB.Exec(sqlQuery)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	return nil
}
