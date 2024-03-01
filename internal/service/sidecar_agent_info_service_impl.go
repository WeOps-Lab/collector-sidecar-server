package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/internal/repo"
)

type SidecarAgentInfoServiceImpl struct {
	repository repo.SidecarAgentInfoRepo
}

func (s SidecarAgentInfoServiceImpl) ListAgentInfo() entity.SidecarAgentInfoListResponse {
	modelResult, _ := s.repository.GetAll()
	var result []entity.SidecarAgentInfo
	for _, model := range modelResult {
		result = append(result, entity.SidecarAgentInfo{
			NodeId:      model.NodeId,
			NodeDetails: model.NodeDetails,
			AgentConfig: model.AgentConfig,
		})
	}
	return entity.SidecarAgentInfoListResponse{
		AgentInfoList: result,
	}
}

func (s SidecarAgentInfoServiceImpl) GetAgentInfo(id string) entity.SidecarAgentInfo {
	model, _ := s.repository.GetByNodeId(id)
	return entity.SidecarAgentInfo{
		NodeId:      model.NodeId,
		NodeDetails: model.NodeDetails,
		AgentConfig: model.AgentConfig,
	}
}

func (s SidecarAgentInfoServiceImpl) UpdateAgentConfig(id string, target entity.SidecarAgentInfo) error {
	updateEntity := &model.SidecarAgentInfoModel{
		NodeId:      id,
		NodeDetails: target.NodeDetails,
		AgentConfig: target.AgentConfig,
	}
	s.repository.Update(updateEntity)
	return nil
}

func (s SidecarAgentInfoServiceImpl) DeleteAgentInfo(id string) error {
	s.repository.DeleteByNodeId(id)
	return nil
}

func NewSidecarAgentInfoService(repository repo.SidecarAgentInfoRepo) *SidecarAgentInfoServiceImpl {
	return &SidecarAgentInfoServiceImpl{
		repository: repository,
	}
}

var _ SidecarAgentInfoService = (*SidecarAgentInfoServiceImpl)(nil)
