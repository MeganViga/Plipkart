// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: UserAddress.sql

package db

import (
	"context"
	"fmt"
	"time"
)

const createUserAddress = `-- name: CreateUserAddress :one
INSERT INTO user_address(
   user_id,
   address_line1,
   address_line2,
   city,
   postal_code,
   country,
   phone_number,
   created_at,
   modified_at
   ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING id, user_id, address_line1, address_line2, city, postal_code, country, phone_number, created_at, modified_at
`

type CreateUserAddressParams struct {
	UserID       int64     `json:"user_id"`
	AddressLine1 string    `json:"address_line1"`
	AddressLine2 string    `json:"address_line2"`
	City         string    `json:"city"`
	PostalCode   string    `json:"postal_code"`
	Country      string    `json:"country"`
	PhoneNumber  string    `json:"phone_number"`
	CreatedAt    time.Time `json:"created_at"`
	ModifiedAt   time.Time `json:"modified_at"`
}

func (q *Queries) CreateUserAddress(ctx context.Context, arg CreateUserAddressParams) (UserAddress, error) {
	row := q.db.QueryRowContext(ctx, createUserAddress,
		arg.UserID,
		arg.AddressLine1,
		arg.AddressLine2,
		arg.City,
		arg.PostalCode,
		arg.Country,
		arg.PhoneNumber,
		arg.CreatedAt,
		arg.ModifiedAt,
	)
	var i UserAddress
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deleteUserAddress = `-- name: DeleteUserAddress :one
DELETE 
FROM user_address
WHERE id = $1 RETURNING id, user_id, address_line1, address_line2, city, postal_code, country, phone_number, created_at, modified_at
`

func (q *Queries) DeleteUserAddress(ctx context.Context, id int64) (UserAddress, error) {
	row := q.db.QueryRowContext(ctx, deleteUserAddress, id)
	var i UserAddress
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.AddressLine1,
		&i.AddressLine2,
		&i.City,
		&i.PostalCode,
		&i.Country,
		&i.PhoneNumber,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const listUserAddresses = `-- name: ListUserAddresses :many
SELECT id, user_id, address_line1, address_line2, city, postal_code, country, phone_number, created_at, modified_at 
FROM user_address
WHERE user_id = $1
`

func (q *Queries) ListUserAddresses(ctx context.Context, userID int64) ([]UserAddress, error) {
	rows, err := q.db.QueryContext(ctx, listUserAddresses, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []UserAddress
	for rows.Next() {
		var i UserAddress
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.AddressLine1,
			&i.AddressLine2,
			&i.City,
			&i.PostalCode,
			&i.Country,
			&i.PhoneNumber,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		fmt.Println("I",i)
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	fmt.Println("Items",items)
	return items, nil
}
