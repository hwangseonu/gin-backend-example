package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/jwt"
	"github.com/hwangseonu/gin-backend/models"
)

func Auth(c *gin.Context) {
	var payload struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil{
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	u, err := models.GetUser(payload.Username)
	if err != nil {
		c.JSON(404, gin.H{"message": err.Error()})
		return
	} else if !u.Verify(payload.Password) {
		c.JSON(422, gin.H{"message": "incorrect password"})
		return
	}

	access, err1 := jwt.GenerateToken("access", "hwangseonu12")
	refresh, err2 := jwt.GenerateToken("refresh", "hwangseonu12")

	if err1 != nil || err2 != nil {
		c.JSON(500, gin.H{
			"access_error": err1.Error(),
			"refresh_error": err2.Error(),
		})
	} else {
		c.JSON(200, gin.H{"access": access, "refresh": refresh})
	}
}

