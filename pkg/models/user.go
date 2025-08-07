package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	Name     string `json:"name"`
	PwdHash  string `json:"pwd_hash"`
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
	sql_query := "SELECT user_id, user_name, name, role FROM user WHERE user_id = ?;"

	row := DB.QueryRow(sql_query, id)
	// defer row.Close()

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

func GetUserByUsername(userName string) (bool, error) {
	sql_query := "SELECT user_id FROM user WHERE user_name = ?;"

	row := DB.QueryRow(sql_query, userName)

	var userId string
	err := row.Scan(&userId)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, fmt.Errorf("error in running query %v", err)
	}

	return true, nil
}

func DeleteUserById(id string) error {
	sql_query := "DELETE FROM user WHERE user_id = ?;"

	_, err := DB.Exec(sql_query, id)
	if err != nil {
		return fmt.Errorf("unable to delete user, %v", err)
	}

	return nil
}

func AddNewUser(u User) error {
	sql_query := "INSERT INTO user (user_name, name, pwd_hash, role) VALUES (?,?,?,?)"

	_, err := DB.Exec(sql_query, u.UserName, u.Name, u.PwdHash, u.Role)
	if err != nil {
		return fmt.Errorf("unable to delete user, %v", err)
	}

	return nil
}
