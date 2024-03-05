package service

import (
	"collector-sidecar-server/internal/entity"
)

type SidecarBackendService interface {
	ListBackend(current int, size int, queryParams map[string][]string) entity.SidecarBackendListEntity
	GetBackend(id uint) entity.SidecarBackendWrapperEntity
	CreateBackend(target entity.SidecarBackendEntity) error
	UpdateBackend(target entity.SidecarBackendWrapperEntity) error
	DeleteBackend(id uint) error
}
