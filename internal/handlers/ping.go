package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping is simple keep-alive/ping handler
func Ping() gin.HandlerFunc {
	return func(c *gin.Contect) {
		c.String(http.StatusOK, "OK")
	}
}
