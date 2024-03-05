package entity

type SidecarTemplateConfigListEntity struct {
	PagerEntity
	Items []SidecarTemplateConfigWrapperEntity `json:"items"`
}

type SidecarTemplateConfigWrapperEntity struct {
	Id uint `json:"id"`
	SidecarTemplateConfigEntity
}

type SidecarTemplateConfigEntity struct {
	Name           string `json:"name"`
	ConfigTemplate string `json:"config_template"`
	BackendId      uint   `json:"backend_id"`
}
