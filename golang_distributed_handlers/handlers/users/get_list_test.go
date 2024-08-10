package users_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/users"
	"github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/users/types"
)

func TestGetList(t *testing.T) {

	resp, err := users.GETList(
		context.Background(),
		types.GETListRequest{},
	)

	require.NoError(t, err)
	require.Equal(t, 2, len(resp.Users))
}
