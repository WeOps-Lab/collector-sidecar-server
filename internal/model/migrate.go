package model

import (
	"gorm.io/gorm"
)

func MigrateAllModel(ds *gorm.DB) {
	ds.AutoMigrate(&SidecarAgentInfoModel{})
	ds.AutoMigrate(&SidecarBackendModel{})
	ds.AutoMigrate(&SidecarTemplateConfigModel{})
}
