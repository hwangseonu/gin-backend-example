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

func UpdateComment(c *gin.Context) {
	body, _ := c.Get("body")
	u, _ := c.Get("user")
	req := body.(*requests.AddCommentRequest)
	user := u.(*models.User)

	var postId int
	var commentId int
	var err error
	var post *models.Post

	param := c.Param("post_id")
	if postId, err = strconv.Atoi(param); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	param = c.Param("comment_id")
	if commentId, err = strconv.Atoi(param); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if post = models.FindPostById(postId); post == nil {
		c.Status(http.StatusNotFound)
		return
	}

	for i, com := range post.Comments {
		if com.Id == commentId {
			if com.Writer.String() != user.Id.String() {
				c.Status(http.StatusForbidden)
				return
			}
			com.Content = req.Content
			comments := post.Comments[:i]
			comments = append(comments, com)
			comments = append(comments, post.Comments[i+1:]...)
			if err := post.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, responses.NewPostResponse(post))
			return
		}
	}
	c.Status(http.StatusNotFound)
	return
}

func DeleteComment(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)

	var postId int
	var commentId int
	var err error
	var post *models.Post

	param := c.Param("post_id")
	if postId, err = strconv.Atoi(param); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	param = c.Param("comment_id")
	if commentId, err = strconv.Atoi(param); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if post = models.FindPostById(postId); post == nil {
		c.Status(http.StatusNotFound)
		return
	}

	for i, com := range post.Comments {
		if com.Id == commentId {
			if com.Writer.String() != user.Id.String() {
				c.Status(http.StatusForbidden)
				return
			}
			post.Comments = append(post.Comments[:i], post.Comments[i+1:]...)
			if err := post.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.JSON(http.StatusOK, responses.NewPostResponse(post))
			return
		}
	}
	c.Status(http.StatusNotFound)
	return
}