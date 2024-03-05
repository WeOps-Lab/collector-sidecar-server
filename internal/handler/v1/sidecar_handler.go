package v1

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/errors"
	"collector-sidecar-server/pkg/errors/ecode"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil"
	"net/http"
)

type SidecarHandler struct {
	sidecarService service.SidecarService
}

func NewSidecarHandler(sidecarService service.SidecarService) *SidecarHandler {
	return &SidecarHandler{
		sidecarService: sidecarService,
	}
}

// ServerInfo godoc
// @Summary ServerInfo
// @Description ServerInfo
// Tags: sidecar
// @Accept json
// @Produce json
// @Success 200 {object} entity.ServerVersionEntity
// @Router /api [get]
func (h *SidecarHandler) ServerInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		result := h.sidecarService.GetServerInfo()
		c.JSON(http.StatusOK, result)
	}
}

// UpdateSidecarNodeInfo godoc
// @Summary UpdateSidecarNodeInfo
// @Description UpdateSidecarNodeInfo
// Tags: sidecar
// @Accept json
// @Produce json
// @Param node_id path string true "node_id"
// @Param req body entity.RegistrationSidecarEntity true "请求体"
// @Success 200 {object} entity.CollectorRegistrationEntity
// @Router /api/sidecars/{node_id} [put]
func (h *SidecarHandler) UpdateSidecarNodeInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		nodeId := c.Param("node_id")
		// get If-None-Match from header
		checkSum := c.GetHeader("If-None-Match")

		req := entity.RegistrationSidecarEntity{}
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, errors.WithCode(ecode.ValidateErr, err.Error()))
		}

		result := h.sidecarService.UpdateSidecarNodeInfo(nodeId, req)
		eTag := h.sidecarService.GetConfigETag(nodeId)

		// set eTag to response Header
		c.Header("ETag", eTag)
		if fmt.Sprintf("\"%s\"", eTag) == checkSum {
			c.JSON(http.StatusNotModified, result)
		} else {
			c.JSON(http.StatusAccepted, result)
		}
	}
}

// ListCollectors godoc
// @Summary ListCollectors
// @Description ListCollectors
// Tags: sidecar
// @Accept json
// @Produce json
// @Success 200 {object} entity.BackendListEntity
// @Router /api/sidecar/collectors [get]
func (h *SidecarHandler) ListCollectors() gin.HandlerFunc {
	return func(c *gin.Context) {
		checkSum := c.GetHeader("If-None-Match")
		result := h.sidecarService.ListCollectors()
		jsonResult, _ := json.Marshal(result)

		eTag := h.sidecarService.GetConfigBackendListETag(string(jsonResult))
		c.Header("ETag", eTag)
		if fmt.Sprintf("\"%s\"", eTag) == checkSum {
			c.JSON(http.StatusNotModified, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}

// GetConfiguration godoc
// @Summary GetConfiguration
// @Description GetConfiguration
// Tags: sidecar
// @Accept json
// @Produce json
// @Param node_id path string true "node_id"
// @Param configuration_id path string true "configuration_id"
// @Success 200 {object} entity.CollectorConfigurationEntity
// @Router /api/sidecar/configurations/render/{node_id}/{configuration_id} [get]
func (h *SidecarHandler) GetConfiguration() gin.HandlerFunc {
	return func(c *gin.Context) {
		checkSum := c.GetHeader("If-None-Match")

		nodeId := c.Param("node_id")
		configurationId := c.Param("configuration_id")

		eTag := h.sidecarService.GetConfigurationETag(goutil.Uint(configurationId))
		result := h.sidecarService.RenderConfiguration(nodeId, goutil.Uint(configurationId))
		c.Header("ETag", eTag)
		if fmt.Sprintf("\"%s\"", eTag) == checkSum {
			c.JSON(http.StatusNotModified, result)
		} else {
			c.JSON(http.StatusOK, result)
		}
	}
}
