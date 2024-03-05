package model

import "gorm.io/gorm"

type SidecarTemplateConfigModel struct {
	gorm.Model
	Name             string
	ConfigTemplate   string `gorm:"type:text"`
	SidecarBackend   SidecarBackendModel
	SidecarBackendID uint
}

func (SidecarTemplateConfigModel) TableName() string {
	return "sidecar_template_config"
}
