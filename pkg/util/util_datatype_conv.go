package util

import (
	"database/sql"
	"fmt"
	"strconv"
)

func StringToSqlNullInt64(s string) (sql.NullInt64, error) {
	var ans sql.NullInt64
	parsedInt, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return ans, fmt.Errorf("Invalid input")
	}
	ans.Int64 = parsedInt
	ans.Valid = true

	return ans, nil
}
