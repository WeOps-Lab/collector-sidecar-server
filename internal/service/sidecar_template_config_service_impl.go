package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/mitchellh/mapstructure"
)

type SidecarTemplateConfigImpl struct {
}

func (s SidecarTemplateConfigImpl) BuildModelFromEntity(target entity.SidecarTemplateConfigEntity) model.SidecarTemplateConfigModel {
	var result model.SidecarTemplateConfigModel
	mapstructure.Decode(target, &result)
	return result
}

func (s SidecarTemplateConfigImpl) BuildEntityFromModel(m *model.SidecarTemplateConfigModel) entity.SidecarTemplateConfigEntity {
	var result entity.SidecarTemplateConfigEntity
	mapstructure.Decode(m, &result)
	return result
}
func (s SidecarTemplateConfigImpl) ListTemplateConfigs(current int, size int, queryParams map[string][]string) entity.SidecarTemplateConfigListEntity {
	page := gplus.NewPage[model.SidecarTemplateConfigModel](current, size)
	pagerList, _ := gplus.SelectPage(page, gplus.BuildQuery[model.SidecarTemplateConfigModel](queryParams))

	var responseList []entity.SidecarTemplateConfigWrapperEntity
	for _, model := range pagerList.Records {
		responseList = append(responseList, entity.SidecarTemplateConfigWrapperEntity{
			Id:                          model.ID,
			SidecarTemplateConfigEntity: s.BuildEntityFromModel(model),
		})
	}
	return entity.SidecarTemplateConfigListEntity{
		Items: responseList,
		PagerEntity: entity.PagerEntity{
			Current: pagerList.Current,
			Size:    pagerList.Size,
			Total:   pagerList.Total,
		},
	}
}

func (s SidecarTemplateConfigImpl) GetTemplateConfig(id uint) entity.SidecarTemplateConfigWrapperEntity {
	target, _ := gplus.SelectById[model.SidecarTemplateConfigModel](id)
	return entity.SidecarTemplateConfigWrapperEntity{
		Id:                          target.ID,
		SidecarTemplateConfigEntity: s.BuildEntityFromModel(target),
	}
}

func (s SidecarTemplateConfigImpl) CreateTemplateConfig(target entity.SidecarTemplateConfigEntity) error {
	result := s.BuildModelFromEntity(target)
	return gplus.Insert[model.SidecarTemplateConfigModel](&result).Error
}

func (s SidecarTemplateConfigImpl) UpdateTemplateConfig(target entity.SidecarTemplateConfigWrapperEntity) error {
	obj := s.BuildModelFromEntity(target.SidecarTemplateConfigEntity)
	obj.ID = target.Id
	return gplus.UpdateById[model.SidecarTemplateConfigModel](&obj).Error
}

func (s SidecarTemplateConfigImpl) DeleteTemplateConfig(id uint) error {
	return gplus.DeleteById[model.SidecarTemplateConfigModel](id).Error
}

func NewSidecarTemplateConfigService() *SidecarTemplateConfigImpl {
	return &SidecarTemplateConfigImpl{}
}

var _ SidecarTemplateConfigService = (*SidecarTemplateConfigImpl)(nil)
