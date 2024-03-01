package repo

import (
	"collector-sidecar-server/pkg/db"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewTransaction,
	wire.Bind(new(db.Transaction), new(*db.CoreTransaction)),
	NewSidecarRepo,
	wire.Bind(new(SidecarRepo), new(*SidecarRepoImpl)),
	NewSidecarAgentInfoRepo,
	wire.Bind(new(SidecarAgentInfoRepo), new(*SidecarAgentInfoRepoImpl)),
	NewSidecarBackendRepo,
	wire.Bind(new(SidecarBackendRepo), new(*SidecarBackendRepoImpl)),
	NewSidecarTemplateConfigRepo,
	wire.Bind(new(SidecarTemplateConfigRepo), new(*SidecarTemplateConfigRepoImpl)),
)
