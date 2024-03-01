package model

type SidecarTemplateConfigModel struct {
	Id               string              `gorm:"column:id;type:varchar(255);not null;primary_key:true"`
	Name             string              `gorm:"column:name;type:varchar(255);not null"`
	ConfigTemplate   string              `gorm:"column:config_template;type:varchar(255);not null"`
	SidecarBackend   SidecarBackendModel `gorm:"foreignkey:SidecarBackendId;references:Id"`
	SidecarBackendId string              `gorm:"column:sidecar_backend_id;type:varchar(255);not null"`
}

func (SidecarTemplateConfigModel) TableName() string {
	return "sidecar_template_config"
}
