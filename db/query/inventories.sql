-- name: GetInventory :one
SELECT * FROM inventories
WHERE id = $1 LIMIT 1;

-- name: ListInventories :many
SELECT * FROM inventories
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateInventory :one
INSERT INTO inventories (
  account_id,
  item_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: DeleteInventory :exec
DELETE FROM inventories
WHERE id = $1;