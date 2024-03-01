package v1

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type SidecarTemplateConfigHandler struct {
	sidecarTemplateConfigService service.SidecarTemplateConfigService
}

func NewSidecarTemplateConfigHandler(sidecarTemplateConfigService service.SidecarTemplateConfigService) *SidecarTemplateConfigHandler {
	return &SidecarTemplateConfigHandler{
		sidecarTemplateConfigService: sidecarTemplateConfigService,
	}
}

// ListTemplateConfigs godoc
// @Summary List all template configurations
// @Description List all template configurations
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_template_config [get]
func (h *SidecarTemplateConfigHandler) ListTemplateConfigs() gin.HandlerFunc {
	return func(c *gin.Context) {
		results := h.sidecarTemplateConfigService.ListTemplateConfigs()
		response.JSON(c, nil, results)
	}
}

// GetTemplateConfig godoc
// @Summary Get a template configuration
// @Description Get a template configuration
// @Produce json
// @Param template_id path string true "template_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_template_config/{template_id} [get]
func (handler *SidecarTemplateConfigHandler) GetTemplateConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		template_id := c.Param("template_id")
		result := handler.sidecarTemplateConfigService.GetTemplateConfig(template_id)
		response.JSON(c, nil, result)
	}
}

// CreateTemplateConfig godoc
// @Summary Create a template configuration
// @Description Create a template configuration
// @Accept json
// @Produce json
// @Param req body entity.SidecarTemplateConfig true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_template_config/ [post]
func (handler *SidecarTemplateConfigHandler) CreateTemplateConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := entity.SidecarTemplateConfig{}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := handler.sidecarTemplateConfigService.CreateTemplateConfig(target)
		response.JSON(c, err, nil)
	}
}

// UpdateTemplateConfig godoc
// @Summary Update a template configuration
// @Description Update a template configuration
// @Accept json
// @Produce json
// @Param template_id path string true "template_id"
// @Param req body entity.SidecarTemplateConfig true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_template_config/{template_id} [put]
func (handler *SidecarTemplateConfigHandler) UpdateTemplateConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		template_id := c.Param("template_id")
		target := entity.SidecarTemplateConfig{}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := handler.sidecarTemplateConfigService.UpdateTemplateConfig(template_id, target)
		response.JSON(c, err, nil)
	}
}

// DeleteTemplateConfig godoc
// @Summary Delete a template configuration
// @Description Delete a template configuration
// @Produce json
// @Param template_id path string true "template_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_template_config/{template_id} [delete]
func (handler *SidecarTemplateConfigHandler) DeleteTemplateConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		template_id := c.Param("template_id")
		err := handler.sidecarTemplateConfigService.DeleteTemplateConfig(template_id)
		response.JSON(c, err, nil)
	}
}
