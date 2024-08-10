package get

import (
	"context"
	"log"
)

type GetRequest struct {
	One int `json:"one"`
	Two string `json:"two"`
}

type GetResponse struct {
	Items []item `json:"items"`
}

type item struct {
	OneVal string `json:"one_val"`
	TwoVal int `json:"two_val"`
}

func GetHandler(ctx context.Context, req GetRequest) (GetResponse, error) {

	log.Println("GET /ccc/ccc: received request", req)

	return GetResponse{
		Items: []item{
			{
				OneVal: "one",
				TwoVal: 1,
			},
			{
				OneVal: "two",
				TwoVal: 2,
			},
		},
	}, nil
}

