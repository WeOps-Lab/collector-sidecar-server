package repo

import "collector-sidecar-server/internal/model"

type SidecarAgentInfoRepo interface {
	Create(agentInfo *model.SidecarAgentInfoModel) error
	Update(agentInfo *model.SidecarAgentInfoModel) error
	GetByNodeId(nodeId string) (*model.SidecarAgentInfoModel, error)
	GetAll() ([]*model.SidecarAgentInfoModel, error)
	DeleteByNodeId(nodeId string) error
}
