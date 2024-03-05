package service

import (
	"bytes"
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/model"
	"collector-sidecar-server/pkg/security"
	"github.com/acmestack/gorm-plus/gplus"
	"github.com/gookit/goutil"
	"text/template"
)

type SidecarServiceImpl struct {
}

func NewSidecarService() *SidecarServiceImpl {
	return &SidecarServiceImpl{}
}

func (s *SidecarServiceImpl) UpdateSidecarNodeInfo(nodeId string, req entity.RegistrationSidecarEntity) entity.CollectorRegistrationEntity {
	query, u := gplus.NewQuery[model.SidecarAgentInfoModel]()
	query.Eq(&u.NodeId, nodeId)
	obj, _ := gplus.SelectOne(query)
	obj.NodeDetails = req.NodeDetails
	gplus.UpdateById[model.SidecarAgentInfoModel](obj)

	return entity.CollectorRegistrationEntity{
		Configuration:         obj.AgentConfig.Configuration,
		ConfigurationOverride: obj.AgentConfig.ConfigurationOverride,
		CollectorActions:      obj.AgentConfig.CollectorActions,
		Assignments:           obj.AgentConfig.Assignments,
	}
}

func (s *SidecarServiceImpl) RenderConfiguration(nodeId string, configurationId uint) entity.CollectorConfigurationEntity {
	templateObj, _ := gplus.SelectById[model.SidecarTemplateConfigModel](configurationId)
	tmpl, _ := template.New("config").Parse(templateObj.ConfigTemplate)
	data := map[string]string{
		"nodeId": nodeId,
	}
	var renderedConfig bytes.Buffer
	tmpl.Execute(&renderedConfig, data)
	result := entity.CollectorConfigurationEntity{
		ConfigurationId: goutil.String(configurationId),
		Name:            templateObj.Name,
		Template:        renderedConfig.String(),
		BackendId:       templateObj.SidecarBackendID,
	}
	return result
}

func (s *SidecarServiceImpl) ListCollectors() entity.BackendListEntity {
	resultList, _ := gplus.SelectList[model.SidecarBackendModel](nil)

	var result entity.BackendListEntity
	// init results do not return nil
	result.Backends = make([]entity.CollectorBackendEntity, 0)

	//convert model to entity, entity.BackendListEntity
	for _, v := range resultList {
		backend := entity.CollectorBackendEntity{
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

func (s *SidecarServiceImpl) GetServerInfo() entity.ServerVersionEntity {
	return entity.ServerVersionEntity{
		Version:   "5.0.0",
		NodeId:    "master",
		ClusterId: "cluster",
	}
}

func (s *SidecarServiceImpl) GetConfigETag(nodeId string) string {
	query, u := gplus.NewQuery[model.SidecarAgentInfoModel]()
	query.Eq(&u.NodeId, nodeId)
	obj, _ := gplus.SelectOne(query)

	return security.Md5(goutil.String(obj.AgentConfig))
	return ""
}

func (s *SidecarServiceImpl) GetConfigBackendListETag(backendConfig string) string {
	return security.Md5(backendConfig)
}

func (s *SidecarServiceImpl) GetConfigurationETag(configurationId uint) string {
	obj, _ := gplus.SelectById[model.SidecarTemplateConfigModel](configurationId)
	return security.Md5(goutil.String(obj.ConfigTemplate))
}
