package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend/models"
)

func NewPost(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)

	var payload struct {
		Title   string `json:"title" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	_, err := models.CreatePost(user, payload.Title, payload.Content)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	c.Status(200)
}
