package entity

import "time"

type RegistrationRequest struct {
	NodeName    string             `json:"node_name"`
	NodeDetails NodeDetailsRequest `json:"node_details"`
}

type NodeDetailsRequest struct {
	OperatingSystem                 string          `json:"operating_system"`
	IP                              string          `json:"ip,omitempty"`
	LogFileList                     []File          `json:"log_file_list,omitempty"`
	Metrics                         *MetricsRequest `json:"metrics,omitempty"`
	Status                          *StatusRequest  `json:"status,omitempty"`
	CollectorConfigurationDirectory string          `json:"collector_configuration_directory,omitempty"`
	Tags                            []string        `json:"tags,omitempty"`
}

type File struct {
	Path    string    `json:"path"`
	ModTime time.Time `json:"mod_time"`
	Size    int64     `json:"size"`
	IsDir   bool      `json:"is_dir"`
}

type MetricsRequest struct {
	Disks75 []string `json:"disks_75"`
	CpuIdle float64  `json:"cpu_idle"`
	Load1   float64  `json:"load_1"`
}

type StatusRequest struct {
	Backends []StatusRequestBackend `json:"collectors"`
	Status   int                    `json:"status"`
	Message  string                 `json:"message"`
}

type StatusRequestBackend struct {
	CollectorId     string `json:"collector_id"`
	ConfigurationId string `json:"configuration_id,omitempty"`
	Status          int    `json:"status"`
	Message         string `json:"message"`
	VerboseMessage  string `json:"verbose_message"`
}

type ServerVersionResponse struct {
	ClusterId string `json:"cluster_id"`
	NodeId    string `json:"node_id"`
	Version   string `json:"version"`
}

type ResponseCollectorRegistration struct {
	Configuration         ResponseCollectorRegistrationConfiguration `json:"configuration"`
	ConfigurationOverride bool                                       `json:"configuration_override"`
	CollectorActions      []ResponseCollectorAction                  `json:"actions,omitempty"`
	Assignments           []ConfigurationAssignment                  `json:"assignments,omitempty"`
}

type ResponseCollectorRegistrationConfiguration struct {
	UpdateInterval int  `json:"update_interval"`
	SendStatus     bool `json:"send_status"`
}

type ResponseCollectorAction struct {
	BackendId  string                 `json:"collector_id"`
	Properties map[string]interface{} `json:"properties"`
}

type ConfigurationAssignment struct {
	BackendId       string `json:"collector_id"`
	ConfigurationId string `json:"configuration_id"`
}

type ResponseBackendList struct {
	Backends []ResponseCollectorBackend `json:"collectors"`
}

type ResponseCollectorBackend struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	ServiceType          string `json:"service_type"`
	OperatingSystem      string `json:"node_operating_system"`
	ExecutablePath       string `json:"executable_path"`
	ExecuteParameters    string `json:"execute_parameters"`
	ValidationParameters string `json:"validation_parameters"`
}

type ResponseCollectorConfiguration struct {
	ConfigurationId string `json:"id"`
	BackendId       string `json:"collector_id"`
	Name            string `json:"name"`
	Template        string `json:"template"`
}
