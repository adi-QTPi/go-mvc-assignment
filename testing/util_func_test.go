package test

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/adi-QTPi/go-mvc-assignment/pkg/util"
)

func TestStringToSqlNullInt64(t *testing.T) {
	var testcases = []struct {
		caseName       string
		givenInput     string
		expectedOutput sql.NullInt64
		errExpected    bool
		errDescription error
	}{
		{
			caseName:   "non null int",
			givenInput: "12345",
			expectedOutput: sql.NullInt64{
				Int64: 12345,
				Valid: true,
			},
			errExpected: false,
		},
		{
			caseName:   "empty string",
			givenInput: "",
			expectedOutput: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
			errExpected:    true,
			errDescription: fmt.Errorf("Invalid input"),
		},
		{
			caseName:   "invalid string",
			givenInput: "afkdsasdfas",
			expectedOutput: sql.NullInt64{
				Int64: 0,
				Valid: false,
			},
			errExpected:    true,
			errDescription: fmt.Errorf("Invalid input"),
		},
		{
			caseName:   "negative int",
			givenInput: "-987",
			expectedOutput: sql.NullInt64{
				Int64: -987,
				Valid: true,
			},
			errExpected: false,
		},
	}

	for _, v := range testcases {
		t.Run(v.caseName, func(t *testing.T) {
			ans, err := util.StringToSqlNullInt64(v.givenInput)

			if err != nil {
				if err.Error() != v.errDescription.Error() {
					t.Errorf("expected error %v, got error %v", v.errDescription, err)
					return
				}
			}

			if ans != v.expectedOutput {
				t.Errorf("Expected %d , got %d", v.expectedOutput.Int64, ans.Int64)
			}
		})
	}
}
