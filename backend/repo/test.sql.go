// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: test.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createTest = `-- name: CreateTest :one
INSERT INTO test (name, bio)
VALUES ($1, $2)
RETURNING id, name, bio
`

type CreateTestParams struct {
	Name string
	Bio  pgtype.Text
}

func (q *Queries) CreateTest(ctx context.Context, arg CreateTestParams) (Test, error) {
	row := q.db.QueryRow(ctx, createTest, arg.Name, arg.Bio)
	var i Test
	err := row.Scan(&i.ID, &i.Name, &i.Bio)
	return i, err
}