//go:build wireinject
// +build wireinject

package main

import (
	"collector-sidecar-server/internal/handler"
	"collector-sidecar-server/internal/repo"
	"collector-sidecar-server/internal/router"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/cache"
	"collector-sidecar-server/pkg/db"
	"collector-sidecar-server/server"

	"github.com/google/wire"
)

// initRouter 初始化router
func initRouter(ds db.IDataSource, cacheClient cache.ICache) server.Router {
	wire.Build(
		providerSet,
		router.ApiRouterProviderSet,
	)
	return nil
}

var providerSet = wire.NewSet(repo.ProviderSet, service.ProviderSet, handler.ProviderSet)
