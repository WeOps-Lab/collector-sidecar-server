package middleware

import (
	"collector-sidecar-server/internal/handler/ping"
	"collector-sidecar-server/pkg/errors"
	"collector-sidecar-server/pkg/errors/ecode"
	"collector-sidecar-server/pkg/response"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// middleware 实现Router接口
// 便于服务启动时加载, middleware本质跟handler无区别
type middleware struct {
}

func NewMiddleware() *middleware {
	return &middleware{}
}

// Load 注册中间件和公共路由
func (m *middleware) Load(g *gin.Engine) {
	// 注册中间件
	g.Use(gin.Recovery())
	g.Use(NoCache())
	g.Use(Options())
	g.Use(Secure())

	// 404
	g.NoRoute(func(c *gin.Context) {
		response.JSON(c, errors.WithCode(ecode.NotFoundErr, "404 not found!"), nil)
	})

	g.GET("/ping", ping.Ping())
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
