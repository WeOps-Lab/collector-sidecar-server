package service

import "collector-sidecar-server/internal/entity"

type SidecarService interface {
	UpdateSidecarNodeInfo(nodeId string, req entity.RegistrationSidecarEntity) entity.CollectorRegistrationEntity
	RenderConfiguration(nodeId string, configurationId uint) entity.CollectorConfigurationEntity
	ListCollectors() entity.BackendListEntity
	GetServerInfo() entity.ServerVersionEntity
	GetConfigETag(nodeId string) string
	GetConfigBackendListETag(backendConfig string) string
	GetConfigurationETag(configurationId uint) string
}
