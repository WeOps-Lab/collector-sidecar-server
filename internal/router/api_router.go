package router

import (
	v1 "collector-sidecar-server/internal/handler/v1"
	"github.com/gin-gonic/gin"
)

type ApiRouter struct {
	sidecarHandler               *v1.SidecarHandler
	sidecarAgentInfoHandler      *v1.SidecarAgentInfoHandler
	sidecarBackendHandler        *v1.SidecarBackendHandler
	sidecarTemplateConfigHandler *v1.SidecarTemplateConfigHandler
}

func NewApiRouter(
	sidecarHandler *v1.SidecarHandler,
	sidecarAgentInfoHandler *v1.SidecarAgentInfoHandler,
	sidecarBackendHandler *v1.SidecarBackendHandler,
	sidecarTemplateConfigHandler *v1.SidecarTemplateConfigHandler) *ApiRouter {
	return &ApiRouter{
		sidecarHandler:               sidecarHandler,
		sidecarAgentInfoHandler:      sidecarAgentInfoHandler,
		sidecarBackendHandler:        sidecarBackendHandler,
		sidecarTemplateConfigHandler: sidecarTemplateConfigHandler,
	}
}

func (ar *ApiRouter) Load(g *gin.Engine) {
	//, middleware.SidecarAuthToken()
	sidecarApi := g.Group("/api")
	{
		sidecarApi.GET("", ar.sidecarHandler.ServerInfo())
		sidecarApi.PUT("/sidecars/:node_id", ar.sidecarHandler.UpdateSidecarNodeInfo())
		sidecarApi.GET("/sidecar/collectors", ar.sidecarHandler.ListCollectors())
		sidecarApi.GET("/sidecar/configurations/:node_id/:configuration_id", ar.sidecarHandler.GetConfiguration())

		sidecarAgentInfo := sidecarApi.Group("sidecar_agent_info")
		{
			sidecarAgentInfo.GET("", ar.sidecarAgentInfoHandler.ListAgentInfo())
			sidecarAgentInfo.GET("/:node_id", ar.sidecarAgentInfoHandler.GetAgentInfo())
			sidecarAgentInfo.PUT("/:node_id", ar.sidecarAgentInfoHandler.UpdateAgentConfig())
			sidecarAgentInfo.DELETE("/:node_id", ar.sidecarAgentInfoHandler.DeleteAgentInfo())
		}

		sidecarBackend := sidecarApi.Group("sidecar_backend")
		{
			sidecarBackend.GET("", ar.sidecarBackendHandler.ListBackend())
			sidecarBackend.GET("/:node_id", ar.sidecarBackendHandler.GetBackend())
			sidecarBackend.POST("", ar.sidecarBackendHandler.CreateBackend())
			sidecarBackend.PUT("/:node_id", ar.sidecarBackendHandler.UpdateBackend())
			sidecarBackend.DELETE("/:node_id", ar.sidecarBackendHandler.DeleteBackend())
		}

		sidecarTemplateConfig := sidecarApi.Group("sidecar_template_config")
		{
			sidecarTemplateConfig.GET("", ar.sidecarTemplateConfigHandler.ListTemplateConfigs())
			sidecarTemplateConfig.GET("/:template_id", ar.sidecarTemplateConfigHandler.GetTemplateConfig())
			sidecarTemplateConfig.POST("", ar.sidecarTemplateConfigHandler.CreateTemplateConfig())
			sidecarTemplateConfig.PUT("/:template_id", ar.sidecarTemplateConfigHandler.UpdateTemplateConfig())
			sidecarTemplateConfig.DELETE("/:template_id", ar.sidecarTemplateConfigHandler.DeleteTemplateConfig())
		}
	}

}
