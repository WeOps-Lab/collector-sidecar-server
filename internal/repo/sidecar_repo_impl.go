package repo

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/pkg/db"
	"context"
	"encoding/json"
)

type SidecarRepoImpl struct {
	ds db.IDataSource
}

func (s SidecarRepoImpl) GetSidecarTemplateConfig(todo context.Context, id string) *model.SidecarTemplateConfigModel {
	// get sidecar template config from db
	var target model.SidecarTemplateConfigModel
	s.ds.Master(todo).Where("id = ?", id).First(&target)
	return &target
}

func (s SidecarRepoImpl) ListAgentBackend(ctx context.Context) []model.SidecarBackendModel {
	// list sidecar agent backend from db
	var target []model.SidecarBackendModel
	s.ds.Master(ctx).Find(&target)
	return target
}

func (s SidecarRepoImpl) GetAgentInfo(ctx context.Context, nodeId string) *model.SidecarAgentInfoModel {
	// get sidecar agent info from db
	var target model.SidecarAgentInfoModel
	s.ds.Master(ctx).Where("node_id = ?", nodeId).First(&target)
	return &target
}

func (s SidecarRepoImpl) SaveAgentInfo(ctx context.Context, nodeId string, req *entity.RegistrationRequest) {
	// create or update sidecar agent info
	// convert req to model
	nodeDetailsBytes, _ := json.Marshal(req.NodeDetails)
	target := model.SidecarAgentInfoModel{
		NodeId: nodeId,
		//serilize req.NodeDetails to String and store
		NodeDetails: string(nodeDetailsBytes),
	}
	// save to db
	s.ds.Master(ctx).Save(&target)
}

func NewSidecarRepo(ds db.IDataSource) *SidecarRepoImpl {
	return &SidecarRepoImpl{
		ds: ds,
	}
}

var _ SidecarRepo = (*SidecarRepoImpl)(nil)
