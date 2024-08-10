package list

import (
	"context"
	"log"
)

type GetRequest struct {
	// add some generate generic type to signify this as a request??
	// httpgen.GetRequest

	Max int `json:"max"`
}

type GetResponse struct {
	Users []user `json:"users"`
}

type user struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

func GetHandler(ctx context.Context, req GetRequest) (GetResponse, error) {

	log.Println("GetHandler: received request", req)

	return GetResponse{
		Users: []user{
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
