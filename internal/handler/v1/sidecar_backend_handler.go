package v1

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil"
)

type SidecarBackendHandler struct {
	sidecarBackendService service.SidecarBackendService
}

func NewSidecarBackendHandler(sidecarBackendService service.SidecarBackendService) *SidecarBackendHandler {
	return &SidecarBackendHandler{
		sidecarBackendService: sidecarBackendService,
	}
}

// ListBackend godoc
// @Schemes
// @Summary ListBackend
// @Description ListBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param current query int false "current"
// @Param size query int false "size"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend [get]
func (h *SidecarBackendHandler) ListBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		current, size := entity.ExtractPageParam(c)
		results := h.sidecarBackendService.ListBackend(current, size, queryParams)
		response.JSON(c, nil, results)
	}
}

// GetBackend godoc
// @Summary GetBackend
// @Description GetBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param id path string true "node_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend/{node_id} [get]
func (h *SidecarBackendHandler) GetBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("node_id")
		result := h.sidecarBackendService.GetBackend(goutil.Uint(id))
		response.JSON(c, nil, result)
	}
}

// CreateBackend godoc
// @Summary CreateBackend
// @Description CreateBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param req body entity.SidecarBackendEntity true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend [post]
func (h *SidecarBackendHandler) CreateBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := entity.SidecarBackendEntity{}
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
// @Param id path string true "node_id"
// @Param req body entity.SidecarBackendEntity true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend/{node_id} [put]
func (h *SidecarBackendHandler) UpdateBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("node_id")
		target := entity.SidecarBackendWrapperEntity{Id: goutil.Uint(id)}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := h.sidecarBackendService.UpdateBackend(target)
		response.JSON(c, err, nil)
	}
}

// DeleteBackend godoc
// @Summary DeleteBackend
// @Description DeleteBackend
// Tags sidecar_backend
// @Accept json
// @Produce json
// @Param id path string true "node_id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_backend/{node_id} [delete]
func (h *SidecarBackendHandler) DeleteBackend() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("node_id")
		err := h.sidecarBackendService.DeleteBackend(goutil.Uint(id))
		response.JSON(c, err, nil)
	}
}
