package models

import (
	"database/sql"
	"fmt"
)

type Category struct {
	CatId          int64          `json:"cat_id"`
	CatName        string         `json:"cat_name"`
	CatDescription sql.NullString `json:"cat_description"`
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
		var oneCat Category
		err := rows.Scan(&oneCat.CatId, &oneCat.CatName, &oneCat.CatDescription)

		if err != nil {
			return nil, fmt.Errorf("error fetching items -> %v", err)
		}
		fetchedCategories = append(fetchedCategories, oneCat)
	}

	return fetchedCategories, nil
}

func AddCategory(c Category) (bool, error) {
	sqlQuery := "INSERT INTO category (cat_name, cat_description)SELECT ?, ? WHERE NOT EXISTS (SELECT 1 FROM category WHERE cat_name = ?);"

	result, err := DB.Exec(sqlQuery, c.CatName, c.CatDescription, c.CatName)
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
