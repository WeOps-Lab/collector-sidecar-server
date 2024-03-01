package service

import (
	"collector-sidecar-server/internal/entity"
)

type SidecarAgentInfoService interface {
	ListAgentInfo() entity.SidecarAgentInfoListResponse
	GetAgentInfo(id string) entity.SidecarAgentInfo
	UpdateAgentConfig(id string, target entity.SidecarAgentInfo) error
	DeleteAgentInfo(id string) error
}
