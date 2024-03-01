package main

import (
	_ "collector-sidecar-server/docs"
	"collector-sidecar-server/internal/middleware"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/pkg/cache"
	"collector-sidecar-server/pkg/config"
	"collector-sidecar-server/pkg/db"
	"collector-sidecar-server/pkg/log"
	"collector-sidecar-server/pkg/version"
	"collector-sidecar-server/server"
	"context"

	_ "github.com/alecthomas/template"
	_ "github.com/swaggo/files"
	_ "github.com/swaggo/gin-swagger"
)

// @title           Swagger  API
// @version         1.0
// @host      localhost:8080
// @BasePath  /
func main() {
	// 解析服务器启动参数
	appOpt := &server.AppOptions{}
	server.ResolveAppOptions(appOpt)
	if appOpt.PrintVersion {
		version.PrintVersion()
	}

	// 加载配置文件
	c := config.Load(appOpt.ConfigFilePath)
	log.InitLogger(&c.LogConfig, c.AppName) // 日志
	defer log.Sync()

	var ds db.IDataSource
	if c.DBConfig.Enable {
		if c.DBConfig.DbType == "mysql" {
			log.Info("开启Mysql数据库连接")
			ds = db.NewDefaultMysql(c.DBConfig)
		}
		if c.DBConfig.DbType == "postgres" {
			log.Info("开启Postgres数据库连接")
			ds = db.NewDefaultPostgres(c.DBConfig)
		}
		if c.DBConfig.AutoMigrate {
			model.MigrateAllModel(ds.Master(context.Background()))
		}
	} else {
		log.Info("未开启数据库连接")
	}

	// 创建HTTPServer
	srv := server.NewHttpServer(config.GlobalConfig)
	srv.RegisterOnShutdown(func() {
		if ds != nil {
			ds.Close()
		}
	})

	var cacheClient cache.ICache
	if c.RedisConfig.Enable {
		log.Info("开启Redis连接")
		cacheClient = cache.NewDefaultRedisCache(c.RedisConfig)
	}
	router := initRouter(ds, cacheClient)

	srv.Run(middleware.NewMiddleware(), router)
}
