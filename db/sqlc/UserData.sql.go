// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: UserData.sql

package db

import (
	"context"
	"time"
)

const createUser = `-- name: CreateUser :one
INSERT INTO user_data(
   username,
   password,
   first_name,
   last_name,
   role,
   phone_number,
   created_at,
   modified_at
   ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8) RETURNING id, username, password, first_name, last_name, role, phone_number, created_at, modified_at
`

type CreateUserParams struct {
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Role        string    `json:"role"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (UserDatum, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Password,
		arg.FirstName,
		arg.LastName,
		arg.Role,
		arg.PhoneNumber,
		arg.CreatedAt,
		arg.ModifiedAt,
	)
	var i UserDatum
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Role,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT id, username, password, first_name, last_name, role, phone_number, created_at, modified_at
FROM user_data
WHERE username = $1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (UserDatum, error) {
	row := q.db.QueryRowContext(ctx, getUserByName, username)
	var i UserDatum
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.FirstName,
		&i.LastName,
		&i.Role,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
