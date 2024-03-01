package entity

type SidecarAgentInfoListResponse struct {
	AgentInfoList []SidecarAgentInfo `json:"agent_info_list"`
}

type SidecarAgentInfo struct {
	NodeId      string `json:"node_id"`
	NodeDetails string `json:"node_details"`
	AgentConfig string `json:"agent_config"`
}
