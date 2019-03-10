package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/responses"
	"net/http"
	"strconv"
)

func AddComment(c *gin.Context) {
	body, _ := c.Get("body")
	u, _ := c.Get("user")
	req := body.(*requests.AddCommentRequest)
	user := u.(*models.User)

	param := c.Param("post_id")
	if id, err := strconv.Atoi(param); err != nil {
		c.Status(http.StatusBadRequest)
		return
	} else {
		if post := models.FindPostById(id); post == nil {
			c.Status(http.StatusNotFound)
			return
		} else {
			comment := models.NewComment(req.Content, user)
			post.AddComment(comment)
			if err := post.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusCreated, responses.NewPostResponse(post))
			return
		}
	}
}
