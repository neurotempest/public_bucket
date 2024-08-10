package types

type User struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

type GETListRequest struct {
	Max int64 `json:"max"`
}

type GETListResponse struct {
	Users []User `json:"users"`
}
