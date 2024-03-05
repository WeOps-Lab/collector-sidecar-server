package service

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"github.com/acmestack/gorm-plus/gplus"
)

type SidecarAgentInfoServiceImpl struct {
}

func (s SidecarAgentInfoServiceImpl) BuildModelFromEntity(target entity.SidecarAgentInfoEntity) model.SidecarAgentInfoModel {
	return model.SidecarAgentInfoModel{
		NodeId:      target.NodeId,
		NodeDetails: target.NodeDetails,
		AgentConfig: target.AgentConfig,
	}
}
func (s SidecarAgentInfoServiceImpl) BuildEntityFromModel(m *model.SidecarAgentInfoModel) entity.SidecarAgentInfoEntity {
	return entity.SidecarAgentInfoEntity{
		NodeId:      m.NodeId,
		NodeDetails: m.NodeDetails,
		AgentConfig: m.AgentConfig,
	}
}
func (s SidecarAgentInfoServiceImpl) ListAgentInfo(current int, size int, queryParams map[string][]string) entity.SidecarAgentInfoListEntity {
	page := gplus.NewPage[model.SidecarAgentInfoModel](current, size)
	pagerList, _ := gplus.SelectPage(page, gplus.BuildQuery[model.SidecarAgentInfoModel](queryParams))

	var result []entity.SidecarAgentInfoWrapperEntity
	for _, obj := range pagerList.Records {
		result = append(result, entity.SidecarAgentInfoWrapperEntity{
			SidecarAgentInfoEntity: s.BuildEntityFromModel(obj),
			Id:                     obj.ID,
		})
	}
	return entity.SidecarAgentInfoListEntity{
		Items: result,
		PagerEntity: entity.PagerEntity{
			Current: pagerList.Current,
			Size:    pagerList.Size,
			Total:   pagerList.Total,
		},
	}
}

func (s SidecarAgentInfoServiceImpl) GetAgentInfo(id uint) entity.SidecarAgentInfoWrapperEntity {
	target, _ := gplus.SelectById[model.SidecarAgentInfoModel](id)
	return entity.SidecarAgentInfoWrapperEntity{
		Id:                     target.ID,
		SidecarAgentInfoEntity: s.BuildEntityFromModel(target),
	}
}

func (s SidecarAgentInfoServiceImpl) UpdateAgentConfig(target entity.SidecarAgentInfoWrapperEntity) error {
	obj := s.BuildModelFromEntity(target.SidecarAgentInfoEntity)
	obj.ID = target.Id
	return gplus.UpdateById[model.SidecarAgentInfoModel](&obj).Error
}

func (s SidecarAgentInfoServiceImpl) DeleteAgentInfo(id uint) error {
	return gplus.DeleteById[model.SidecarAgentInfoModel](id).Error
}

func NewSidecarAgentInfoService() *SidecarAgentInfoServiceImpl {
	return &SidecarAgentInfoServiceImpl{}
}

var _ SidecarAgentInfoService = (*SidecarAgentInfoServiceImpl)(nil)
