package v1

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil"
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
// @Param current query int false "current"
// @Param size query int false "size"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_agent_info [get]
func (h *SidecarAgentInfoHandler) ListAgentInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		current, size := entity.ExtractPageParam(c)
		result := h.sidecarAgentInfoService.ListAgentInfo(current, size, queryParams)
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
		id := c.Param("node_id")
		result := h.sidecarAgentInfoService.GetAgentInfo(goutil.Uint(id))
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
// @Param req body entity.CollectorRegistrationEntity true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_agent_info/{node_id} [put]
func (h *SidecarAgentInfoHandler) UpdateAgentConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("node_id")
		target := entity.SidecarAgentInfoWrapperEntity{Id: goutil.Uint(id)}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := h.sidecarAgentInfoService.UpdateAgentConfig(target)
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
		id := c.Param("node_id")
		err := h.sidecarAgentInfoService.DeleteAgentInfo(goutil.Uint(id))
		response.JSON(c, err, nil)
	}
}
