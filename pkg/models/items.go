package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/adi-QTPi/go-mvc-assignment/cache"
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

type DisplayItem struct {
	ItemId      int64          `json:"item_id"`
	ItemName    string         `json:"item_name"`
	CookTimeMin string         `json:"cook_time_min"`
	Price       string         `json:"price"`
	DisplayPic  sql.NullString `json:"display_pic"`
	CatId       int64          `json:"cat_id"`
	Category    string         `json:"category"`
	SubCatId    int64          `json:"subcat_id"`
	SubCategory string         `json:"subcategory"`
}

func GetAllItems() ([]DisplayItem, error) {
	var fetchedItems []DisplayItem
	if fetchedItems, ok := cache.AppCache.Get("menu"); ok {
		return fetchedItems.([]DisplayItem), nil
	}
	sqlQuery := "SELECT i.item_id, i.item_name, i.cook_time_min, i.price, i.display_pic, i.cat_id, c.cat_name AS cat_name, i.subcat_id, cd.cat_name AS subcat_name FROM item i JOIN category c ON i.cat_id = c.cat_id LEFT JOIN category cd ON i.subcat_id = cd.cat_id WHERE i.is_available = 1"

	rows, err := DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var oneItem DisplayItem
		var displayPic sql.NullString
		err := rows.Scan(&oneItem.ItemId, &oneItem.ItemName, &oneItem.CookTimeMin, &oneItem.Price, &displayPic, &oneItem.CatId, &oneItem.Category, &oneItem.SubCatId, &oneItem.SubCategory)
		if err != nil {
			return nil, fmt.Errorf("error fetching items -> %v", err)
		}

		if displayPic.Valid {
			oneItem.DisplayPic = displayPic
		}
		fetchedItems = append(fetchedItems, oneItem)
	}

	cache.AppCache.Set("menu", fetchedItems, 24*time.Hour)

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
