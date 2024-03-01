package repo

import (
	"context"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/pkg/db"
)

type SidecarTemplateConfigRepoImpl struct {
	ds db.IDataSource
}

func (s SidecarTemplateConfigRepoImpl) Create(agentInfo *model.SidecarTemplateConfigModel) error {
	s.ds.Master(context.TODO()).Create(agentInfo)
	return nil
}

func (s SidecarTemplateConfigRepoImpl) Update(agentInfo *model.SidecarTemplateConfigModel) error {
	s.ds.Master(context.TODO()).Save(agentInfo)
	return nil
}

func (s SidecarTemplateConfigRepoImpl) GetByNodeId(nodeId string) (*model.SidecarTemplateConfigModel, error) {
	result := &model.SidecarTemplateConfigModel{}
	s.ds.Master(context.TODO()).Where("node_id = ?", nodeId).First(&result)
	return result, nil
}

func (s SidecarTemplateConfigRepoImpl) GetAll() ([]*model.SidecarTemplateConfigModel, error) {
	var result []*model.SidecarTemplateConfigModel
	s.ds.Master(context.TODO()).Find(&result)
	return result, nil
}

func (s SidecarTemplateConfigRepoImpl) DeleteByNodeId(nodeId string) error {
	s.ds.Master(context.TODO()).Where("node_id = ?", nodeId).Delete(&model.SidecarTemplateConfigModel{})
	return nil
}

func NewSidecarTemplateConfigRepo(ds db.IDataSource) *SidecarTemplateConfigRepoImpl {
	return &SidecarTemplateConfigRepoImpl{
		ds: ds,
	}
}

var _ SidecarTemplateConfigRepo = (*SidecarTemplateConfigRepoImpl)(nil)
