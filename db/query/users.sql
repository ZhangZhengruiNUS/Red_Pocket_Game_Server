-- name: GetUserById :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: GetUserByName :one
SELECT * FROM users
WHERE user_name = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2;

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