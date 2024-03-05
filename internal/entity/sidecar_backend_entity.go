package entity

type SidecarBackendListEntity struct {
	PagerEntity
	Items []SidecarBackendWrapperEntity `json:"items"`
}

type SidecarBackendWrapperEntity struct {
	SidecarBackendEntity
	Id uint `json:"id"`
}

type SidecarBackendEntity struct {
	Name                 string `json:"name"`
	ServiceType          string `json:"service_type"`
	OperatingSystem      string `json:"operating_system"`
	ExecutablePath       string `json:"executable_path"`
	ExecuteParameters    string `json:"execute_parameters"`
	ValidationParameters string `json:"validation_parameters"`
}
