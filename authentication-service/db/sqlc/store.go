package db

import (
	"context"
	"database/sql"
)

type Store interface {
	Querier
	CreateUserTx(cxt context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(cxt context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(cxt, nil)
	if err != nil {
		return err
	}
	// Pass the transaction object to the queries object
	q := New(tx)
	err = fn(q)
	if err != nil {
		// Rollback the transaction if there is an error
		if rbErr := tx.Rollback(); rbErr != nil {
			return rbErr
		}
		return err
	}
	// Commit the transaction if there is no error
	return tx.Commit()
}
