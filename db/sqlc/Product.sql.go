// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: Product.sql

package db

import (
	"context"
	"database/sql"
	"time"
)

const addProduct = `-- name: AddProduct :one
INSERT INTO product(
    product_name,
    price,
    created_at,
    modified_at,
    deleted_at
 ) VALUES($1,$2,$3,$4,$5) RETURNING id, product_name, price, created_at, modified_at, deleted_at
`

type AddProductParams struct {
	ProductName string       `json:"product_name"`
	Price       int64        `json:"price"`
	CreatedAt   time.Time    `json:"created_at"`
	ModifiedAt  time.Time    `json:"modified_at"`
	DeletedAt   sql.NullTime `json:"deleted_at"`
}

func (q *Queries) AddProduct(ctx context.Context, arg AddProductParams) (Product, error) {
	row := q.db.QueryRowContext(ctx, addProduct,
		arg.ProductName,
		arg.Price,
		arg.CreatedAt,
		arg.ModifiedAt,
		arg.DeletedAt,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Price,
		&i.CreatedAt,
		&i.ModifiedAt,
		&i.DeletedAt,
	)
	return i, err
}

const addProductCategory = `-- name: AddProductCategory :one
 INSERT INTO product_category(
    product_id,
    color,
    quantity
 ) VALUES($1,$2,$3) RETURNING product_id, color, quantity
`

type AddProductCategoryParams struct {
	ProductID int64  `json:"product_id"`
	Color     string `json:"color"`
	Quantity  int32  `json:"quantity"`
}

func (q *Queries) AddProductCategory(ctx context.Context, arg AddProductCategoryParams) (ProductCategory, error) {
	row := q.db.QueryRowContext(ctx, addProductCategory, arg.ProductID, arg.Color, arg.Quantity)
	var i ProductCategory
	err := row.Scan(&i.ProductID, &i.Color, &i.Quantity)
	return i, err
}

const getProductByColor = `-- name: GetProductByColor :many
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id
 WHERE product_id = $1 AND color = $2
`

type GetProductByColorParams struct {
	ProductID int64  `json:"product_id"`
	Color     string `json:"color"`
}

type GetProductByColorRow struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Color       string `json:"color"`
	Quantity    int32  `json:"quantity"`
	Price       int64  `json:"price"`
}
func (g GetProductByColorRow)Get()(int64,string,string,int32,int64){
	return g.ID,g.ProductName,g.Color,g.Quantity,g.Price
}
func (q *Queries) GetProductByColor(ctx context.Context, arg GetProductByColorParams) ([]GetProductByColorRow, error) {
	rows, err := q.db.QueryContext(ctx, getProductByColor, arg.ProductID, arg.Color)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductByColorRow
	for rows.Next() {
		var i GetProductByColorRow
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.Color,
			&i.Quantity,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductById = `-- name: GetProductById :one
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id
 WHERE id = $1
`

type GetProductByIdRow struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Color       string `json:"color"`
	Quantity    int32  `json:"quantity"`
	Price       int64  `json:"price"`
}

func (q *Queries) GetProductById(ctx context.Context, id int64) (GetProductByIdRow, error) {
	row := q.db.QueryRowContext(ctx, getProductById, id)
	var i GetProductByIdRow
	err := row.Scan(
		&i.ID,
		&i.ProductName,
		&i.Color,
		&i.Quantity,
		&i.Price,
	)
	return i, err
}

const getProductByName = `-- name: GetProductByName :many
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id
 WHERE product_name = $1
`

type GetProductByNameRow struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Color       string `json:"color"`
	Quantity    int32  `json:"quantity"`
	Price       int64  `json:"price"`
}

func (q *Queries) GetProductByName(ctx context.Context, productName string) ([]GetProductByNameRow, error) {
	rows, err := q.db.QueryContext(ctx, getProductByName, productName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductByNameRow
	for rows.Next() {
		var i GetProductByNameRow
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.Color,
			&i.Quantity,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProducts = `-- name: GetProducts :many
 SELECT id,product_name,color,quantity,price 
 FROM product 
 INNER JOIN product_category 
 ON id = product_id
`

type GetProductsRow struct {
	ID          int64  `json:"id"`
	ProductName string `json:"product_name"`
	Color       string `json:"color"`
	Quantity    int32  `json:"quantity"`
	Price       int64  `json:"price"`
}
func (g GetProductsRow)Get()(int64,string,string,int32,int64){
	return g.ID,g.ProductName,g.Color,g.Quantity,g.Price
}
func (q *Queries) GetProducts(ctx context.Context) ([]GetProductsRow, error) {
	rows, err := q.db.QueryContext(ctx, getProducts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsRow
	for rows.Next() {
		var i GetProductsRow
		if err := rows.Scan(
			&i.ID,
			&i.ProductName,
			&i.Color,
			&i.Quantity,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
