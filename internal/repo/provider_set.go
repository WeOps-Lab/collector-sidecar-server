package repo

import (
	"collector-sidecar-server/pkg/db"

	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	db.NewTransaction,
	wire.Bind(new(db.Transaction), new(*db.CoreTransaction)),
)
