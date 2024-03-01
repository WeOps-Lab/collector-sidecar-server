package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/internal/repo"
)

type SidecarTemplateConfigImpl struct {
	repository repo.SidecarTemplateConfigRepo
}

func (s SidecarTemplateConfigImpl) ListTemplateConfigs() entity.SidecarTemplateConfigListResponse {
	modelList, _ := s.repository.GetAll()
	var responseList []entity.SidecarTemplateConfig
	for _, model := range modelList {
		responseList = append(responseList, entity.SidecarTemplateConfig{
			Id:             model.Id,
			Name:           model.Name,
			ConfigTemplate: model.ConfigTemplate,
		})
	}
	return entity.SidecarTemplateConfigListResponse{
		TemplateConfigs: responseList,
	}
}

func (s SidecarTemplateConfigImpl) GetTemplateConfig(id string) entity.SidecarTemplateConfig {
	target, _ := s.repository.GetByNodeId(id)
	return entity.SidecarTemplateConfig{
		Id:             target.Id,
		Name:           target.Name,
		ConfigTemplate: target.ConfigTemplate,
		BackendId:      target.SidecarBackendId,
	}
}

func (s SidecarTemplateConfigImpl) CreateTemplateConfig(target entity.SidecarTemplateConfig) error {
	result := model.SidecarTemplateConfigModel{
		Id:               target.Id,
		Name:             target.Name,
		ConfigTemplate:   target.ConfigTemplate,
		SidecarBackendId: target.BackendId,
	}
	return s.repository.Create(&result)
}

func (s SidecarTemplateConfigImpl) UpdateTemplateConfig(id string, target entity.SidecarTemplateConfig) error {
	result := model.SidecarTemplateConfigModel{
		Id:               target.Id,
		Name:             target.Name,
		ConfigTemplate:   target.ConfigTemplate,
		SidecarBackendId: target.BackendId,
	}
	return s.repository.Update(&result)
}

func (s SidecarTemplateConfigImpl) DeleteTemplateConfig(id string) error {
	return s.repository.DeleteByNodeId(id)
}

func NewSidecarTemplateConfigService(repository repo.SidecarTemplateConfigRepo) *SidecarTemplateConfigImpl {
	return &SidecarTemplateConfigImpl{
		repository: repository,
	}
}

var _ SidecarTemplateConfigService = (*SidecarTemplateConfigImpl)(nil)
