package service

import "collector-sidecar-server/internal/entity"

type SidecarBackendService interface {
	ListBackends() entity.SidecarBackendListResponse
	GetBackend(id string) entity.SidecarBackend
	CreateBackend(target entity.SidecarBackend) error
	UpdateBackend(id string, target entity.SidecarBackend) error
	DeleteBackend(id string) error
}
