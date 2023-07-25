package go_sql_example_test

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thecodedproject/sqltest"
)

type Custom struct {
	V int64
}

func (c Custom) String() string {
	return "Custom" + strconv.FormatInt(c.V, 10)
}

func (c Custom) Value() (driver.Value, error) {
	return c.String(), nil
}

func (c *Custom) Scan(v any) error {
	var s sql.NullString
	if err := s.Scan(v); err != nil {
		return err
	}

	if !s.Valid {
		return errors.New("not valid string")
	}

	if !strings.HasPrefix(s.String, "Custom") {
		return errors.New("not custom type")
	}

	sStr := strings.TrimPrefix(s.String, "Custom")

	i, err := strconv.ParseInt(sStr, 10, 64)
	if err != nil {
		return err
	}

	c.V = i

	return nil
}

func TestInsertCustomTypeIntoVarchar(t *testing.T) {

	testCases := []struct{
		Name string
		Val any
		Expected string
	}{
		{
			Name: "a string",
			Val: string("hello world"),
			Expected: "hello world",
		},
		{
			Name: "a valuer type",
			Val: Custom{V: 100},
			Expected: "Custom100",
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()
			db := sqltest.OpenMysql(t, "schema.sql")

			r, err := db.ExecContext(
				ctx,
				"insert into conversion (vc) values (?)",
				test.Val,
			)
			require.NoError(t, err)

			id, err := r.LastInsertId()
			require.NoError(t, err)

			var res string
			err = db.QueryRowContext(
				ctx,
				"select vc from conversion where id = ?",
				id,
			).Scan(
				&res,
			)
			require.NoError(t, err)

			require.Equal(t, test.Expected, res)
		})
	}

}

func TestScanIntoCustomTypeFromVarchar(t *testing.T) {

	testCases := []struct{
		Name string
		Val Custom
		Expected Custom
	}{
		{
			Name: "a valuer type",
			Val: Custom{V: 100},
			Expected: Custom{V: 100},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()
			db := sqltest.OpenMysql(t, "schema.sql")

			r, err := db.ExecContext(
				ctx,
				"insert into conversion (vc) values (?)",
				test.Val,
			)
			require.NoError(t, err)

			id, err := r.LastInsertId()
			require.NoError(t, err)

			var res Custom
			err = db.QueryRowContext(
				ctx,
				"select vc from conversion where id = ?",
				id,
			).Scan(
				&res,
			)
			require.NoError(t, err)

			require.Equal(t, test.Expected, res)
		})
	}

}
