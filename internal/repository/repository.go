package repository

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	WithTx(tx *gorm.DB) Repository
}

type TransactionManager interface {
	Transaction(ctx context.Context, fn func(ctx context.Context) error) error
}

type transactionManager struct {
	db *gorm.DB
}

func NewTransactionManager(db *gorm.DB) TransactionManager {
	return &transactionManager{db: db}
}

func (tm *transactionManager) Transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	return tm.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, "tx", tx)
		return fn(ctx)
	})
}

func GetTxFromContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return nil
}
