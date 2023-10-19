-- name: GetInventory :one
SELECT * FROM inventories
WHERE inventory_id = $1 LIMIT 1;

-- name: GetInventoryByUserIDItemID :one
SELECT * FROM inventories
WHERE user_id = $1
AND item_id = $2 LIMIT 1;

-- name: ListInventories :many
SELECT * FROM inventories
ORDER BY inventory_id
LIMIT $1
OFFSET $2;

-- name: CreateInventory :one
INSERT INTO inventories (
  user_id,
  item_id,
  quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteInventory :exec
DELETE FROM inventories
WHERE inventory_id = $1;

-- name: UpdateInventoryQuantity :one
UPDATE inventories
SET quantity = quantity + sqlc.arg(amount)
WHERE inventory_id = sqlc.arg(inventory_id)
RETURNING *;
