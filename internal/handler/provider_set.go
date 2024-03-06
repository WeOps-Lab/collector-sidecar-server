package handler

import (
	v1 "collector-sidecar-server/internal/handler/v1"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	v1.NewSidecarHandler,
	v1.NewSidecarAgentInfoHandler,
	v1.NewSidecarBackendHandler,
	v1.NewSidecarTemplateConfigHandler,
	v1.NewSidecarTokenHandler,
)
