package entity

type SidecarBackendListResponse struct {
	Backends []SidecarBackend `json:"backends"`
}

type SidecarBackend struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	ServiceType          string `json:"service_type"`
	OperatingSystem      string `json:"operating_system"`
	ExecutablePath       string `json:"executable_path"`
	ExecuteParameters    string `json:"execute_parameters"`
	ValidationParameters string `json:"validation_parameters"`
}
