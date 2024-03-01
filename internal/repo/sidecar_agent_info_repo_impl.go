package repo

import (
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/pkg/db"
	"context"
)

type SidecarAgentInfoRepoImpl struct {
	ds db.IDataSource
}

func (s SidecarAgentInfoRepoImpl) Create(agentInfo *model.SidecarAgentInfoModel) error {
	s.ds.Master(context.TODO()).Create(agentInfo)
	return nil
}

func (s SidecarAgentInfoRepoImpl) Update(agentInfo *model.SidecarAgentInfoModel) error {
	s.ds.Master(context.TODO()).Save(agentInfo)
	return nil
}

func (s SidecarAgentInfoRepoImpl) GetByNodeId(nodeId string) (*model.SidecarAgentInfoModel, error) {
	result := &model.SidecarAgentInfoModel{}
	s.ds.Master(context.TODO()).Where("node_id = ?", nodeId).First(&result)
	return result, nil
}

func (s SidecarAgentInfoRepoImpl) GetAll() ([]*model.SidecarAgentInfoModel, error) {
	var result []*model.SidecarAgentInfoModel
	s.ds.Master(context.TODO()).Find(&result)
	return result, nil
}

func (s SidecarAgentInfoRepoImpl) DeleteByNodeId(nodeId string) error {
	s.ds.Master(context.TODO()).Where("node_id = ?", nodeId).Delete(&model.SidecarAgentInfoModel{})
	return nil
}

func NewSidecarAgentInfoRepo(ds db.IDataSource) *SidecarAgentInfoRepoImpl {
	return &SidecarAgentInfoRepoImpl{
		ds: ds,
	}
}

var _ SidecarAgentInfoRepo = (*SidecarAgentInfoRepoImpl)(nil)
