package user_test

import (
	context "context"
	fmt "fmt"
	libdb "github.com/neurotempest/public_bucket/go_fsm/libdb"
	user "github.com/neurotempest/public_bucket/go_fsm/libdb/user"
	require "github.com/stretchr/testify/require"
	lib "github.com/thecodedproject/dbcrudgen/lib"
	assert "github.com/thecodedproject/gotest/assert"
	gotest_time "github.com/thecodedproject/gotest/time"
	sqltest "github.com/thecodedproject/sqltest"
	testing "testing"
	time "time"
)

func populateDataModelFromNonce(nonce int64) libdb.User {

	return libdb.User{
		CreatedAt: time.Unix(nonce, 0),
		DeletedAt: time.Unix(nonce, 0),
		FirstName: "some_str" + fmt.Sprint(nonce),
		LastName: "some_str" + fmt.Sprint(nonce),
		State: libdb.UserState(nonce),
	}
}

func populateDataModelFromNonceWithIDAndTimestamp(
	nonce int64,
	id int64,
	t time.Time,
) libdb.User {

	d := populateDataModelFromNonce(nonce)
	d.ID = id
	d.UpdatedAt = t.Round(gotest_time.Second)
	return d
}

func queryFromNonce(nonce int64) map[string]any {

	return map[string]any{
		"created_at": time.Unix(nonce, 0),
		"deleted_at": time.Unix(nonce, 0),
		"first_name": "some_str" + fmt.Sprint(nonce),
		"last_name": "some_str" + fmt.Sprint(nonce),
		"state": libdb.UserState(nonce),
	}
}

func TestInsertAndSelect(t *testing.T) {

	now := gotest_time.SetTimeNowForTesting(t)

	testCases := []struct{
		Name string
		ToInsert []libdb.User
		Query map[string]any
		Expected []libdb.User
		ExpectErr bool
	}{
		{
			Name: "selects nothing when nothing inserted",
		},
		{
			Name: "insert one and select",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(11),
			},
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(11, 1, now),
			},
		},
		{
			Name: "insert many and select",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(11),
				populateDataModelFromNonce(21),
				populateDataModelFromNonce(31),
				populateDataModelFromNonce(41),
			},
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(11, 1, now),
				populateDataModelFromNonceWithIDAndTimestamp(21, 2, now),
				populateDataModelFromNonceWithIDAndTimestamp(31, 3, now),
				populateDataModelFromNonceWithIDAndTimestamp(41, 4, now),
			},
		},
		{
			Name: "insert many and select with query",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(22),
				populateDataModelFromNonce(45),
				populateDataModelFromNonce(45),
				populateDataModelFromNonce(1),
				populateDataModelFromNonce(45),
			},
			Query: queryFromNonce(45),
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(45, 2, now),
				populateDataModelFromNonceWithIDAndTimestamp(45, 3, now),
				populateDataModelFromNonceWithIDAndTimestamp(45, 5, now),
			},
		},
		{
			Name: "select query field which is not in data model returns error",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(1),
			},
			Query: map[string]any{
				"some_field_not_in_User": 1,
			},
			ExpectErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			db := sqltest.OpenMysql(t, "schema.sql")
			ctx := lib.ContextWithDB(context.Background(), db)

			for _, d := range test.ToInsert {
				_, err := user.Insert(ctx, d)
				require.NoError(t, err)
			}

			actual, err := user.Select(ctx, test.Query)
			if test.ExpectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			require.Equal(t, len(test.Expected), len(actual))

			for i := range actual {
				assert.LogicallyEqual(t, test.Expected[i], actual[i], fmt.Sprint(i) + "th element not equal")
			}
		})
	}
}

func TestSelectByID(t *testing.T) {

	now := gotest_time.SetTimeNowForTesting(t)

	testCases := []struct{
		Name string
		ToInsert []libdb.User
		ID int64
		Expected libdb.User
		ExpectErr bool
	}{
		{
			Name: "when ID not found returns error",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(100),
				populateDataModelFromNonce(200),
			},
			ID: 12345,
			ExpectErr: true,
		},
		{
			Name: "when ID is found returns row",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(100),
				populateDataModelFromNonce(200),
				populateDataModelFromNonce(300),
			},
			ID: 2,
			Expected: populateDataModelFromNonceWithIDAndTimestamp(200, 2, now),
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			db := sqltest.OpenMysql(t, "schema.sql")
			ctx := lib.ContextWithDB(context.Background(), db)

			for _, d := range test.ToInsert {
				_, err := user.Insert(ctx, d)
				require.NoError(t, err)
			}

			actual, err := user.SelectByID(ctx, test.ID)
			if test.ExpectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.LogicallyEqual(t, test.Expected, actual)
		})
	}
}

func TestUpdate(t *testing.T) {

	now := gotest_time.SetTimeNowForTesting(t)

	testCases := []struct{
		Name string
		ToInsert []libdb.User
		Updates map[string]any
		Query map[string]any
		ExpectedNumUpdates int64
		Expected []libdb.User
		ExpectErr bool
	}{
		{
			Name: "empty params does nothing",
		},
		{
			Name: "update unknown field throws error",
			Updates: map[string]any{
				"field_not_in_the_libdb.User_type": "update",
			},
			ExpectErr: true,
		},
		{
			Name: "query unknown field throws error",
			Updates: queryFromNonce(1),
			Query: map[string]any{
				"field_not_in_User": "update",
			},
			ExpectErr: true,
		},
		{
			Name: "update all records",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(123),
				populateDataModelFromNonce(124),
				populateDataModelFromNonce(125),
				populateDataModelFromNonce(126),
			},
			Updates: queryFromNonce(111),
			ExpectedNumUpdates: 4,
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(111, 1, now),
				populateDataModelFromNonceWithIDAndTimestamp(111, 2, now),
				populateDataModelFromNonceWithIDAndTimestamp(111, 3, now),
				populateDataModelFromNonceWithIDAndTimestamp(111, 4, now),
			},
		},
		{
			Name: "update records with query",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(123),
				populateDataModelFromNonce(125),
				populateDataModelFromNonce(124),
				populateDataModelFromNonce(125),
				populateDataModelFromNonce(126),
				populateDataModelFromNonce(125),
			},
			Updates: queryFromNonce(999),
			Query: queryFromNonce(125),
			ExpectedNumUpdates: 3,
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(123, 1, now),
				populateDataModelFromNonceWithIDAndTimestamp(999, 2, now),
				populateDataModelFromNonceWithIDAndTimestamp(124, 3, now),
				populateDataModelFromNonceWithIDAndTimestamp(999, 4, now),
				populateDataModelFromNonceWithIDAndTimestamp(126, 5, now),
				populateDataModelFromNonceWithIDAndTimestamp(999, 6, now),
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			db := sqltest.OpenMysql(t, "schema.sql")
			ctx := lib.ContextWithDB(context.Background(), db)

			for _, d := range test.ToInsert {
				_, err := user.Insert(ctx, d)
				require.NoError(t, err)
			}

			numUpdates, err := user.Update(ctx, test.Updates, test.Query)
			if test.ExpectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			require.Equal(t, test.ExpectedNumUpdates, numUpdates)

			actual, err := user.Select(ctx, nil)
			require.NoError(t, err)

			require.Equal(t, len(test.Expected), len(actual))

			for i := range actual {
				assert.LogicallyEqual(t, test.Expected[i], actual[i], fmt.Sprint(i) + "th element not equal")
			}
		})
	}
}

func TestUpdateByID(t *testing.T) {

	now := gotest_time.SetTimeNowForTesting(t)

	testCases := []struct{
		Name string
		ToInsert []libdb.User
		ID int64
		Updates map[string]any
		Expected []libdb.User
		ExpectErr bool
	}{
		{
			Name: "no updates does not error - even if ID does not exist",
		},
		{
			Name: "when there are updates and ID not found throws error",
			ID: 1234,
			Updates: queryFromNonce(1),
			ExpectErr: true,
		},
		{
			Name: "when update field not in schema throws error",
			ID: 1,
			Updates: map[string]any{
				"field_not_in_the_libdb.User_type": "update",
			},
			ExpectErr: true,
		},
		{
			Name: "insert many and update one by id",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(101),
				populateDataModelFromNonce(102),
				populateDataModelFromNonce(103),
				populateDataModelFromNonce(104),
			},
			ID: 3,
			Updates: queryFromNonce(555),
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(101, 1, now),
				populateDataModelFromNonceWithIDAndTimestamp(102, 2, now),
				populateDataModelFromNonceWithIDAndTimestamp(555, 3, now),
				populateDataModelFromNonceWithIDAndTimestamp(104, 4, now),
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			db := sqltest.OpenMysql(t, "schema.sql")
			ctx := lib.ContextWithDB(context.Background(), db)

			for _, d := range test.ToInsert {
				_, err := user.Insert(ctx, d)
				require.NoError(t, err)
			}

			err := user.UpdateByID(ctx, test.ID, test.Updates)

			if test.ExpectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			actual, err := user.Select(ctx, nil)
			require.NoError(t, err)

			require.Equal(t, len(test.Expected), len(actual))

			for i := range actual {
				assert.LogicallyEqual(t, test.Expected[i], actual[i], fmt.Sprint(i) + "th element not equal")
			}
		})
	}
}

func TestDelete(t *testing.T) {

	now := gotest_time.SetTimeNowForTesting(t)

	testCases := []struct{
		Name string
		ToInsert []libdb.User
		Query map[string]any
		ExpectedNumDeleted int64
		Expected []libdb.User
		ExpectErr bool
	}{
		{
			Name: "empty query deletes all records",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(1000),
				populateDataModelFromNonce(1001),
				populateDataModelFromNonce(1002),
				populateDataModelFromNonce(1003),
				populateDataModelFromNonce(1004),
			},
			ExpectedNumDeleted: 5,
		},
		{
			Name: "delete records using query",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(1000),
				populateDataModelFromNonce(1001),
				populateDataModelFromNonce(1002),
				populateDataModelFromNonce(1002),
				populateDataModelFromNonce(1003),
				populateDataModelFromNonce(1004),
				populateDataModelFromNonce(1002),
			},
			Query: queryFromNonce(1002),
			ExpectedNumDeleted: 3,
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(1000, 1, now),
				populateDataModelFromNonceWithIDAndTimestamp(1001, 2, now),
				populateDataModelFromNonceWithIDAndTimestamp(1003, 5, now),
				populateDataModelFromNonceWithIDAndTimestamp(1004, 6, now),
			},
		},
		{
			Name: "when query contains field not in data model returns error",
			Query: map[string]any{
				"some_field_not_in_User": 1,
			},
			ExpectErr: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			db := sqltest.OpenMysql(t, "schema.sql")
			ctx := lib.ContextWithDB(context.Background(), db)

			for _, d := range test.ToInsert {
				_, err := user.Insert(ctx, d)
				require.NoError(t, err)
			}

			numDeleted, err := user.Delete(ctx, test.Query)
			if test.ExpectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			require.Equal(t, test.ExpectedNumDeleted, numDeleted)

			actual, err := user.Select(ctx, nil)
			require.NoError(t, err)

			require.Equal(t, len(test.Expected), len(actual))

			for i := range actual {
				assert.LogicallyEqual(t, test.Expected[i], actual[i], fmt.Sprint(i) + "th element not equal")
			}
		})
	}
}

func TestDeleteByID(t *testing.T) {

	now := gotest_time.SetTimeNowForTesting(t)

	testCases := []struct{
		Name string
		ToInsert []libdb.User
		ID int64
		Expected []libdb.User
		ExpectErr bool
	}{
		{
			Name: "when ID not found returns error",
			ExpectErr: true,
		},
		{
			Name: "insert many and delete by ID",
			ToInsert: []libdb.User{
				populateDataModelFromNonce(101),
				populateDataModelFromNonce(102),
				populateDataModelFromNonce(103),
				populateDataModelFromNonce(104),
			},
			ID: 3,
			Expected: []libdb.User{
				populateDataModelFromNonceWithIDAndTimestamp(101, 1, now),
				populateDataModelFromNonceWithIDAndTimestamp(102, 2, now),
				populateDataModelFromNonceWithIDAndTimestamp(104, 4, now),
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.Name, func(t *testing.T) {
			db := sqltest.OpenMysql(t, "schema.sql")
			ctx := lib.ContextWithDB(context.Background(), db)

			for _, d := range test.ToInsert {
				_, err := user.Insert(ctx, d)
				require.NoError(t, err)
			}

			err := user.DeleteByID(ctx, test.ID)

			if test.ExpectErr {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			actual, err := user.Select(ctx, nil)
			require.NoError(t, err)

			require.Equal(t, len(test.Expected), len(actual))

			for i := range actual {
				assert.LogicallyEqual(t, test.Expected[i], actual[i], fmt.Sprint(i) + "th element not equal")
			}
		})
	}
}

