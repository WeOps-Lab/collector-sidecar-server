package repo

import "collector-sidecar-server/internal/model"

type SidecarBackendRepo interface {
	Create(agentInfo *model.SidecarBackendModel) error
	Update(agentInfo *model.SidecarBackendModel) error
	GetByNodeId(nodeId string) (*model.SidecarBackendModel, error)
	GetAll() ([]*model.SidecarBackendModel, error)
	DeleteByNodeId(nodeId string) error
}
