package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// miss router
func MissRouter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, nil)
	}
}
