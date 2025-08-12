package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	UserId     string `json:"user_id"`
	UserName   string `json:"user_name"`
	Name       string `json:"name"`
	Role       string `json:"role"`
	DisplayPic string `json:"display_pic"`
}

func GetAllUsers() ([]User, error) {
	sqlQuery := "SELECT user_id, user_name, name, role FROM user;"
	rows, err := DB.Query(sqlQuery)
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

func GetUserById(id string) (User, error) {
	sqlQuery := "SELECT user_id, user_name, name, role FROM user WHERE user_id = ?;"

	row := DB.QueryRow(sqlQuery, id)
	// defer row.Close() not work here

	var fetchedUser User

	err := row.Scan(&fetchedUser.UserId, &fetchedUser.UserName, &fetchedUser.Name, &fetchedUser.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return fetchedUser, fmt.Errorf("user not found")
		}
		return fetchedUser, fmt.Errorf("error scanning user: %v", err)
	}

	return fetchedUser, nil
}

func GetUserByUsername(userName string) (bool, error, string) {
	sqlQuery := "SELECT user_id FROM user WHERE user_name = ?;"

	row := DB.QueryRow(sqlQuery, userName)

	var userId string
	err := row.Scan(&userId)
	if err == sql.ErrNoRows {
		//no usr found case
		return false, nil, userId
	} else if err != nil {
		return false, fmt.Errorf("error in running query %v", err), userId
	}

	return true, nil, userId
}

func DeleteUserById(id string) error {
	sqlQuery := "DELETE FROM user WHERE user_id = ?;"

	_, err := DB.Exec(sqlQuery, id)
	if err != nil {
		return fmt.Errorf("unable to delete user, %v", err)
	}

	return nil
}

func FetchHashedPassword(id string) string {
	sqlQuery := "SELECT pwd_hash FROM user WHERE user_id = ?;"

	row := DB.QueryRow(sqlQuery, id)
	var pwd_hash string
	err := row.Scan(&pwd_hash)
	if err != nil {
		return ""
	}

	return pwd_hash

}

func AddNewUser(u User, pwdHash string) error {
	sqlQuery := "INSERT INTO user (user_name, name, pwd_hash, role) VALUES (?,?,?,?)"

	_, err := DB.Exec(sqlQuery, u.UserName, u.Name, pwdHash, u.Role)
	if err != nil {
		return fmt.Errorf("unable to add user, %v", err)
	}

	return nil
}
