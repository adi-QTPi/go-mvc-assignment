package models

import (
	"database/sql"
	"fmt"
	"strconv"
)

type Item struct {
	ItemId      int64          `json:"item_id"`
	ItemName    string         `json:"item_name"`
	CookTimeMin sql.NullInt64  `json:"cook_time_min"`
	Price       int64          `json:"price"`
	DisplayPic  sql.NullString `json:"display_pic"`
	CatId       int64          `json:"cat_id"`
	SubCatId    sql.NullInt64  `json:"subcat_id"`
	IsAvailable bool           `json:"is_available"`
}

func GetAllItems() ([]Item, error) {
	sqlQuery := "SELECT * FROM item;"

	rows, err := DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var fetchedItems []Item

	for rows.Next() {
		var oneItem Item
		err := rows.Scan(&oneItem.ItemId, &oneItem.ItemName, &oneItem.CookTimeMin, &oneItem.Price, &oneItem.DisplayPic, &oneItem.CatId, &oneItem.SubCatId, &oneItem.IsAvailable)

		if err != nil {
			return nil, fmt.Errorf("error fetching items -> %v", err)
		}
		fetchedItems = append(fetchedItems, oneItem)
	}

	return fetchedItems, nil
}

func AddItem(i Item) error {
	sqlQuery := "INSERT INTO item (item_name, cat_id, cook_time_min, price, display_pic, subcat_id) VALUES (?, ?, ?, ?, ?, ?);"

	_, err := DB.Exec(sqlQuery, i.ItemName, i.CatId, i.CookTimeMin, i.Price, i.DisplayPic, i.SubCatId)
	if err != nil {
		return fmt.Errorf("Error in adding item, %v", err)
	}

	return nil
}

func DeleteItemById(idString string) error {
	itemId, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid input for itemId : %v", err)
	}

	sqlQuery := "UPDATE item SET is_available = 0 WHERE item_id = ?;"

	_, err = DB.Exec(sqlQuery, itemId)
	if err != nil {
		return fmt.Errorf("Error in adding item, %v", err)
	}

	return nil
}
