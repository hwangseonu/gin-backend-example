package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"net/http"
)

func SignUp(c *gin.Context) {
	body, _ := c.Get("body")
	req := body.(*requests.SignUpRequest)
	c.String(http.StatusOK, req.Username)
}
