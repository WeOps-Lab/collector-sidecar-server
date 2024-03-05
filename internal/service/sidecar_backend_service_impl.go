package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"github.com/acmestack/gorm-plus/gplus"
)

type SidecarBackendServiceImpl struct {
}

func (s SidecarBackendServiceImpl) BuildModelFromEntity(target entity.SidecarBackendEntity) model.SidecarBackendModel {
	return model.SidecarBackendModel{
		Name:                 target.Name,
		ServiceType:          target.ServiceType,
		OperatingSystem:      target.OperatingSystem,
		ExecutablePath:       target.ExecutablePath,
		ExecuteParameters:    target.ExecuteParameters,
		ValidationParameters: target.ValidationParameters,
	}
}
func (s SidecarBackendServiceImpl) BuildEntityFromModel(m *model.SidecarBackendModel) entity.SidecarBackendEntity {
	return entity.SidecarBackendEntity{
		Name:                 m.Name,
		ServiceType:          m.ServiceType,
		OperatingSystem:      m.OperatingSystem,
		ExecutablePath:       m.ExecutablePath,
		ExecuteParameters:    m.ExecuteParameters,
		ValidationParameters: m.ValidationParameters,
	}
}

func (s SidecarBackendServiceImpl) ListBackend(current int, size int, queryParams map[string][]string) entity.SidecarBackendListEntity {
	page := gplus.NewPage[model.SidecarBackendModel](current, size)
	pagerList, _ := gplus.SelectPage(page, gplus.BuildQuery[model.SidecarBackendModel](queryParams))

	var responseList []entity.SidecarBackendWrapperEntity
	for _, obj := range pagerList.Records {
		item := entity.SidecarBackendWrapperEntity{
			SidecarBackendEntity: s.BuildEntityFromModel(obj),
			Id:                   obj.ID,
		}
		responseList = append(responseList, item)
	}

	return entity.SidecarBackendListEntity{
		Items: responseList,
		PagerEntity: entity.PagerEntity{
			Current: pagerList.Current,
			Size:    pagerList.Size,
			Total:   pagerList.Total,
		},
	}
}

func (s SidecarBackendServiceImpl) GetBackend(id uint) entity.SidecarBackendWrapperEntity {
	target, _ := gplus.SelectById[model.SidecarBackendModel](id)

	return entity.SidecarBackendWrapperEntity{
		Id:                   target.ID,
		SidecarBackendEntity: s.BuildEntityFromModel(target),
	}
}

func (s SidecarBackendServiceImpl) CreateBackend(target entity.SidecarBackendEntity) error {
	result := s.BuildModelFromEntity(target)
	return gplus.Insert[model.SidecarBackendModel](&result).Error
}

func (s SidecarBackendServiceImpl) UpdateBackend(target entity.SidecarBackendWrapperEntity) error {
	obj := s.BuildModelFromEntity(target.SidecarBackendEntity)
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
