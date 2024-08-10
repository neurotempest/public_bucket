package users

import (
	"context"

	usertypes "github.com/neurotempest/public_bucket/golang_distributed_handlers/handlers/users/types"
)


func GETList(ctx context.Context, req usertypes.GETListRequest) (usertypes.GETListResponse, error) {

	return usertypes.GETListResponse{
		Users: []usertypes.User{
			{
				FirstName: "A",
				LastName: "Aa",
			},
			{
				FirstName: "B",
				LastName: "Bb",
			},
		},
	}, nil
}
