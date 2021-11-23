package infrastructure

import (
	"context"
	"database/sql"
)

type txContextKey string

const txCtxKey txContextKey = "tx"

type sqlHandler struct {
	conn *sql.DB
}

type SqlHandler interface {
	DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

func NewSqlHandler(conn *sql.DB) SqlHandler {
	return &sqlHandler{
		conn,
	}
}

func (h *sqlHandler) DoInTx(ctx context.Context, f func(ctx context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := h.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}
	ctx = context.WithValue(ctx, txCtxKey, tx)
	v, err := f(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		return nil, err
	}
	return v, nil
}

func (h *sqlHandler) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	if tx, ok := ctx.Value(txCtxKey).(*sql.Tx); ok {
		return tx.PrepareContext(ctx, query)
	}
	return h.conn.PrepareContext(ctx, query)
}
