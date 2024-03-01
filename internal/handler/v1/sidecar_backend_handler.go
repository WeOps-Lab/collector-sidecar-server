package v1

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/response"
	"github.com/gin-gonic/gin"
)

type SidecarBackendHandler struct {
	sidecarBackendService service.SidecarBackendService
}

func NewSidecarBackendHandler(sidecarBackendService service.SidecarBackendService) *SidecarBackendHandler {
	return &SidecarBackendHandler{
		sidecarBackendService: sidecarBackendService,
	}
}

// ListBackends godoc
// @Summary ListBackends
// @Description ListBackends
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend [get]
func (h *SidecarBackendHandler) ListBackends() gin.HandlerFunc {
	return func(c *gin.Context) {
		results := h.sidecarBackendService.ListBackends()
		response.JSON(c, nil, results)
	}
}

// GetBackend godoc
// @Summary GetBackend
// @Description GetBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param backend_id path string true "backend_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend/{backend_id} [get]
func (h *SidecarBackendHandler) GetBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		backendId := c.Param("backend_id")
		result := h.sidecarBackendService.GetBackend(backendId)
		response.JSON(c, nil, result)
	}
}

// CreateBackend godoc
// @Summary CreateBackend
// @Description CreateBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param req body entity.SidecarBackend true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend [post]
func (h *SidecarBackendHandler) CreateBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := entity.SidecarBackend{}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := h.sidecarBackendService.CreateBackend(target)
		response.JSON(c, err, nil)
	}
}

// UpdateBackend godoc
// @Summary UpdateBackend
// @Description UpdateBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param backend_id path string true "backend_id"
// @Param req body entity.SidecarBackend true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend/{backend_id} [put]
func (h *SidecarBackendHandler) UpdateBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		backendId := c.Param("backend_id")
		target := entity.SidecarBackend{}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := h.sidecarBackendService.UpdateBackend(backendId, target)
		response.JSON(c, err, nil)
	}
}

// DeleteBackend godoc
// @Summary DeleteBackend
// @Description DeleteBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param backend_id path string true "backend_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend/{backend_id} [delete]
func (h *SidecarBackendHandler) DeleteBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		backendId := c.Param("backend_id")
		err := h.sidecarBackendService.DeleteBackend(backendId)
		response.JSON(c, err, nil)
	}
}
