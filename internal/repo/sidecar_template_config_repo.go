package repo

import "collector-sidecar-server/internal/model"

type SidecarTemplateConfigRepo interface {
	Create(agentInfo *model.SidecarTemplateConfigModel) error
	Update(agentInfo *model.SidecarTemplateConfigModel) error
	GetByNodeId(nodeId string) (*model.SidecarTemplateConfigModel, error)
	GetAll() ([]*model.SidecarTemplateConfigModel, error)
	DeleteByNodeId(nodeId string) error
}
