package libdb

import (
	"time"

	"github.com/thecodedproject/dbcrudgen/dbcrudgen"
)

type UserState int64

const (
	UserStateCreated UserState = 0
	UserStateActive UserState = 1
	UserStateDeleted UserState = 2
)

type User struct {
	dbcrudgen.DataModel

	ID int64
	State UserState

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time

	FirstName string
	LastName string
}
