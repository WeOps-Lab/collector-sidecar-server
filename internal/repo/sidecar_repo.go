package repo

import (
	"context"
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
)

type SidecarRepo interface {
	SaveAgentInfo(ctx context.Context, nodeId string, req *entity.RegistrationRequest)
	GetAgentInfo(ctx context.Context, nodeId string) *model.SidecarAgentInfoModel
	ListAgentBackend(ctx context.Context) []model.SidecarBackendModel
	GetSidecarTemplateConfig(todo context.Context, id string) *model.SidecarTemplateConfigModel
}
