package post

import (
	"context"
	"log"
)

type PostRequest struct {
	One int `json:"one"`
	Two string `json:"two"`
}

type PostResponse struct {
	Success bool `json:"success"`
}

func PostHandler(ctx context.Context, req PostRequest) (PostResponse, error) {

	log.Println("POST /aaa: received request", req)

	return PostResponse{
		Success: true,
	}, nil
}

