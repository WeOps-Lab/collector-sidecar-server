package entity

type SidecarBackendPageList struct {
	PagerEntity
	Items []SidecarBackendWrapper `json:"items"`
}

type SidecarBackendWrapper struct {
	SidecarBackend
	Id uint `json:"id"`
}

type SidecarBackend struct {
	Name                 string `json:"name"`
	ServiceType          string `json:"service_type"`
	OperatingSystem      string `json:"operating_system"`
	ExecutablePath       string `json:"executable_path"`
	ExecuteParameters    string `json:"execute_parameters"`
	ValidationParameters string `json:"validation_parameters"`
}
