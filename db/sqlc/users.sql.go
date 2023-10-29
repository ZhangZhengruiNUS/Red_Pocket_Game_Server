// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  user_name,
  password,
  role_type
) VALUES (
  $1, $2, $3
)
RETURNING user_id, user_name, password, credit, coupon, role_type, create_time
`

type CreateUserParams struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	RoleType int32  `json:"roleType"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser, arg.UserName, arg.Password, arg.RoleType)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.Password,
		&i.Credit,
		&i.Coupon,
		&i.RoleType,
		&i.CreateTime,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userID int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, userID)
	return err
}

const getUserById = `-- name: GetUserById :one
SELECT user_id, user_name, password, credit, coupon, role_type, create_time FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) GetUserById(ctx context.Context, userID int64) (User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, userID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.Password,
		&i.Credit,
		&i.Coupon,
		&i.RoleType,
		&i.CreateTime,
	)
	return i, err
}

const getUserByName = `-- name: GetUserByName :one
SELECT user_id, user_name, password, credit, coupon, role_type, create_time FROM users
WHERE user_name = $1 LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, userName string) (User, error) {
	row := q.queryRow(ctx, q.getUserByNameStmt, getUserByName, userName)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.Password,
		&i.Credit,
		&i.Coupon,
		&i.RoleType,
		&i.CreateTime,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, user_name, password, credit, coupon, role_type, create_time FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.UserID,
			&i.UserName,
			&i.Password,
			&i.Credit,
			&i.Coupon,
			&i.RoleType,
			&i.CreateTime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserCoupon = `-- name: UpdateUserCoupon :one
UPDATE users
SET coupon = coupon + $1
WHERE user_id = $2
RETURNING user_id, user_name, password, credit, coupon, role_type, create_time
`

type UpdateUserCouponParams struct {
	Amount int32 `json:"amount"`
	UserID int64 `json:"userId"`
}

func (q *Queries) UpdateUserCoupon(ctx context.Context, arg UpdateUserCouponParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserCouponStmt, updateUserCoupon, arg.Amount, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.Password,
		&i.Credit,
		&i.Coupon,
		&i.RoleType,
		&i.CreateTime,
	)
	return i, err
}

const updateUserCredit = `-- name: UpdateUserCredit :one
UPDATE users
SET credit = credit + $1
WHERE user_id = $2
RETURNING user_id, user_name, password, credit, coupon, role_type, create_time
`

type UpdateUserCreditParams struct {
	Amount int32 `json:"amount"`
	UserID int64 `json:"userId"`
}

func (q *Queries) UpdateUserCredit(ctx context.Context, arg UpdateUserCreditParams) (User, error) {
	row := q.queryRow(ctx, q.updateUserCreditStmt, updateUserCredit, arg.Amount, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.Password,
		&i.Credit,
		&i.Coupon,
		&i.RoleType,
		&i.CreateTime,
	)
	return i, err
}