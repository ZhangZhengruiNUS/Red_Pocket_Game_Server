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

const createUserByUserId = `-- name: CreateUserByUserId :one
INSERT INTO users (
  user_id,
  user_name,
  password,
  role_type
) VALUES (
  $1, $2, $3, $4
)
RETURNING user_id, user_name, password, credit, coupon, role_type, create_time
`

type CreateUserByUserIdParams struct {
	UserID   int64  `json:"userId"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	RoleType int32  `json:"roleType"`
}

func (q *Queries) CreateUserByUserId(ctx context.Context, arg CreateUserByUserIdParams) (User, error) {
	row := q.queryRow(ctx, q.createUserByUserIdStmt, createUserByUserId,
		arg.UserID,
		arg.UserName,
		arg.Password,
		arg.RoleType,
	)
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

const getAverageCouponCount = `-- name: GetAverageCouponCount :one
SELECT COALESCE(SUM(COUPON)/NULLIF(COUNT(*), 0), 0)::int AS average_coupon 
FROM USERS LIMIT 1
`

func (q *Queries) GetAverageCouponCount(ctx context.Context) (int32, error) {
	row := q.queryRow(ctx, q.getAverageCouponCountStmt, getAverageCouponCount)
	var average_coupon int32
	err := row.Scan(&average_coupon)
	return average_coupon, err
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

const listCouponByUserID = `-- name: ListCouponByUserID :one
SELECT coupon FROM users
WHERE user_id = $1 LIMIT 1
`

func (q *Queries) ListCouponByUserID(ctx context.Context, userID int64) (int32, error) {
	row := q.queryRow(ctx, q.listCouponByUserIDStmt, listCouponByUserID, userID)
	var coupon int32
	err := row.Scan(&coupon)
	return coupon, err
}

const listUsers = `-- name: ListUsers :many
SELECT user_id, user_name, password, credit, coupon, role_type, create_time FROM users
ORDER BY user_id
LIMIT $2::int
OFFSET (($1::int - 1) * $2::int)
`

type ListUsersParams struct {
	Page     int32 `json:"page"`
	Pagesize int32 `json:"pagesize"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers, arg.Page, arg.Pagesize)
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

const listWarehouse01ByUserID = `-- name: ListWarehouse01ByUserID :one
SELECT credit, coupon FROM users
WHERE user_id = $1 LIMIT 1
`

type ListWarehouse01ByUserIDRow struct {
	Credit int32 `json:"credit"`
	Coupon int32 `json:"coupon"`
}

func (q *Queries) ListWarehouse01ByUserID(ctx context.Context, userID int64) (ListWarehouse01ByUserIDRow, error) {
	row := q.queryRow(ctx, q.listWarehouse01ByUserIDStmt, listWarehouse01ByUserID, userID)
	var i ListWarehouse01ByUserIDRow
	err := row.Scan(&i.Credit, &i.Coupon)
	return i, err
}

const updateCoupon = `-- name: UpdateCoupon :one
UPDATE users
SET coupon = coupon + $1
WHERE user_id = $2
RETURNING coupon
`

type UpdateCouponParams struct {
	Amount int32 `json:"amount"`
	UserID int64 `json:"userId"`
}

func (q *Queries) UpdateCoupon(ctx context.Context, arg UpdateCouponParams) (int32, error) {
	row := q.queryRow(ctx, q.updateCouponStmt, updateCoupon, arg.Amount, arg.UserID)
	var coupon int32
	err := row.Scan(&coupon)
	return coupon, err
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
