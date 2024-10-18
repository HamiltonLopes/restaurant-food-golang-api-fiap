// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: products.sql

package fiapRestaurantDb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const getProductById = `-- name: GetProductById :many
SELECT 
    p.id, 
    p.name, 
    p.description, 
    p.price, 
    p.created_at,
    p.updated_at,
    c.id AS category_id, 
    c.name AS category_name,
    c.handle AS category_handle,
    c.created_at AS category_created_at,
    c.updated_at AS category_updated_at,
    pi.id AS image_id, 
    pi.image,
    pi.created_at AS image_created_at,
    pi.updated_at AS image_updated_at
FROM products p
LEFT JOIN categories c ON p.category_id = c.id AND c.deleted_at IS NULL
LEFT JOIN products_images pi ON p.id = pi.product_id AND pi.deleted_at IS NULL
WHERE p.id = $1 AND p.deleted_at IS NULL
`

type GetProductByIdRow struct {
	ID                int32
	Name              string
	Description       string
	Price             pgtype.Numeric
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
	CategoryID        pgtype.Int4
	CategoryName      pgtype.Text
	CategoryHandle    pgtype.Text
	CategoryCreatedAt pgtype.Timestamptz
	CategoryUpdatedAt pgtype.Timestamptz
	ImageID           pgtype.Int4
	Image             pgtype.Text
	ImageCreatedAt    pgtype.Timestamptz
	ImageUpdatedAt    pgtype.Timestamptz
}

func (q *Queries) GetProductById(ctx context.Context, id int32) ([]GetProductByIdRow, error) {
	rows, err := q.db.Query(ctx, getProductById, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductByIdRow
	for rows.Next() {
		var i GetProductByIdRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CategoryID,
			&i.CategoryName,
			&i.CategoryHandle,
			&i.CategoryCreatedAt,
			&i.CategoryUpdatedAt,
			&i.ImageID,
			&i.Image,
			&i.ImageCreatedAt,
			&i.ImageUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getProductsByIds = `-- name: GetProductsByIds :many
SELECT 
    p.id, 
    p.name, 
    p.description, 
    p.price, 
    p.created_at,
    p.updated_at,
    c.id AS category_id, 
    c.name AS category_name,
    c.created_at AS category_created_at,
    c.updated_at AS category_updated_at,
    pi.id AS image_id, 
    pi.image,
    pi.created_at AS image_created_at,
    pi.updated_at AS image_updated_at
FROM products p
LEFT JOIN categories c ON p.category_id = c.id
LEFT JOIN products_images pi ON p.id = pi.product_id
WHERE p.deleted_at IS NULL AND p.id = ANY($1::int[])
`

type GetProductsByIdsRow struct {
	ID                int32
	Name              string
	Description       string
	Price             pgtype.Numeric
	CreatedAt         pgtype.Timestamptz
	UpdatedAt         pgtype.Timestamptz
	CategoryID        pgtype.Int4
	CategoryName      pgtype.Text
	CategoryCreatedAt pgtype.Timestamptz
	CategoryUpdatedAt pgtype.Timestamptz
	ImageID           pgtype.Int4
	Image             pgtype.Text
	ImageCreatedAt    pgtype.Timestamptz
	ImageUpdatedAt    pgtype.Timestamptz
}

func (q *Queries) GetProductsByIds(ctx context.Context, ids []int32) ([]GetProductsByIdsRow, error) {
	rows, err := q.db.Query(ctx, getProductsByIds, ids)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProductsByIdsRow
	for rows.Next() {
		var i GetProductsByIdsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Price,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CategoryID,
			&i.CategoryName,
			&i.CategoryCreatedAt,
			&i.CategoryUpdatedAt,
			&i.ImageID,
			&i.Image,
			&i.ImageCreatedAt,
			&i.ImageUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :one
UPDATE products
SET name = $2, description = $3, price = $4, category_id = $5
WHERE id = $1
RETURNING id, name, description, price, category_id, created_at, updated_at, deleted_at
`

type UpdateProductParams struct {
	ID          int32
	Name        string
	Description string
	Price       pgtype.Numeric
	CategoryID  int32
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, updateProduct,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.CategoryID,
	)
	var i Product
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Price,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
