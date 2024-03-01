package service

import "collector-sidecar-server/internal/entity"

type SidecarService interface {
	UpdateSidecarNodeInfo(nodeId string, req entity.RegistrationRequest) entity.ResponseCollectorRegistration
	RenderConfiguration(nodeId string, configurationId string) entity.ResponseCollectorConfiguration
	ListCollectors() entity.ResponseBackendList
	GetServerInfo() entity.ServerVersionResponse
	GetConfigETag(nodeId string) string
	GetConfigBackendListETag(backendConfig string) string
	GetConfigurationETag(configurationId string) string
}
