package router

import (
	"collector-sidecar-server/server"

	"github.com/google/wire"
)

var ApiRouterProviderSet = wire.NewSet(
	NewApiRouter,
	wire.Bind(new(server.Router), new(*ApiRouter)),
)
