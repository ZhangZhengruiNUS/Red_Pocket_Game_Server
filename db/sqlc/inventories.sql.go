// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: inventories.sql

package db

import (
	"context"
)

const createInventory = `-- name: CreateInventory :one
INSERT INTO inventories (
  user_id,
  item_id,
  quantity
) VALUES (
  $1, $2, $3
)
RETURNING user_id, item_id, quantity, create_time
`

type CreateInventoryParams struct {
	UserID   int64 `json:"userId"`
	ItemID   int64 `json:"itemId"`
	Quantity int32 `json:"quantity"`
}

func (q *Queries) CreateInventory(ctx context.Context, arg CreateInventoryParams) (Inventory, error) {
	row := q.queryRow(ctx, q.createInventoryStmt, createInventory, arg.UserID, arg.ItemID, arg.Quantity)
	var i Inventory
	err := row.Scan(
		&i.UserID,
		&i.ItemID,
		&i.Quantity,
		&i.CreateTime,
	)
	return i, err
}

const getInventoryByUserIDItemID = `-- name: GetInventoryByUserIDItemID :one
SELECT user_id, item_id, quantity, create_time FROM inventories
WHERE user_id = $1
AND item_id = $2 LIMIT 1
`

type GetInventoryByUserIDItemIDParams struct {
	UserID int64 `json:"userId"`
	ItemID int64 `json:"itemId"`
}

func (q *Queries) GetInventoryByUserIDItemID(ctx context.Context, arg GetInventoryByUserIDItemIDParams) (Inventory, error) {
	row := q.queryRow(ctx, q.getInventoryByUserIDItemIDStmt, getInventoryByUserIDItemID, arg.UserID, arg.ItemID)
	var i Inventory
	err := row.Scan(
		&i.UserID,
		&i.ItemID,
		&i.Quantity,
		&i.CreateTime,
	)
	return i, err
}

const listInventoriesByUserID = `-- name: ListInventoriesByUserID :many
SELECT t1.item_id, COALESCE(t2.item_name, ' '), COALESCE(t2.describe, ' '), t1.quantity, COALESCE(t2.pic_path, ' ')  FROM inventories t1
LEFT JOIN items t2 ON t1.item_id = t2.item_id
WHERE user_id = $3
ORDER BY t1.item_id
LIMIT $1
OFFSET $2
`

type ListInventoriesByUserIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	UserID int64 `json:"userId"`
}

type ListInventoriesByUserIDRow struct {
	ItemID   int64  `json:"itemId"`
	ItemName string `json:"itemName"`
	Describe string `json:"describe"`
	Quantity int32  `json:"quantity"`
	PicPath  string `json:"picPath"`
}

func (q *Queries) ListInventoriesByUserID(ctx context.Context, arg ListInventoriesByUserIDParams) ([]ListInventoriesByUserIDRow, error) {
	rows, err := q.query(ctx, q.listInventoriesByUserIDStmt, listInventoriesByUserID, arg.Limit, arg.Offset, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListInventoriesByUserIDRow{}
	for rows.Next() {
		var i ListInventoriesByUserIDRow
		if err := rows.Scan(
			&i.ItemID,
			&i.ItemName,
			&i.Describe,
			&i.Quantity,
			&i.PicPath,
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

const listWarehouse02ByUserID = `-- name: ListWarehouse02ByUserID :many
SELECT t1.item_id, COALESCE(t2.item_name, ' '), COALESCE(t2.describe, ' '), t1.quantity, COALESCE(t2.price, 0), COALESCE(t2.pic_path, ' ')  FROM inventories t1
LEFT JOIN items t2 ON t1.item_id = t2.item_id
WHERE user_id = $3
ORDER BY t1.item_id
LIMIT $1
OFFSET $2
`

type ListWarehouse02ByUserIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
	UserID int64 `json:"userId"`
}

type ListWarehouse02ByUserIDRow struct {
	ItemID   int64  `json:"itemId"`
	ItemName string `json:"itemName"`
	Describe string `json:"describe"`
	Quantity int32  `json:"quantity"`
	Price    int32  `json:"price"`
	PicPath  string `json:"picPath"`
}

func (q *Queries) ListWarehouse02ByUserID(ctx context.Context, arg ListWarehouse02ByUserIDParams) ([]ListWarehouse02ByUserIDRow, error) {
	rows, err := q.query(ctx, q.listWarehouse02ByUserIDStmt, listWarehouse02ByUserID, arg.Limit, arg.Offset, arg.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListWarehouse02ByUserIDRow{}
	for rows.Next() {
		var i ListWarehouse02ByUserIDRow
		if err := rows.Scan(
			&i.ItemID,
			&i.ItemName,
			&i.Describe,
			&i.Quantity,
			&i.Price,
			&i.PicPath,
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

const updateInventoryQuantity = `-- name: UpdateInventoryQuantity :one
UPDATE inventories
SET quantity = quantity + $1
WHERE user_id = $2
AND item_id = $3
RETURNING user_id, item_id, quantity, create_time
`

type UpdateInventoryQuantityParams struct {
	Amount int32 `json:"amount"`
	UserID int64 `json:"userId"`
	ItemID int64 `json:"itemId"`
}

func (q *Queries) UpdateInventoryQuantity(ctx context.Context, arg UpdateInventoryQuantityParams) (Inventory, error) {
	row := q.queryRow(ctx, q.updateInventoryQuantityStmt, updateInventoryQuantity, arg.Amount, arg.UserID, arg.ItemID)
	var i Inventory
	err := row.Scan(
		&i.UserID,
		&i.ItemID,
		&i.Quantity,
		&i.CreateTime,
	)
	return i, err
}
