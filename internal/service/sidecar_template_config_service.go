package service

import "collector-sidecar-server/internal/entity"

type SidecarTemplateConfigService interface {
	ListTemplateConfigs() entity.SidecarTemplateConfigListResponse
	GetTemplateConfig(id string) entity.SidecarTemplateConfig
	CreateTemplateConfig(target entity.SidecarTemplateConfig) error
	UpdateTemplateConfig(id string, target entity.SidecarTemplateConfig) error
	DeleteTemplateConfig(id string) error
}
