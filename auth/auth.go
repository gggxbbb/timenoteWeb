package auth

import (
	"github.com/gin-gonic/gin"
	"timenoteWeb/config"
)

func BasicAuth(c *gin.Context, config *config.Config) bool {
	username, password, ok := c.Request.BasicAuth()
	if !ok || username != config.Username || password != config.Password {
		c.AbortWithStatus(401)
		return false
	}
	return true
}

func BasicAuthFunc(config *config.Config) gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		config.Username: config.Password,
	})
}
