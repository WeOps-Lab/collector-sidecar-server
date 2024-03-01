package db

import (
	"collector-sidecar-server/pkg/config"
	"context"

	"gorm.io/gorm"
)

var _ IDataSource = (*defaultMysqlDataSource)(nil)

// defaultMysqlDataSource 默认mysql数据源实现
type defaultMysqlDataSource struct {
	master *gorm.DB // 定义私有属性，用来持有主库链接，防止每次创建，创建后直接返回该变量。
}

func (d *defaultMysqlDataSource) Master(ctx context.Context) *gorm.DB {
	// 事物, 根据事物的key取出tx
	tx, ok := ctx.Value(ContextTxKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	if d.master == nil {
		panic("The [master] connection is nil, Please initialize it first.")
	}
	return d.master
}

func (d *defaultMysqlDataSource) Close() {
	// 关闭主库链接
	if d.master != nil {
		m, err := d.master.DB()
		if err != nil {
			_ = m.Close()
		}
	}
}

func NewDefaultMysql(c config.DBConfig) *defaultMysqlDataSource {
	return &defaultMysqlDataSource{
		master: GetMysqlConn(
			c.Dsn,
			c.MaximumPoolSize,
			c.MaximumIdleSize),
	}

}
