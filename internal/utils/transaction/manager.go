package transaction

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type txKeyType struct{}

var txKey = txKeyType{}

type TransactionManager interface {
	Do(ctx context.Context, fn func(ctx context.Context) error) error
}

type transactionManager struct {
	pool *pgxpool.Pool
}

func NewTransactionManager(pool *pgxpool.Pool) TransactionManager {
	return &transactionManager{pool: pool}
}

func (tm *transactionManager) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	tx, err := tm.pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			_ = tx.Rollback(ctx)
		}
	}()

	ctx = context.WithValue(ctx, txKey, tx)

	err = fn(ctx)
	if err != nil {
		return err
	}

	err = tx.Commit(ctx)
	return err
}

func TxFromCtx(ctx context.Context) (pgx.Tx, error) {
	tx, ok := ctx.Value(txKey).(pgx.Tx)
	if !ok {
		return nil, errors.New("no transaction in context")
	}
	return tx, nil
}
