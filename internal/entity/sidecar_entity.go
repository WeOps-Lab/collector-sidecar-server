package entity

import "time"

type RegistrationSidecarEntity struct {
	NodeName    string            `json:"node_name"`
	NodeDetails NodeDetailsEntity `json:"node_details"`
}

type NodeDetailsEntity struct {
	OperatingSystem                 string         `json:"operating_system"`
	IP                              string         `json:"ip,omitempty"`
	LogFileList                     []FileEntity   `json:"log_file_list,omitempty"`
	Metrics                         *MetricsEntity `json:"metrics,omitempty"`
	Status                          *StatusEntity  `json:"status,omitempty"`
	CollectorConfigurationDirectory string         `json:"collector_configuration_directory,omitempty"`
	Tags                            []string       `json:"tags,omitempty"`
}

type FileEntity struct {
	Path    string    `json:"path"`
	ModTime time.Time `json:"mod_time"`
	Size    int64     `json:"size"`
	IsDir   bool      `json:"is_dir"`
}

type MetricsEntity struct {
	Disks75 []string `json:"disks_75"`
	CpuIdle float64  `json:"cpu_idle"`
	Load1   float64  `json:"load_1"`
}

type StatusEntity struct {
	Backends []StatusBackendEntity `json:"collectors"`
	Status   int                   `json:"status"`
	Message  string                `json:"message"`
}

type StatusBackendEntity struct {
	CollectorId     string `json:"collector_id"`
	ConfigurationId string `json:"configuration_id,omitempty"`
	Status          int    `json:"status"`
	Message         string `json:"message"`
	VerboseMessage  string `json:"verbose_message"`
}

type ServerVersionEntity struct {
	ClusterId string `json:"cluster_id"`
	NodeId    string `json:"node_id"`
	Version   string `json:"version"`
}

type CollectorRegistrationEntity struct {
	Configuration         ResponseCollectorRegistrationConfigurationEntity `json:"configuration"`
	ConfigurationOverride bool                                             `json:"configuration_override"`
	CollectorActions      []ResponseCollectorActionEntity                  `json:"actions,omitempty"`
	Assignments           []ConfigurationAssignmentEntity                  `json:"assignments,omitempty"`
}

type ResponseCollectorRegistrationConfigurationEntity struct {
	UpdateInterval int  `json:"update_interval"`
	SendStatus     bool `json:"send_status"`
}

type ResponseCollectorActionEntity struct {
	BackendId  string                 `json:"collector_id"`
	Properties map[string]interface{} `json:"properties"`
}

type ConfigurationAssignmentEntity struct {
	BackendId       string `json:"collector_id"`
	ConfigurationId string `json:"configuration_id"`
}

type BackendListEntity struct {
	Backends []CollectorBackendEntity `json:"collectors"`
}

type CollectorBackendEntity struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	ServiceType          string `json:"service_type"`
	OperatingSystem      string `json:"node_operating_system"`
	ExecutablePath       string `json:"executable_path"`
	ExecuteParameters    string `json:"execute_parameters"`
	ValidationParameters string `json:"validation_parameters"`
}

type CollectorConfigurationEntity struct {
	ConfigurationId string `json:"id"`
	BackendId       uint   `json:"collector_id"`
	Name            string `json:"name"`
	Template        string `json:"template"`
}
