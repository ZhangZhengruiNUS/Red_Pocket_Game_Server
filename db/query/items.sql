-- name: GetItem :one
SELECT * FROM items
WHERE item_id = $1 LIMIT 1;

-- name: GetItemByItemName :one
SELECT * FROM items
WHERE item_name = $1 LIMIT 1;

-- name: ListItems :many
SELECT * FROM items
ORDER BY item_id
LIMIT sqlc.arg(pageSize)::int
OFFSET ((sqlc.arg(page)::int - 1) * sqlc.arg(pageSize)::int);

-- name: CreateItem :one
INSERT INTO items (
  item_name,
  describe,
  pic_path,
  price,
  creator_id,
  create_time
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: DeleteItem :exec
DELETE FROM items
WHERE item_id = $1;

-- name: UpdateItem :one
UPDATE items
SET item_name = $2,
 describe = $3, 
 price = $4, 
 pic_path = $5, 
 reviser_id = sqlc.arg(reviserId)::bigint,
 revise_time = sqlc.arg(reviseTime)::timestamp
WHERE item_id = $1
RETURNING *;