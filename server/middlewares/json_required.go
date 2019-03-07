package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonRequired(json interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := c.ShouldBindJSON(&json); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Set("body", json)
	}
}
