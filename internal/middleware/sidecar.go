package middleware

import (
	"collector-sidecar-server/pkg/log"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func SidecarAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationInfo := c.Request.Header.Get("Authorization")
		if authorizationInfo == "" {
			c.JSON(401, gin.H{
				"message": "Authorization header is empty",
			})
			c.Abort()
			return
		}

		auth := strings.SplitN(authorizationInfo, " ", 2)
		if len(auth) != 2 || auth[0] != "Basic" {
			c.JSON(401, gin.H{
				"message": "Authorization header is not Bearer",
			})
			c.Abort()
			return
		}
		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)
		if len(pair) != 2 {
			c.JSON(401, gin.H{
				"message": "Authorization header is not valid",
			})
			c.Abort()
			return
		}

		log.Info(fmt.Sprintf("token: %s", pair[0]))
		c.Next()
	}
}
