package entity

type SidecarTemplateConfigListResponse struct {
	TemplateConfigs []SidecarTemplateConfig `json:"template_configs"`
}

type SidecarTemplateConfig struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	ConfigTemplate string `json:"config_template"`
	BackendId      uint   `json:"backend_id"`
}
