package service

import "collector-sidecar-server/internal/entity"

type SidecarTokenService interface {
	List(current int, size int, queryParams map[string][]string) entity.SidecarTokenListEntity
	Get(id uint) entity.SidecarTokenWrapperEntity
	Create(target entity.SidecarTokenEntity) error
	Update(target entity.SidecarTokenWrapperEntity) error
	Delete(id uint) error
}
