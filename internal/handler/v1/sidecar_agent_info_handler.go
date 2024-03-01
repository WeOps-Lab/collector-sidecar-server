package v1

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type SidecarAgentInfoHandler struct {
	sidecarAgentInfoService service.SidecarAgentInfoService
}

func NewSidecarAgentInfoHandler(sidecarAgentInfoService service.SidecarAgentInfoService) *SidecarAgentInfoHandler {
	return &SidecarAgentInfoHandler{
		sidecarAgentInfoService: sidecarAgentInfoService,
	}
}

// ListAgentInfo godoc
// @Summary ListAgentInfo
// @Description ListAgentInfo
// Tags: sidecar
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_agent_info [get]
func (h *SidecarAgentInfoHandler) ListAgentInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := h.sidecarAgentInfoService.ListAgentInfo()
		response.JSON(c, nil, result)
	}
}

// GetAgentInfo godoc
// @Summary GetAgentInfo
// @Description GetAgentInfo
// Tags: sidecar
// @Accept json
// @Produce json
// @Param node_id path string true "node_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_agent_info/{node_id} [get]
func (h *SidecarAgentInfoHandler) GetAgentInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeId := c.Param("node_id")
		result := h.sidecarAgentInfoService.GetAgentInfo(nodeId)
		response.JSON(c, nil, result)
	}
}

// UpdateAgentConfig  godoc
// @Summary GetAgentInfo
// @Description GetAgentInfo
// Tags: sidecar
// @Accept json
// @Produce json
// @Param node_id path string true "node_id"
// @Param req body entity.ResponseCollectorRegistration true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_agent_info/{node_id} [put]
func (h *SidecarAgentInfoHandler) UpdateAgentConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeId := c.Param("node_id")
		target := entity.SidecarAgentInfo{}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := h.sidecarAgentInfoService.UpdateAgentConfig(nodeId, target)
		response.JSON(c, err, nil)
	}

}

// DeleteAgentInfo godoc
// @Summary DeleteAgentInfo
// @Description DeleteAgentInfo
// Tags: sidecar
// @Accept json
// @Produce json
// @Param node_id path string true "node_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_agent_info/{node_id} [delete]
func (h *SidecarAgentInfoHandler) DeleteAgentInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeId := c.Param("node_id")
		err := h.sidecarAgentInfoService.DeleteAgentInfo(nodeId)
		response.JSON(c, err, nil)
	}
}
