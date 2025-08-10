package models

import (
	"database/sql"
	"fmt"
)

type Category struct {
	CategoryId          int64          `json:"cat_id"`
	CategoryName        string         `json:"cat_name"`
	CategoryDescription sql.NullString `json:"cat_description"`
}

func GetAllCategories() ([]Category, error) {
	sqlQuery := "SELECT * FROM category;"

	rows, err := DB.Query(sqlQuery)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var fetchedCategories []Category

	for rows.Next() {
		var category Category
		err := rows.Scan(&category.CategoryId, &category.CategoryName, &category.CategoryDescription)

		if err != nil {
			return nil, fmt.Errorf("error fetching items -> %v", err)
		}
		fetchedCategories = append(fetchedCategories, category)
	}

	return fetchedCategories, nil
}

func AddCategory(c Category) (bool, error) {
	sqlQuery := "INSERT INTO category (cat_name, cat_description)SELECT ?, ? WHERE NOT EXISTS (SELECT 1 FROM category WHERE cat_name = ?);"

	result, err := DB.Exec(sqlQuery, c.CategoryName, c.CategoryDescription, c.CategoryName)
	if err != nil {
		return true, fmt.Errorf("Error in adding category, %v", err)
	}

	a, err := result.RowsAffected()
	if err != nil {
		return true, fmt.Errorf("Error while verifying effect of query, %v", err)
	}
	if a == 0 {
		return false, nil
	}

	return true, nil
}
