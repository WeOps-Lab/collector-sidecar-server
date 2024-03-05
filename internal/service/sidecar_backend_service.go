package service

import (
	"collector-sidecar-server/internal/entity"
)

type SidecarBackendService interface {
	ListBackend(current int, size int, queryParams map[string][]string) entity.SidecarBackendPageList
	GetBackend(id uint) entity.SidecarBackendWrapper
	CreateBackend(target entity.SidecarBackend) error
	UpdateBackend(target entity.SidecarBackendWrapper) error
	DeleteBackend(id uint) error
}
