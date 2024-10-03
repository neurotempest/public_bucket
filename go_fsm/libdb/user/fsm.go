package user

import (
	"context"
	"time"

	"github.com/neurotempest/public_bucket/go_fsm/libdb"
)

type Create struct {
	FirstName string
	LastName string
}

type CreatedToActive struct {
	ID int64
}

func (c Create) Exec(ctx context.Context) (int64, error) {

	return Insert(ctx, libdb.User{
		State: libdb.UserStateCreated,

		// TODO Allow time fields to be nil
		CreatedAt: time.Now(),
		DeletedAt: time.Now(),

		FirstName: c.FirstName,
		LastName: c.LastName,
	})
}

func (c CreatedToActive) Exec(ctx context.Context) (bool, error) {

	n, err := Update(
		ctx,
		map[string]any{
			"state": libdb.UserStateActive,
		},
		map[string]any{
			"id": c.ID,
			"state": libdb.UserStateCreated,
		},
	)
	if err != nil {
		return false, err
	}

	return n==1, nil
}
