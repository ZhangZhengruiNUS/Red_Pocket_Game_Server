-- name: GetInventory :one
SELECT * FROM inventories
WHERE inventory_id = $1 LIMIT 1;

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