package user


import (
	usertypes "github.com/thecodedproject/publicbucket/golang_distributed_handlers/handlers/user/types"
)


func GETList(ctx context.Context, req usertypes.GETListRequest) (usertypes.GETListResponse, error) {

	log.Println("GET users/list: received request:", req)

	return usertypes.GETUserResponse{
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
	}
}
