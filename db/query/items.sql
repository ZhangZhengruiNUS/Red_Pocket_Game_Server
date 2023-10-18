-- name: GetItem :one
SELECT * FROM items
WHERE item_id = $1 LIMIT 1;

-- name: ListItems :many
SELECT * FROM items
ORDER BY item_id
LIMIT $1
OFFSET $2;

-- name: CreateItem :one
INSERT INTO items (
  item_name,
  describe,
  pic_path,
  price,
  creator_id
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE item_id = $1;