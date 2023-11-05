-- name: CreatePrize :one
INSERT INTO prizes (
    prize_name,
    pic_path,
    weight
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPrizeByPrizeName :one
SELECT * FROM prizes
WHERE prize_name = $1 LIMIT 1;

-- name: DeletePrize :exec
DELETE FROM prizes
WHERE prize_name = $1;

-- name: UpdatePrize :one
UPDATE prizes
SET pic_path = sqlc.arg(picPath), weight = sqlc.arg(weight)
WHERE prize_name = sqlc.arg(prize_name)
RETURNING *;

-- name: ListRolltable :many
SELECT t1.prize_name, t1.pic_path, t1.weight FROM prizes t1
ORDER BY t1.prize_id
LIMIT sqlc.arg(pageSize)::int
OFFSET ((sqlc.arg(page)::int - 1) * sqlc.arg(pageSize)::int);