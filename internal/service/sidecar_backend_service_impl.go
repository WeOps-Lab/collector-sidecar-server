package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"github.com/acmestack/gorm-plus/gplus"
)

type SidecarBackendServiceImpl struct {
}

func (s SidecarBackendServiceImpl) NewSidecarBackendEntityFromModel(m *model.SidecarBackendModel) entity.SidecarBackend {
	return entity.SidecarBackend{
		Name:                 m.Name,
		ServiceType:          m.ServiceType,
		OperatingSystem:      m.OperatingSystem,
		ExecutablePath:       m.ExecutablePath,
		ExecuteParameters:    m.ExecuteParameters,
		ValidationParameters: m.ValidationParameters,
	}

}

func (s SidecarBackendServiceImpl) ListBackend(current int, size int, queryParams map[string][]string) entity.SidecarBackendPageList {
	page := gplus.NewPage[model.SidecarBackendModel](current, size)
	modelList, _ := gplus.SelectPage(page, gplus.BuildQuery[model.SidecarBackendModel](queryParams))

	var responseList []entity.SidecarBackendWrapper
	for _, obj := range modelList.Records {
		item := entity.SidecarBackendWrapper{
			SidecarBackend: s.NewSidecarBackendEntityFromModel(obj),
			Id:             obj.ID,
		}
		responseList = append(responseList, item)
	}

	return entity.SidecarBackendPageList{
		Items: responseList,
		PagerEntity: entity.PagerEntity{
			Current: modelList.Current,
			Size:    modelList.Size,
			Total:   modelList.Total,
		},
	}
}

func (s SidecarBackendServiceImpl) GetBackend(id uint) entity.SidecarBackendWrapper {
	target, _ := gplus.SelectById[model.SidecarBackendModel](id)

	return entity.SidecarBackendWrapper{
		Id:             target.ID,
		SidecarBackend: s.NewSidecarBackendEntityFromModel(target),
	}
}

func (s SidecarBackendServiceImpl) CreateBackend(target entity.SidecarBackend) error {
	result := model.SidecarBackendModel{
		Name:                 target.Name,
		ServiceType:          target.ServiceType,
		OperatingSystem:      target.OperatingSystem,
		ExecutablePath:       target.ExecutablePath,
		ExecuteParameters:    target.ExecuteParameters,
		ValidationParameters: target.ValidationParameters,
	}
	return gplus.Insert[model.SidecarBackendModel](&result).Error
}

func (s SidecarBackendServiceImpl) UpdateBackend(target entity.SidecarBackendWrapper) error {
	obj := model.SidecarBackendModel{
		Name:                 target.Name,
		ServiceType:          target.ServiceType,
		OperatingSystem:      target.OperatingSystem,
		ExecutablePath:       target.ExecutablePath,
		ExecuteParameters:    target.ExecuteParameters,
		ValidationParameters: target.ValidationParameters,
	}
	obj.ID = target.Id
	return gplus.UpdateById[model.SidecarBackendModel](&obj).Error
}

func (s SidecarBackendServiceImpl) DeleteBackend(id uint) error {
	return gplus.DeleteById[model.SidecarBackendModel](id).Error
}

func NewSidecarBackendService() *SidecarBackendServiceImpl {
	return &SidecarBackendServiceImpl{}
}

var _ SidecarBackendService = (*SidecarBackendServiceImpl)(nil)
