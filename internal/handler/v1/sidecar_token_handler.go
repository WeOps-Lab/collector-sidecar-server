package v1

import (
	"collector-sidecar-server/internal/entity"
	"collector-sidecar-server/internal/service"
	"collector-sidecar-server/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil"
)

type SidecarTokenHandler struct {
	sidecarTokenService service.SidecarTokenService
}

func NewSidecarTokenHandler(sidecarTokenService service.SidecarTokenService) *SidecarTokenHandler {
	return &SidecarTokenHandler{
		sidecarTokenService: sidecarTokenService,
	}
}

// List godoc
// @Schemes
// @Summary List
// @Description List
// @Accept json
// @Produce json
// @Param current query int false "current"
// @Param size query int false "size"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_token/ [get]
func (h *SidecarTokenHandler) List() gin.HandlerFunc {
	return func(c *gin.Context) {
		queryParams := c.Request.URL.Query()
		current, size := entity.ExtractPageParam(c)
		results := h.sidecarTokenService.List(current, size, queryParams)
		response.JSON(c, nil, results)
	}
}

// GetEntity godoc
// @Summary GetEntity
// @Description GetEntity
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_template_config/{id} [get]
func (handler *SidecarTokenHandler) GetEntity() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		result := handler.sidecarTokenService.Get(goutil.Uint(id))
		response.JSON(c, nil, result)
	}
}

// Create godoc
// @Summary Create
// @Description Create
// @Accept json
// @Produce json
// @Param req body entity.SidecarTokenEntity true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_token/ [post]
func (handler *SidecarTokenHandler) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		target := entity.SidecarTokenEntity{}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := handler.sidecarTokenService.Create(target)
		response.JSON(c, err, nil)
	}
}

// Update godoc
// @Summary Update
// @Description Update
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param req body entity.SidecarTokenWrapperEntity true "请求体"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_token/{id} [put]
func (handler *SidecarTokenHandler) Update() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		target := entity.SidecarTokenWrapperEntity{
			Id: goutil.Uint(id),
		}
		if err := c.ShouldBindJSON(&target); err != nil {
			response.JSON(c, err, nil)
			return
		}
		err := handler.sidecarTokenService.Update(target)
		response.JSON(c, err, nil)
	}
}

// DeleteEntity godoc
// @Summary DeleteEntity
// @Description DeleteEntity
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} response.ApiResponse
// @Router /api/sidecar_token/{id} [delete]
func (handler *SidecarTokenHandler) DeleteEntity() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		err := handler.sidecarTokenService.Delete(goutil.Uint(id))
		response.JSON(c, err, nil)
	}
}
