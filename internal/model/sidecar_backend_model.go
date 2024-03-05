package model

import "gorm.io/gorm"

type SidecarBackendModel struct {
	gorm.Model
	Name                 string
	ServiceType          string
	OperatingSystem      string
	ExecutablePath       string
	ExecuteParameters    string
	ValidationParameters string
}

func (receiver SidecarBackendModel) TableName() string {
	return "sidecar_backend"
}
