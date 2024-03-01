package repo

import (
	"context"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/pkg/db"
)

type SidecarBackendRepoImpl struct {
	ds db.IDataSource
}

func (s SidecarBackendRepoImpl) Create(agentInfo *model.SidecarBackendModel) error {
	s.ds.Master(context.TODO()).Create(agentInfo)
	return nil
}

func (s SidecarBackendRepoImpl) Update(agentInfo *model.SidecarBackendModel) error {
	s.ds.Master(context.TODO()).Save(agentInfo)
	return nil
}

func (s SidecarBackendRepoImpl) GetByNodeId(nodeId string) (*model.SidecarBackendModel, error) {
	result := &model.SidecarBackendModel{}
	s.ds.Master(context.TODO()).Where("node_id = ?", nodeId).First(&result)
	return result, nil
}

func (s SidecarBackendRepoImpl) GetAll() ([]*model.SidecarBackendModel, error) {
	var result []*model.SidecarBackendModel
	s.ds.Master(context.TODO()).Find(&result)
	return result, nil
}

func (s SidecarBackendRepoImpl) DeleteByNodeId(nodeId string) error {
	s.ds.Master(context.TODO()).Where("node_id = ?", nodeId).Delete(&model.SidecarBackendModel{})
	return nil
}

func NewSidecarBackendRepo(ds db.IDataSource) *SidecarBackendRepoImpl {
	return &SidecarBackendRepoImpl{
		ds: ds,
	}
}

var _ SidecarBackendRepo = (*SidecarBackendRepoImpl)(nil)
