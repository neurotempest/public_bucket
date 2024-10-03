package user

import (
	context "context"
	errors "errors"
	fmt "fmt"
	libdb "github.com/neurotempest/public_bucket/go_fsm/libdb"
	lib "github.com/thecodedproject/dbcrudgen/lib"
	time "github.com/thecodedproject/gotest/time"
)

func Insert(
	ctx context.Context,
	d libdb.User,
) (int64, error) {

	db, err := lib.DBFromContext(ctx)
	if err != nil {
		return 0, err
	}

	r, err := db.ExecContext(
		ctx,
		"insert into user set state=?, created_at=?, updated_at=?, deleted_at=?, first_name=?, last_name=?",
		d.State,
		d.CreatedAt,
		time.Now(),
		d.DeletedAt,
		d.FirstName,
		d.LastName,
	)
	if err != nil {
		return 0, err
	}

	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func SelectByID(
	ctx context.Context,
	id int64,
) (libdb.User, error) {

	r, err := Select(
		ctx,
		map[string]any{
			"id": id,
		},
	)
	if err != nil {
		return libdb.User{}, err
	}

	if len(r) == 0 {
		return libdb.User{}, errors.New("SelectByID: id not found - " + fmt.Sprint(id))
	}

	if len(r) > 1 {
		return libdb.User{}, errors.New("found more than one entry with id")
	}

	return r[0], nil
}

func Select(
	ctx context.Context,
	queryParams map[string]any,
) ([]libdb.User, error) {

	db, err := lib.DBFromContext(ctx)
	if err != nil {
		return nil, err
	}

	q := "select id, state, created_at, updated_at, deleted_at, first_name, last_name from user"

	if len(queryParams) > 0 {
		q += " where "
	}

	queryVals := make([]any, 0, len(queryParams))
	i := 0
	for k, v := range queryParams {
		if !modelContainsField(k) {
			return nil, errors.New("Select: no such field to query - " + k)
		}

		q += k + "=?"
		i++
		if i < len(queryParams) {
			q += " and "
		}
		queryVals = append(queryVals, v)
	}

	r, err := db.QueryContext(
		ctx,
		q,
		queryVals...,
	)
	if err != nil {
		return nil, err
	}

	// TODO: make this a configurable param
	maxResponses := 1000
	res := make([]libdb.User, 0, maxResponses)
	for r.Next() {

		if len(res) >= maxResponses {
			return nil, errors.New("select query exceeded max responses")
		}

		var d libdb.User
		err := r.Scan(
			&d.ID,
			&d.State,
			&d.CreatedAt,
			&d.UpdatedAt,
			&d.DeletedAt,
			&d.FirstName,
			&d.LastName,
		)
		if err != nil {
			return nil, err
		}

		res = append(res, d)
	}

	return res, nil
}

func Update(
	ctx context.Context,
	updates map[string]any,
	queryParams map[string]any,
) (int64, error) {

	db, err := lib.DBFromContext(ctx)
	if err != nil {
		return 0, err
	}

	if len(updates) == 0 {
		return 0, nil
	}

	query := "update user set "
	queryArgs := make([]any, 0, len(updates) + len(queryParams))
	i := 0
	for k, v := range updates {
		if !modelContainsField(k) {
			return 0, errors.New("Update: no such field to update - " + k)
		}

		query += k + "=?"
		i++
		if i < len(updates) {
			query += ", "
		}

		queryArgs = append(queryArgs, v)
	}

	if len(queryParams) > 0 {
		query += " where "
	}
	i = 0
	for k, v := range queryParams {
		if !modelContainsField(k) {
			return 0, errors.New("Update: no such field to query - " + k)
		}

		query += k + "=?"
		i++
		if i < len(queryParams) {
			query += " and "
		}

		queryArgs = append(queryArgs, v)
	}

	r, err := db.ExecContext(
		ctx,
		query,
		queryArgs...,
	)
	if err != nil {
		return 0, err
	}

	count, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func UpdateByID(
	ctx context.Context,
	id int64,
	updates map[string]any,
) error {

	if len(updates) == 0 {
		return nil
	}

	n, err := Update(
		ctx,
		updates,
		map[string]any{
			"id": id,
		},
	)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("UpdateByID: no such ID")
	}

	return nil
}

func Delete(
	ctx context.Context,
	queryParams map[string]any,
) (int64, error) {

	db, err := lib.DBFromContext(ctx)
	if err != nil {
		return 0, err
	}

	query := "delete from user"

	if len(queryParams) > 0 {
		query += " where "
	}
	i := 0
	queryArgs := make([]any, 0, len(queryParams))
	for k, v := range queryParams {
		if !modelContainsField(k) {
			return 0, errors.New("Delete: no such field to query - " + k)
		}

		query += k + "=?"
		i++
		if i < len(queryParams) {
			query += " and "
		}

		queryArgs = append(queryArgs, v)
	}

	r, err := db.ExecContext(
		ctx,
		query,
		queryArgs...,
	)
	if err != nil {
		return 0, err
	}

	count, err := r.RowsAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func DeleteByID(
	ctx context.Context,
	id int64,
) error {

	n, err := Delete(
		ctx,
		map[string]any{
			"id": id,
		},
	)
	if err != nil {
		return err
	}

	if n == 0 {
		return errors.New("DeleteByID: no such ID")
	}

	return nil
}

func modelContainsField(field string) bool {

	modelFields := map[string]bool{
		"id": true,
		"state": true,
		"created_at": true,
		"updated_at": true,
		"deleted_at": true,
		"first_name": true,
		"last_name": true,
	}

	return modelFields[field]
}

