package model

import (
	"collector-sidecar-server/internal/entity"
	"gorm.io/gorm"
)

type SidecarAgentInfoModel struct {
	gorm.Model
	NodeId      string
	NodeDetails entity.NodeDetailsEntity
	AgentConfig entity.CollectorRegistrationEntity `gorm:"type:json"`
}

func (receiver SidecarAgentInfoModel) TableName() string {
	return "sidecar_agent_info"
}
