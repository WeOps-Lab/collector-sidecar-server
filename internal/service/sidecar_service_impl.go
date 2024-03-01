package service

import (
	"bytes"
	"context"
	"encoding/json"
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/repo"
	"collector-sidecar-server/pkg/security"
	"text/template"
)

type SidecarServiceImpl struct {
	repository repo.SidecarRepo
}

func NewSidecarService(repository repo.SidecarRepo) *SidecarServiceImpl {
	return &SidecarServiceImpl{
		repository: repository,
	}
}

func (s *SidecarServiceImpl) UpdateSidecarNodeInfo(nodeId string, req entity.RegistrationRequest) entity.ResponseCollectorRegistration {
	s.repository.SaveAgentInfo(context.TODO(), nodeId, &req)
	targetEntity := s.repository.GetAgentInfo(context.TODO(), nodeId)
	//try to umarshal the string to json ,type is entity.ResponseCollectorRegistration
	var agentConfig entity.ResponseCollectorRegistration
	json.Unmarshal([]byte(targetEntity.NodeDetails), &agentConfig)
	return agentConfig
}

func (s *SidecarServiceImpl) RenderConfiguration(nodeId string, configurationId string) entity.ResponseCollectorConfiguration {
	target := s.repository.GetSidecarTemplateConfig(context.TODO(), configurationId)
	tmpl, _ := template.New("config").Parse(target.ConfigTemplate)
	data := map[string]string{
		"nodeId": nodeId,
	}
	var renderedConfig bytes.Buffer
	tmpl.Execute(&renderedConfig, data)
	result := entity.ResponseCollectorConfiguration{
		ConfigurationId: configurationId,
		Name:            target.Name,
		Template:        renderedConfig.String(),
		BackendId:       target.SidecarBackendId,
	}
	return result
}

func (s *SidecarServiceImpl) ListCollectors() entity.ResponseBackendList {
	entityList := s.repository.ListAgentBackend(context.TODO())

	var result entity.ResponseBackendList
	// init results do not return nil
	result.Backends = make([]entity.ResponseCollectorBackend, 0)

	//convert model to entity, entity.ResponseBackendList
	for _, v := range entityList {
		backend := entity.ResponseCollectorBackend{
			Id:                   v.Id,
			Name:                 v.Name,
			ServiceType:          v.ServiceType,
			OperatingSystem:      v.OperatingSystem,
			ExecutablePath:       v.ExecutablePath,
			ExecuteParameters:    v.ExecuteParameters,
			ValidationParameters: v.ValidationParameters,
		}
		result.Backends = append(result.Backends, backend)
	}
	return result
}

func (s *SidecarServiceImpl) GetServerInfo() entity.ServerVersionResponse {
	return entity.ServerVersionResponse{
		Version:   "5.0.0",
		NodeId:    "master",
		ClusterId: "cluster",
	}
}

func (s *SidecarServiceImpl) GetConfigETag(nodeId string) string {
	agentEntity := s.repository.GetAgentInfo(context.TODO(), nodeId)
	return security.Md5(agentEntity.AgentConfig)
}

func (s *SidecarServiceImpl) GetConfigBackendListETag(backendConfig string) string {
	return security.Md5(backendConfig)
}

func (s *SidecarServiceImpl) GetConfigurationETag(configurationId string) string {
	targetEntity := s.repository.GetSidecarTemplateConfig(context.TODO(), configurationId)
	return security.Md5(targetEntity.ConfigTemplate)
}
