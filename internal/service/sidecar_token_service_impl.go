package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"github.com/acmestack/gorm-plus/gplus"
)

type SidecarTokenServiceImpl struct{}

func (s SidecarTokenServiceImpl) List(current int, size int, queryParams map[string][]string) entity.SidecarTokenListEntity {
	page := gplus.NewPage[model.SidecarTokenModel](current, size)
	pagerList, _ := gplus.SelectPage(page, gplus.BuildQuery[model.SidecarTokenModel](queryParams))

	var responseList []entity.SidecarTokenWrapperEntity
	for _, model := range pagerList.Records {
		responseList = append(responseList, entity.SidecarTokenWrapperEntity{
			Id:                 model.ID,
			SidecarTokenEntity: s.BuildEntityFromModel(model),
		})
	}

	return entity.SidecarTokenListEntity{
		Items: responseList,
		PagerEntity: entity.PagerEntity{
			Current: pagerList.Current,
			Size:    pagerList.Size,
			Total:   pagerList.Total,
		},
	}
}

func (s SidecarTokenServiceImpl) Get(id uint) entity.SidecarTokenWrapperEntity {
	target, _ := gplus.SelectById[model.SidecarTokenModel](id)
	return entity.SidecarTokenWrapperEntity{
		Id:                 target.ID,
		SidecarTokenEntity: s.BuildEntityFromModel(target),
	}
}

func (s SidecarTokenServiceImpl) Create(target entity.SidecarTokenEntity) error {
	result := s.BuildModelFromEntity(target)
	return gplus.Insert[model.SidecarTokenModel](&result).Error
}

func (s SidecarTokenServiceImpl) Update(target entity.SidecarTokenWrapperEntity) error {
	obj := s.BuildModelFromEntity(target.SidecarTokenEntity)
	obj.ID = target.Id
	return gplus.UpdateById[model.SidecarTokenModel](&obj).Error
}

func (s SidecarTokenServiceImpl) Delete(id uint) error {
	return gplus.DeleteById[model.SidecarTokenModel](id).Error
}

func (s SidecarTokenServiceImpl) BuildModelFromEntity(target entity.SidecarTokenEntity) model.SidecarTokenModel {
	return model.SidecarTokenModel{
		Token: target.Token,
	}
}

func (s SidecarTokenServiceImpl) BuildEntityFromModel(m *model.SidecarTokenModel) entity.SidecarTokenEntity {
	return entity.SidecarTokenEntity{
		Token: m.Token,
	}
}

func NewSidecarTokenService() *SidecarTokenServiceImpl {
	return &SidecarTokenServiceImpl{}
}

var _ SidecarTokenService = (*SidecarTokenServiceImpl)(nil)
