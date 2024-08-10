package list_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/users/list"
)

func TestGetHandler(t *testing.T) {
	resp, err := list.GetHandler(
		context.Background(),
		list.GetRequest{},
	)

	require.NoError(t, err)
	require.Equal(t, 2, len(resp.Users))
}

