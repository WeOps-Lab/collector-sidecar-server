package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/internal/repo"
)

type SidecarBackendServiceImpl struct {
	repository repo.SidecarBackendRepo
}

func (s SidecarBackendServiceImpl) ListBackends() entity.SidecarBackendListResponse {
	modelList, _ := s.repository.GetAll()
	var responseList []entity.SidecarBackend
	for _, model := range modelList {
		responseList = append(responseList, entity.SidecarBackend{
			Id:   model.Id,
			Name: model.Name,
		})
	}
	return entity.SidecarBackendListResponse{
		Backends: responseList,
	}
}

func (s SidecarBackendServiceImpl) GetBackend(id string) entity.SidecarBackend {
	target, _ := s.repository.GetByNodeId(id)
	return entity.SidecarBackend{
		Id:   target.Id,
		Name: target.Name,
	}
}

func (s SidecarBackendServiceImpl) CreateBackend(target entity.SidecarBackend) error {
	result := model.SidecarBackendModel{
		Id:   target.Id,
		Name: target.Name,
	}
	return s.repository.Create(&result)
}

func (s SidecarBackendServiceImpl) UpdateBackend(id string, target entity.SidecarBackend) error {
	result := model.SidecarBackendModel{
		Id:   target.Id,
		Name: target.Name,
	}
	return s.repository.Update(&result)
}

func (s SidecarBackendServiceImpl) DeleteBackend(id string) error {
	return s.repository.DeleteByNodeId(id)
}

func NewSidecarBackendService(repository repo.SidecarBackendRepo) *SidecarBackendServiceImpl {
	return &SidecarBackendServiceImpl{
		repository: repository,
	}
}

var _ SidecarBackendService = (*SidecarBackendServiceImpl)(nil)
