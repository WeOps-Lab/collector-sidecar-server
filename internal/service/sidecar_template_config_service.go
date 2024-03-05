package service

import "collector-sidecar-server/internal/entity"

type SidecarTemplateConfigService interface {
	ListTemplateConfigs(current int, size int, queryParams map[string][]string) entity.SidecarTemplateConfigListEntity
	GetTemplateConfig(id uint) entity.SidecarTemplateConfigWrapperEntity
	CreateTemplateConfig(target entity.SidecarTemplateConfigEntity) error
	UpdateTemplateConfig(target entity.SidecarTemplateConfigWrapperEntity) error
	DeleteTemplateConfig(id uint) error
}
