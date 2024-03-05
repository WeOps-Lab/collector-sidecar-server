package service

import (
	"collector-sidecar-server/internal/entity"
)

type SidecarAgentInfoService interface {
	ListAgentInfo(current int, size int, queryParams map[string][]string) entity.SidecarAgentInfoListEntity
	GetAgentInfo(id uint) entity.SidecarAgentInfoWrapperEntity
	UpdateAgentConfig(target entity.SidecarAgentInfoWrapperEntity) error
	DeleteAgentInfo(id uint) error
}
