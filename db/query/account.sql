-- name: GetAccountById :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountByName :one
SELECT * FROM accounts
WHERE username = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateAccount :one
INSERT INTO accounts (
  username,
  password,
  role_type
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateAccountPwd :exec
UPDATE accounts
  set password = $2
WHERE id = $1;

-- name: DeleteAuthor :exec
DELETE FROM accounts
WHERE id = $1;