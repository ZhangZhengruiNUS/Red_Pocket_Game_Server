-- name: GetUserById :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT sqlc.arg(pageSize)::int
OFFSET ((sqlc.arg(page)::int - 1) * sqlc.arg(pageSize)::int);

-- name: CreateUser :one
INSERT INTO users (
  user_name,
  password,
  role_type
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1;

-- name: UpdateUserCredit :one
UPDATE users
SET credit = credit + sqlc.arg(amount)
WHERE user_id = sqlc.arg(user_id)
RETURNING *;

-- name: UpdateUserCoupon :one
UPDATE users
SET coupon = coupon + sqlc.arg(amount)
WHERE user_id = sqlc.arg(user_id)
RETURNING *;

-- name: GetAverageCouponCount :one
SELECT COALESCE(SUM(COUPON)/NULLIF(COUNT(*), 0), 0)::int AS average_coupon 
FROM USERS LIMIT 1;
