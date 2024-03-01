package model

type SidecarBackendModel struct {
	Id                   string `gorm:"column:id;type:varchar(255);not null`
	Name                 string `gorm:"column:name;type:varchar(255);not null`
	ServiceType          string `gorm:"column:service_type;type:varchar(255);not null`
	OperatingSystem      string `gorm:"column:operating_system;type:varchar(255);not null`
	ExecutablePath       string `gorm:"column:executable_path;type:varchar(255);not null`
	ExecuteParameters    string `gorm:"column:execute_parameters;type:varchar(255);not null`
	ValidationParameters string `gorm:"column:validation_parameters;type:varchar(255);not null`
}

func (receiver SidecarBackendModel) TableName() string {
	return "sidecar_backend"
}
