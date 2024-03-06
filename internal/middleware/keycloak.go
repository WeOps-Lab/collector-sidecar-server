package middleware

import (
	"collector-sidecar-server/pkg/config"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type KeyCloakMiddleware struct {
	client       *gocloak.GoCloak
	realm        string
	clientId     string
	clientSecret string
}

func NewKeyCloakMiddleware(config config.KeyCloakConfig) *KeyCloakMiddleware {
	return &KeyCloakMiddleware{
		client:       gocloak.NewClient(config.ServerUrl),
		realm:        config.Realm,
		clientId:     config.ClientId,
		clientSecret: config.ClientSecret,
	}
}

func (receiver KeyCloakMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		if config.GlobalConfig.Mode == "debug" {
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not provided"})
			return
		}
		authArr := strings.Fields(authHeader)
		if len(authArr) != 2 || authArr[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Authorization header format"})
			return
		}
		token := authArr[1]
		rptResult, err := receiver.client.RetrospectToken(c, token, receiver.clientId, receiver.clientSecret, receiver.realm)
		if err != nil || !*rptResult.Active {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	}
}
