-- name: ListWarehouse01ByUserID :one
SELECT credit, coupon FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListWarehouse02ByUserID :many
SELECT t1.item_id, COALESCE(t2.item_name, ' '), COALESCE(t2.describe, ' '), t1.quantity, COALESCE(t2.price, 0), COALESCE(t2.pic_path, ' ')  FROM inventories t1
LEFT JOIN items t2 ON t1.item_id = t2.item_id
WHERE user_id = $3
ORDER BY t1.item_id
LIMIT $1
OFFSET $2;

-- name: ListRolltable :many
SELECT t1.prize_name, t1.pic_path, t1.weight FROM prizes t1
ORDER BY t1.prize_id
LIMIT $1
OFFSET $2;

-- name: ListCouponByUserID :one
SELECT coupon FROM users
WHERE user_id = $1 LIMIT 1;

-- name: UpdateCoupon :one
UPDATE users
SET coupon = coupon + sqlc.arg(amount)
WHERE user_id = sqlc.arg(user_id)
RETURNING coupon;