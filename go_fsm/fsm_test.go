package go_fsm_test

import (
	"context"
	"testing"
	"time"

	"github.com/thecodedproject/sqltest"
	"github.com/stretchr/testify/require"
	"github.com/thecodedproject/dbcrudgen/lib"

	"github.com/neurotempest/public_bucket/go_fsm/libdb"
	"github.com/neurotempest/public_bucket/go_fsm/libdb/user"
)

const schemaFile = "libdb/user/schema.sql"

func TestCreateUser(t *testing.T) {

	db := sqltest.OpenMysql(t, schemaFile)
	ctx := lib.ContextWithDB(context.Background(), db)

	id, err := user.Create{
		FirstName: "Aaa",
		LastName: "Bbb",
	}.Exec(ctx)
	require.NoError(t, err)

	user, err := user.SelectByID(ctx, id)
	require.NoError(t, err)

	require.Equal(t, libdb.UserStateCreated, user.State)
}

func TestCreatedToActive(t *testing.T) {

	db := sqltest.OpenMysql(t, schemaFile)
	ctx := lib.ContextWithDB(context.Background(), db)

	id, err := user.Create{
		FirstName: "Aaa",
		LastName: "Bbb",
	}.Exec(ctx)
	require.NoError(t, err)

	ok, err := user.CreatedToActive{
		ID: id,
	}.Exec(ctx)
	require.NoError(t, err)

	require.True(t, ok)

	user, err := user.SelectByID(ctx, id)
	require.NoError(t, err)

	require.Equal(t, libdb.UserStateActive, user.State)
}

func TestCreatedToActiveWhenNotInCreatedStateReturnsFalse(t *testing.T) {

	db := sqltest.OpenMysql(t, schemaFile)
	ctx := lib.ContextWithDB(context.Background(), db)

	id, err := user.Insert(
		ctx,
		libdb.User{
			FirstName: "Aaa",
			LastName: "Bbb",
			State: libdb.UserStateDeleted,
			CreatedAt: time.Now(),
			DeletedAt: time.Now(),
		},
	)
	require.NoError(t, err)

	ok, err := user.CreatedToActive{
		ID: id,
	}.Exec(ctx)
	require.NoError(t, err)

	require.False(t, ok)

	user, err := user.SelectByID(ctx, id)
	require.NoError(t, err)

	require.Equal(t, libdb.UserStateDeleted, user.State)
}
