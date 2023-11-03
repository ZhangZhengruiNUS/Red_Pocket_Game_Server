// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createDiffLvStmt, err = db.PrepareContext(ctx, createDiffLv); err != nil {
		return nil, fmt.Errorf("error preparing query CreateDiffLv: %w", err)
	}
	if q.createInventoryStmt, err = db.PrepareContext(ctx, createInventory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateInventory: %w", err)
	}
	if q.createItemStmt, err = db.PrepareContext(ctx, createItem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateItem: %w", err)
	}
	if q.createPrizeStmt, err = db.PrepareContext(ctx, createPrize); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePrize: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteDiffLvStmt, err = db.PrepareContext(ctx, deleteDiffLv); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteDiffLv: %w", err)
	}
	if q.deleteItemStmt, err = db.PrepareContext(ctx, deleteItem); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteItem: %w", err)
	}
	if q.deletePrizeStmt, err = db.PrepareContext(ctx, deletePrize); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePrize: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getInventoryByUserIDItemIDStmt, err = db.PrepareContext(ctx, getInventoryByUserIDItemID); err != nil {
		return nil, fmt.Errorf("error preparing query GetInventoryByUserIDItemID: %w", err)
	}
	if q.getItemStmt, err = db.PrepareContext(ctx, getItem); err != nil {
		return nil, fmt.Errorf("error preparing query GetItem: %w", err)
	}
	if q.getItemByItemNameStmt, err = db.PrepareContext(ctx, getItemByItemName); err != nil {
		return nil, fmt.Errorf("error preparing query GetItemByItemName: %w", err)
	}
	if q.getPrizeByPrizeNameStmt, err = db.PrepareContext(ctx, getPrizeByPrizeName); err != nil {
		return nil, fmt.Errorf("error preparing query GetPrizeByPrizeName: %w", err)
	}
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getUserByNameStmt, err = db.PrepareContext(ctx, getUserByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByName: %w", err)
	}
	if q.listCouponByUserIDStmt, err = db.PrepareContext(ctx, listCouponByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query ListCouponByUserID: %w", err)
	}
	if q.listGameDiffSetsStmt, err = db.PrepareContext(ctx, listGameDiffSets); err != nil {
		return nil, fmt.Errorf("error preparing query ListGameDiffSets: %w", err)
	}
	if q.listGameDiffSetsByDiffLvStmt, err = db.PrepareContext(ctx, listGameDiffSetsByDiffLv); err != nil {
		return nil, fmt.Errorf("error preparing query ListGameDiffSetsByDiffLv: %w", err)
	}
	if q.listInventoriesByUserIDStmt, err = db.PrepareContext(ctx, listInventoriesByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query ListInventoriesByUserID: %w", err)
	}
	if q.listItemsStmt, err = db.PrepareContext(ctx, listItems); err != nil {
		return nil, fmt.Errorf("error preparing query ListItems: %w", err)
	}
	if q.listRolltableStmt, err = db.PrepareContext(ctx, listRolltable); err != nil {
		return nil, fmt.Errorf("error preparing query ListRolltable: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.listWarehouse01ByUserIDStmt, err = db.PrepareContext(ctx, listWarehouse01ByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query ListWarehouse01ByUserID: %w", err)
	}
	if q.listWarehouse02ByUserIDStmt, err = db.PrepareContext(ctx, listWarehouse02ByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query ListWarehouse02ByUserID: %w", err)
	}
	if q.updateCouponStmt, err = db.PrepareContext(ctx, updateCoupon); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCoupon: %w", err)
	}
	if q.updateDiffLvStmt, err = db.PrepareContext(ctx, updateDiffLv); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateDiffLv: %w", err)
	}
	if q.updateInventoryQuantityStmt, err = db.PrepareContext(ctx, updateInventoryQuantity); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateInventoryQuantity: %w", err)
	}
	if q.updateItemStmt, err = db.PrepareContext(ctx, updateItem); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateItem: %w", err)
	}
	if q.updatePrizeStmt, err = db.PrepareContext(ctx, updatePrize); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePrize: %w", err)
	}
	if q.updateUserCouponStmt, err = db.PrepareContext(ctx, updateUserCoupon); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserCoupon: %w", err)
	}
	if q.updateUserCreditStmt, err = db.PrepareContext(ctx, updateUserCredit); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUserCredit: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createDiffLvStmt != nil {
		if cerr := q.createDiffLvStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createDiffLvStmt: %w", cerr)
		}
	}
	if q.createInventoryStmt != nil {
		if cerr := q.createInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createInventoryStmt: %w", cerr)
		}
	}
	if q.createItemStmt != nil {
		if cerr := q.createItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createItemStmt: %w", cerr)
		}
	}
	if q.createPrizeStmt != nil {
		if cerr := q.createPrizeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPrizeStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteDiffLvStmt != nil {
		if cerr := q.deleteDiffLvStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteDiffLvStmt: %w", cerr)
		}
	}
	if q.deleteItemStmt != nil {
		if cerr := q.deleteItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteItemStmt: %w", cerr)
		}
	}
	if q.deletePrizeStmt != nil {
		if cerr := q.deletePrizeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePrizeStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.getInventoryByUserIDItemIDStmt != nil {
		if cerr := q.getInventoryByUserIDItemIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInventoryByUserIDItemIDStmt: %w", cerr)
		}
	}
	if q.getItemStmt != nil {
		if cerr := q.getItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getItemStmt: %w", cerr)
		}
	}
	if q.getItemByItemNameStmt != nil {
		if cerr := q.getItemByItemNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getItemByItemNameStmt: %w", cerr)
		}
	}
	if q.getPrizeByPrizeNameStmt != nil {
		if cerr := q.getPrizeByPrizeNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPrizeByPrizeNameStmt: %w", cerr)
		}
	}
	if q.getUserByIdStmt != nil {
		if cerr := q.getUserByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByIdStmt: %w", cerr)
		}
	}
	if q.getUserByNameStmt != nil {
		if cerr := q.getUserByNameStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserByNameStmt: %w", cerr)
		}
	}
	if q.listCouponByUserIDStmt != nil {
		if cerr := q.listCouponByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listCouponByUserIDStmt: %w", cerr)
		}
	}
	if q.listGameDiffSetsStmt != nil {
		if cerr := q.listGameDiffSetsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listGameDiffSetsStmt: %w", cerr)
		}
	}
	if q.listGameDiffSetsByDiffLvStmt != nil {
		if cerr := q.listGameDiffSetsByDiffLvStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listGameDiffSetsByDiffLvStmt: %w", cerr)
		}
	}
	if q.listInventoriesByUserIDStmt != nil {
		if cerr := q.listInventoriesByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInventoriesByUserIDStmt: %w", cerr)
		}
	}
	if q.listItemsStmt != nil {
		if cerr := q.listItemsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listItemsStmt: %w", cerr)
		}
	}
	if q.listRolltableStmt != nil {
		if cerr := q.listRolltableStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listRolltableStmt: %w", cerr)
		}
	}
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.listWarehouse01ByUserIDStmt != nil {
		if cerr := q.listWarehouse01ByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listWarehouse01ByUserIDStmt: %w", cerr)
		}
	}
	if q.listWarehouse02ByUserIDStmt != nil {
		if cerr := q.listWarehouse02ByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listWarehouse02ByUserIDStmt: %w", cerr)
		}
	}
	if q.updateCouponStmt != nil {
		if cerr := q.updateCouponStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCouponStmt: %w", cerr)
		}
	}
	if q.updateDiffLvStmt != nil {
		if cerr := q.updateDiffLvStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateDiffLvStmt: %w", cerr)
		}
	}
	if q.updateInventoryQuantityStmt != nil {
		if cerr := q.updateInventoryQuantityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateInventoryQuantityStmt: %w", cerr)
		}
	}
	if q.updateItemStmt != nil {
		if cerr := q.updateItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateItemStmt: %w", cerr)
		}
	}
	if q.updatePrizeStmt != nil {
		if cerr := q.updatePrizeStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePrizeStmt: %w", cerr)
		}
	}
	if q.updateUserCouponStmt != nil {
		if cerr := q.updateUserCouponStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserCouponStmt: %w", cerr)
		}
	}
	if q.updateUserCreditStmt != nil {
		if cerr := q.updateUserCreditStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserCreditStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                             DBTX
	tx                             *sql.Tx
	createDiffLvStmt               *sql.Stmt
	createInventoryStmt            *sql.Stmt
	createItemStmt                 *sql.Stmt
	createPrizeStmt                *sql.Stmt
	createUserStmt                 *sql.Stmt
	deleteDiffLvStmt               *sql.Stmt
	deleteItemStmt                 *sql.Stmt
	deletePrizeStmt                *sql.Stmt
	deleteUserStmt                 *sql.Stmt
	getInventoryByUserIDItemIDStmt *sql.Stmt
	getItemStmt                    *sql.Stmt
	getItemByItemNameStmt          *sql.Stmt
	getPrizeByPrizeNameStmt        *sql.Stmt
	getUserByIdStmt                *sql.Stmt
	getUserByNameStmt              *sql.Stmt
	listCouponByUserIDStmt         *sql.Stmt
	listGameDiffSetsStmt           *sql.Stmt
	listGameDiffSetsByDiffLvStmt   *sql.Stmt
	listInventoriesByUserIDStmt    *sql.Stmt
	listItemsStmt                  *sql.Stmt
	listRolltableStmt              *sql.Stmt
	listUsersStmt                  *sql.Stmt
	listWarehouse01ByUserIDStmt    *sql.Stmt
	listWarehouse02ByUserIDStmt    *sql.Stmt
	updateCouponStmt               *sql.Stmt
	updateDiffLvStmt               *sql.Stmt
	updateInventoryQuantityStmt    *sql.Stmt
	updateItemStmt                 *sql.Stmt
	updatePrizeStmt                *sql.Stmt
	updateUserCouponStmt           *sql.Stmt
	updateUserCreditStmt           *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                             tx,
		tx:                             tx,
		createDiffLvStmt:               q.createDiffLvStmt,
		createInventoryStmt:            q.createInventoryStmt,
		createItemStmt:                 q.createItemStmt,
		createPrizeStmt:                q.createPrizeStmt,
		createUserStmt:                 q.createUserStmt,
		deleteDiffLvStmt:               q.deleteDiffLvStmt,
		deleteItemStmt:                 q.deleteItemStmt,
		deletePrizeStmt:                q.deletePrizeStmt,
		deleteUserStmt:                 q.deleteUserStmt,
		getInventoryByUserIDItemIDStmt: q.getInventoryByUserIDItemIDStmt,
		getItemStmt:                    q.getItemStmt,
		getItemByItemNameStmt:          q.getItemByItemNameStmt,
		getPrizeByPrizeNameStmt:        q.getPrizeByPrizeNameStmt,
		getUserByIdStmt:                q.getUserByIdStmt,
		getUserByNameStmt:              q.getUserByNameStmt,
		listCouponByUserIDStmt:         q.listCouponByUserIDStmt,
		listGameDiffSetsStmt:           q.listGameDiffSetsStmt,
		listGameDiffSetsByDiffLvStmt:   q.listGameDiffSetsByDiffLvStmt,
		listInventoriesByUserIDStmt:    q.listInventoriesByUserIDStmt,
		listItemsStmt:                  q.listItemsStmt,
		listRolltableStmt:              q.listRolltableStmt,
		listUsersStmt:                  q.listUsersStmt,
		listWarehouse01ByUserIDStmt:    q.listWarehouse01ByUserIDStmt,
		listWarehouse02ByUserIDStmt:    q.listWarehouse02ByUserIDStmt,
		updateCouponStmt:               q.updateCouponStmt,
		updateDiffLvStmt:               q.updateDiffLvStmt,
		updateInventoryQuantityStmt:    q.updateInventoryQuantityStmt,
		updateItemStmt:                 q.updateItemStmt,
		updatePrizeStmt:                q.updatePrizeStmt,
		updateUserCouponStmt:           q.updateUserCouponStmt,
		updateUserCreditStmt:           q.updateUserCreditStmt,
	}
}
