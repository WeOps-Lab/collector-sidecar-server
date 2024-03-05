package db

import (
	"context"
	"github.com/acmestack/gorm-plus/gplus"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

// IDataSource 定义数据库数据源接口，按照业务需求可以返回主库链接Master和从库链接Slave
type IDataSource interface {
	Master(ctx context.Context) *gorm.DB
	Close()
}

func GetSqliteConn(dsn string, maxPoolSize, maxIdle int) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		PrepareStmt: true,
	})
	gplus.Init(db)
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	return db

}

// 创建Postgres链接
func GetPostgresConn(dsn string, maxPoolSize, maxIdle int) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 缓存每一条sql语句，提高执行速度
	})
	gplus.Init(db)
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	return db
}

// GetMysqlConn 创建Mysql链接
func GetMysqlConn(dsn string, maxPoolSize, maxIdle int) *gorm.DB {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true, // 缓存每一条sql语句，提高执行速度
	})
	gplus.Init(db)
	if err != nil {
		panic(err)
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Hour)
	// 设置连接池大小
	sqlDb.SetMaxOpenConns(maxPoolSize)
	sqlDb.SetMaxIdleConns(maxIdle)
	return db
}
