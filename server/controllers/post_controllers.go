package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"net/http"
)

func CreatePost(c *gin.Context) {
	body, _ := c.Get("body")
	u, _ := c.Get("user")
	req := body.(*requests.CreatePostRequest)
	user := u.(*models.User)

	post := models.NewPost(req.Title, req.Content, user)
	if err := post.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, post)
	return
}