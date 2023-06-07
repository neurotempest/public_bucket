package go_sql_example_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/thecodedproject/sqltest"

	"github.com/neurotempest/public_bucket/go_sql_example"
)

func TestSimpleInsertAndQuery(t *testing.T) {

	testCases := []struct{
		Name string
		Input go_sql_example.Simple
		Expected go_sql_example.Simple
	}{
		{
			Name: "Insert without ID",
			Input: go_sql_example.Simple{
				A: 1234,
				B: "hello",
			},
			Expected: go_sql_example.Simple{
				ID: 1,
				A: 1234,
				B: "hello",
			},
		},
		{
			Name: "Insert with ID",
			Input: go_sql_example.Simple{
				ID: 23,
				A: 1234,
				B: "hello",
			},
			Expected: go_sql_example.Simple{
				ID: 23,
				A: 1234,
				B: "hello",
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()
			db := sqltest.OpenMysql(t, "schema.sql")

			id, err := go_sql_example.SimpleInsert(ctx, db, test.Input)
			require.NoError(t, err)

			actual, err := go_sql_example.SimpleQueryRowByID(ctx, db, id)
			require.NoError(t, err)

			require.Equal(t, test.Expected, actual)
		})
	}
}

func TestSimpleInsertMultipleAndQueryMultiple(t *testing.T) {

	testCases := []struct{
		Name string
		Inputs []go_sql_example.Simple
		Expected []go_sql_example.Simple
	}{
		{
			Name: "Insert without IDs",
			Inputs: []go_sql_example.Simple{
				{
					A: 1234,
					B: "hello",
				},
				{
					A: 5678,
					B: "world",
				},
			},
			Expected: []go_sql_example.Simple{
				{
					ID: 1,
					A: 1234,
					B: "hello",
				},
				{
					ID: 2,
					A: 5678,
					B: "world",
				},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()
			db := sqltest.OpenMysql(t, "schema.sql")

			err := go_sql_example.SimpleInsertMany(ctx, db, test.Inputs...)
			require.NoError(t, err)

			res, err := go_sql_example.SimpleQuery(ctx, db, nil)
			require.NoError(t, err)

			require.Equal(t, test.Expected, res)
		})
	}
}

func TestSimpleInsertMultipleAndQueryWithParams(t *testing.T) {

	testCases := []struct{
		Name string
		Inputs []go_sql_example.Simple
		QueryParams go_sql_example.Query
		Expected []go_sql_example.Simple
	}{
		{
			Name: "Insert two - query for 1",
			Inputs: []go_sql_example.Simple{
				{
					A: 1234,
					B: "hello",
				},
				{
					A: 5678,
					B: "world",
				},
			},
			QueryParams: go_sql_example.Query{
				"a": 1234,
				"b": "hello",
			},
			Expected: []go_sql_example.Simple{
				{
					ID: 1,
					A: 1234,
					B: "hello",
				},
			},
		},
		{
			Name: "Insert two - query for 2",
			Inputs: []go_sql_example.Simple{
				{
					A: 1234,
					B: "hello",
				},
				{
					A: 5678,
					B: "world",
				},
			},
			QueryParams: go_sql_example.Query{
				"b": "world",
			},
			Expected: []go_sql_example.Simple{
				{
					ID: 2,
					A: 5678,
					B: "world",
				},
			},
		},
		{
			Name: "Insert two - query for none",
			Inputs: []go_sql_example.Simple{
				{
					A: 1234,
					B: "hello",
				},
				{
					A: 5678,
					B: "world",
				},
			},
			QueryParams: go_sql_example.Query{
				"b": "jkfdsjkdlsa",
			},
			Expected: []go_sql_example.Simple{
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			ctx := context.Background()
			db := sqltest.OpenMysql(t, "schema.sql")

			err := go_sql_example.SimpleInsertMany(ctx, db, test.Inputs...)
			require.NoError(t, err)

			res, err := go_sql_example.SimpleQuery(ctx, db, test.QueryParams)
			require.NoError(t, err)

			require.Equal(t, test.Expected, res)
		})
	}
}
