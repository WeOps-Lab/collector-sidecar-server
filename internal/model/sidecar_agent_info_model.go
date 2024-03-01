package model

type SidecarAgentInfoModel struct {
	NodeId      string `gorm:"column:node_id;type:varchar(255);not null;primary_key" json:"node_id"`
	NodeDetails string `gorm:"column:node_details;type:text;default:'{}';" json:"node_details"`
	AgentConfig string `gorm:"column:agent_config;type:text;default:'{}';" json:"agent_config"`
}

func (receiver SidecarAgentInfoModel) TableName() string {
	return "sidecar_agent_info"
}
