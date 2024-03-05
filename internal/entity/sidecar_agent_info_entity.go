package entity

type SidecarAgentInfoListEntity struct {
	PagerEntity
	Items []SidecarAgentInfoWrapperEntity `json:"items"`
}

type SidecarAgentInfoWrapperEntity struct {
	SidecarAgentInfoEntity
	Id uint `json:"id"`
}

type SidecarAgentInfoEntity struct {
	NodeId      string                      `json:"node_id"`
	NodeDetails NodeDetailsEntity           `json:"node_details"`
	AgentConfig CollectorRegistrationEntity `json:"agent_config"`
}
