package go_sql_example

import (
	"context"
	"database/sql"

	//"fmt"
)

type Query map[string]any

type Simple struct {
	ID int64
  A int64
  B string
}

func SimpleInsert(
	ctx context.Context,
	db *sql.DB,
	s Simple,
) (int64, error) {

	r, err := db.ExecContext(
		ctx,
		"insert into simple (id, a, b) values (?,?,?)",
		s.ID,
		s.A,
		s.B,
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

// This is probably not a good idea to have as a generic function
// You want to tailor the behaviour when one of the inserts fails to the specific
// use case (e.g. should it be best effort, or maybe use ExecTx to fail all
// if one cant be inserted)
// Probably better use the single insert instead
func SimpleInsertMany(
	ctx context.Context,
	db *sql.DB,
	entries ...Simple,
) (error) {

	for _, s := range entries {
		_, err := db.ExecContext(
			ctx,
			"insert into simple (id, a, b) values (?,?,?)",
			s.ID,
			s.A,
			s.B,
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func SimpleQueryRowByID(
	ctx context.Context,
	db *sql.DB,
	id int64,
) (Simple, error) {

	var a sql.NullInt64
	var b sql.NullString

	var s Simple

	err := db.QueryRowContext(
		ctx,
		"select id, a, b from simple where id = ?",
		id,
	).Scan(
		&s.ID,
		&a,
		&b,
	)
	if err != nil {
		return Simple{}, err
	}

	if a.Valid {
		s.A = a.Int64
	}

	if b.Valid {
		s.B = b.String
	}

	return s, nil
}

func SimpleQuery(
	ctx context.Context,
	db *sql.DB,
	queryParams Query,
) ([]Simple, error) {

	baseQ := "select id, a, b from simple"

	queryVals := make([]any, 0, len(queryParams))
	if len(queryParams) != 0 {

		baseQ += " where "

		i := 0
		for k, v := range(queryParams) {
			baseQ += k + " = ?"
			i++
			if i < len(queryParams) {
				baseQ += " and "
			}

			queryVals = append(queryVals, v)
		}
	}

	r, err := db.QueryContext(
		ctx,
		baseQ,
		queryVals...,
	)
	if err != nil {
		return nil, err
	}


	maxResponses := 1000
	res := make([]Simple, 0, maxResponses)
	for r.Next() {

		if len(res) >= maxResponses {
			return res, nil
		}

		var a sql.NullInt64
		var b sql.NullString

		var s Simple
		err := r.Scan(
			&s.ID,
			&a,
			&b,
		)
		if err != nil {
			return nil, err
		}


		if a.Valid {
			s.A = a.Int64
		}

		if b.Valid {
			s.B = b.String
		}

		res = append(res, s)
	}

	return res, nil
}
