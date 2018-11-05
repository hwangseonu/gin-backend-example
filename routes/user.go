package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/models"
)

func SignUp(c *gin.Context) {
	var payload struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
		Nickname string `json:"nickname" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil{
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if _, err := models.CreateUser(payload.Username, payload.Password, payload.Nickname, payload.Email); err != nil {
		c.JSON(409, gin.H{"message": err.Error()})
		return
	}

	c.Status(201)
}