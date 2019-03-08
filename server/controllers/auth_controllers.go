package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/responses"
	"github.com/hwangseonu/gin-backend-example/server/security"
	"net/http"
)

func SignIn(c *gin.Context) {
	body, _ := c.Get("body")
	req := body.(*requests.SignInRequest)

	user := models.FindByUsername(req.Username)
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "cannot find user by username"})
	} else if user.Password != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "password is mismatch"})
	}

	access, err1 := security.GenerateToken(security.ACCESS, user.Username)
	refresh, err2 := security.GenerateToken(security.REFRESH, user.Username)
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err1.Error() + "\n" + err2.Error()})
	}
	c.JSON(http.StatusOK, responses.AuthResponse{
		Access:  access,
		Refresh: refresh,
	})
}
