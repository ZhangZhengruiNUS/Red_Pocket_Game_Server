// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

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
	if q.createInventoryStmt, err = db.PrepareContext(ctx, createInventory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateInventory: %w", err)
	}
	if q.createItemStmt, err = db.PrepareContext(ctx, createItem); err != nil {
		return nil, fmt.Errorf("error preparing query CreateItem: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteItemStmt, err = db.PrepareContext(ctx, deleteItem); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteItem: %w", err)
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
	if q.getUserByIdStmt, err = db.PrepareContext(ctx, getUserById); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserById: %w", err)
	}
	if q.getUserByNameStmt, err = db.PrepareContext(ctx, getUserByName); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserByName: %w", err)
	}
	if q.listGameDiffSetsStmt, err = db.PrepareContext(ctx, listGameDiffSets); err != nil {
		return nil, fmt.Errorf("error preparing query ListGameDiffSets: %w", err)
	}
	if q.listInventoriesByUserIDStmt, err = db.PrepareContext(ctx, listInventoriesByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query ListInventoriesByUserID: %w", err)
	}
	if q.listItemsStmt, err = db.PrepareContext(ctx, listItems); err != nil {
		return nil, fmt.Errorf("error preparing query ListItems: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
	}
	if q.updateInventoryQuantityStmt, err = db.PrepareContext(ctx, updateInventoryQuantity); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateInventoryQuantity: %w", err)
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
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteItemStmt != nil {
		if cerr := q.deleteItemStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteItemStmt: %w", cerr)
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
	if q.listGameDiffSetsStmt != nil {
		if cerr := q.listGameDiffSetsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listGameDiffSetsStmt: %w", cerr)
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
	if q.listUsersStmt != nil {
		if cerr := q.listUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listUsersStmt: %w", cerr)
		}
	}
	if q.updateInventoryQuantityStmt != nil {
		if cerr := q.updateInventoryQuantityStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateInventoryQuantityStmt: %w", cerr)
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
	createInventoryStmt            *sql.Stmt
	createItemStmt                 *sql.Stmt
	createUserStmt                 *sql.Stmt
	deleteItemStmt                 *sql.Stmt
	deleteUserStmt                 *sql.Stmt
	getInventoryByUserIDItemIDStmt *sql.Stmt
	getItemStmt                    *sql.Stmt
	getUserByIdStmt                *sql.Stmt
	getUserByNameStmt              *sql.Stmt
	listGameDiffSetsStmt           *sql.Stmt
	listInventoriesByUserIDStmt    *sql.Stmt
	listItemsStmt                  *sql.Stmt
	listUsersStmt                  *sql.Stmt
	updateInventoryQuantityStmt    *sql.Stmt
	updateUserCouponStmt           *sql.Stmt
	updateUserCreditStmt           *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                             tx,
		tx:                             tx,
		createInventoryStmt:            q.createInventoryStmt,
		createItemStmt:                 q.createItemStmt,
		createUserStmt:                 q.createUserStmt,
		deleteItemStmt:                 q.deleteItemStmt,
		deleteUserStmt:                 q.deleteUserStmt,
		getInventoryByUserIDItemIDStmt: q.getInventoryByUserIDItemIDStmt,
		getItemStmt:                    q.getItemStmt,
		getUserByIdStmt:                q.getUserByIdStmt,
		getUserByNameStmt:              q.getUserByNameStmt,
		listGameDiffSetsStmt:           q.listGameDiffSetsStmt,
		listInventoriesByUserIDStmt:    q.listInventoriesByUserIDStmt,
		listItemsStmt:                  q.listItemsStmt,
		listUsersStmt:                  q.listUsersStmt,
		updateInventoryQuantityStmt:    q.updateInventoryQuantityStmt,
		updateUserCouponStmt:           q.updateUserCouponStmt,
		updateUserCreditStmt:           q.updateUserCreditStmt,
	}
}
