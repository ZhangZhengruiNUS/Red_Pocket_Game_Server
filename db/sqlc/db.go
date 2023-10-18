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
	if q.deleteInventoryStmt, err = db.PrepareContext(ctx, deleteInventory); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteInventory: %w", err)
	}
	if q.deleteItemStmt, err = db.PrepareContext(ctx, deleteItem); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteItem: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.getInventoryStmt, err = db.PrepareContext(ctx, getInventory); err != nil {
		return nil, fmt.Errorf("error preparing query GetInventory: %w", err)
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
	if q.listInventoriesStmt, err = db.PrepareContext(ctx, listInventories); err != nil {
		return nil, fmt.Errorf("error preparing query ListInventories: %w", err)
	}
	if q.listItemsStmt, err = db.PrepareContext(ctx, listItems); err != nil {
		return nil, fmt.Errorf("error preparing query ListItems: %w", err)
	}
	if q.listUsersStmt, err = db.PrepareContext(ctx, listUsers); err != nil {
		return nil, fmt.Errorf("error preparing query ListUsers: %w", err)
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
	if q.deleteInventoryStmt != nil {
		if cerr := q.deleteInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteInventoryStmt: %w", cerr)
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
	if q.getInventoryStmt != nil {
		if cerr := q.getInventoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getInventoryStmt: %w", cerr)
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
	if q.listInventoriesStmt != nil {
		if cerr := q.listInventoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing listInventoriesStmt: %w", cerr)
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
	db                  DBTX
	tx                  *sql.Tx
	createInventoryStmt *sql.Stmt
	createItemStmt      *sql.Stmt
	createUserStmt      *sql.Stmt
	deleteInventoryStmt *sql.Stmt
	deleteItemStmt      *sql.Stmt
	deleteUserStmt      *sql.Stmt
	getInventoryStmt    *sql.Stmt
	getItemStmt         *sql.Stmt
	getUserByIdStmt     *sql.Stmt
	getUserByNameStmt   *sql.Stmt
	listInventoriesStmt *sql.Stmt
	listItemsStmt       *sql.Stmt
	listUsersStmt       *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		createInventoryStmt: q.createInventoryStmt,
		createItemStmt:      q.createItemStmt,
		createUserStmt:      q.createUserStmt,
		deleteInventoryStmt: q.deleteInventoryStmt,
		deleteItemStmt:      q.deleteItemStmt,
		deleteUserStmt:      q.deleteUserStmt,
		getInventoryStmt:    q.getInventoryStmt,
		getItemStmt:         q.getItemStmt,
		getUserByIdStmt:     q.getUserByIdStmt,
		getUserByNameStmt:   q.getUserByNameStmt,
		listInventoriesStmt: q.listInventoriesStmt,
		listItemsStmt:       q.listItemsStmt,
		listUsersStmt:       q.listUsersStmt,
	}
}
