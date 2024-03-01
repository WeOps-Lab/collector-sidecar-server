package db

import (
	"context"

	"gorm.io/gorm"
)

type ContextTxKey struct{}

// 事物默认实现
type CoreTransaction struct {
	ds IDataSource
}

func NewTransaction(_ds IDataSource) *CoreTransaction {
	return &CoreTransaction{ds: _ds}
}

func (t *CoreTransaction) Execute(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.ds.Master(ctx).Transaction(func(tx *gorm.DB) error {
		withValue := context.WithValue(ctx, ContextTxKey{}, tx)
		return fn(withValue)
	})
}
