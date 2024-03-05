package db

import (
	"collector-sidecar-server/pkg/config"
	"context"
	"gorm.io/gorm"
)

var _ IDataSource = (*defaultSqliteDataSource)(nil)

type defaultSqliteDataSource struct {
	master *gorm.DB
}

func (d defaultSqliteDataSource) Master(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(ContextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	if d.master == nil {
		panic("The [master] connection is nil, Please initialize it first.")
	}
	return d.master
}

func (d defaultSqliteDataSource) Close() {
	// 关闭主库链接
	if d.master != nil {
		m, err := d.master.DB()
		if err != nil {
			_ = m.Close()
		}
	}
}

func NewDefaultSqlite(c config.DBConfig) *defaultSqliteDataSource {
	return &defaultSqliteDataSource{
		master: GetSqliteConn(
			c.Dsn,
			c.MaximumPoolSize,
			c.MaximumIdleSize),
	}

}
