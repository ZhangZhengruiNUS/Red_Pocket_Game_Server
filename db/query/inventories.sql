-- name: GetInventoryByUserIDItemID :one
SELECT * FROM inventories
WHERE user_id = $1
AND item_id = $2 LIMIT 1;

-- name: ListInventoriesByUserID :many
SELECT t1.item_id, COALESCE(t2.item_name, ' '), COALESCE(t2.describe, ' '), t1.quantity, COALESCE(t2.pic_path, ' ')  FROM inventories t1
LEFT JOIN items t2 ON t1.item_id = t2.item_id
WHERE user_id = $3
ORDER BY t1.item_id
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

-- name: UpdateInventoryQuantity :one
UPDATE inventories
SET quantity = quantity + sqlc.arg(amount)
WHERE user_id = sqlc.arg(user_id)
AND item_id = sqlc.arg(item_id)
RETURNING *;
