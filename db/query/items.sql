-- name: GetItem :one
SELECT * FROM items
WHERE id = $1 LIMIT 1;

-- name: ListItems :many
SELECT * FROM items
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateItem :one
INSERT INTO items (
  item_name,
  item_price,
  creator_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateItem :exec
UPDATE items
  set item_name = $2,
  item_price = $3
WHERE id = $1;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;