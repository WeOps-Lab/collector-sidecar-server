package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewSidecarService,
	wire.Bind(new(SidecarService), new(*SidecarServiceImpl)),
	NewSidecarAgentInfoService,
	wire.Bind(new(SidecarAgentInfoService), new(*SidecarAgentInfoServiceImpl)),
	NewSidecarBackendService,
	wire.Bind(new(SidecarBackendService), new(*SidecarBackendServiceImpl)),
	NewSidecarTemplateConfigService,
	wire.Bind(new(SidecarTemplateConfigService), new(*SidecarTemplateConfigImpl)),
	NewSidecarTokenService,
	wire.Bind(new(SidecarTokenService), new(*SidecarTokenServiceImpl)),
)
