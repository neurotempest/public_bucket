package types

struct User {
	FirstName string
	LastName string
}

struct GETListRequest {
	Max int64 `json:"max"`
}

struct GETListResponse {
	Users []User `json:"users"`
}
