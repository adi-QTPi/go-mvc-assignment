package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

func GetAllUsers() ([]User, error) {
	sql_query := "SELECT user_id, user_name, name, role FROM user;"
	rows, err := DB.Query(sql_query)
	if err != nil {
		return nil, fmt.Errorf("error querying users: %v", err)
	}
	defer rows.Close()

	var fetchedUsers []User

	for rows.Next() {
		var oneUser User
		err := rows.Scan(&oneUser.UserId, &oneUser.UserName, &oneUser.Name, &oneUser.Role)
		if err != nil {
			return nil, fmt.Errorf("error fetching users -> %v", err)
		}
		fetchedUsers = append(fetchedUsers, oneUser)
	}

	return fetchedUsers, nil
}

func GetUserById(id string) (*User, error) {
	sql_query := "SELECT user_name, name, profile_pic, role FROM user WHERE user_id = ?;"

	row := DB.QueryRow(sql_query, id)

	var fetchedUser User

	err := row.Scan(&fetchedUser.UserId, &fetchedUser.UserName, &fetchedUser.Name, &fetchedUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error scanning user: %v", err)
	}

	return &fetchedUser, nil
}
