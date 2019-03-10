package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/responses"
	"net/http"
	"strconv"
	"time"
)

func CreatePost(c *gin.Context) {
	body, _ := c.Get("body")
	u, _ := c.Get("user")
	req := body.(*requests.CreatePostRequest)
	user := u.(*models.User)

	if len(req.Title) <= 0 || len(req.Content) <= 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	post := models.NewPost(req.Title, req.Content, user)
	if err := post.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.PostResponse{
		Id:      post.Id,
		Title:   post.Title,
		Content: post.Content,
		Writer: responses.UserResponse{
			Username: user.Username,
			Nickname: user.Nickname,
			Email:    user.Email,
		},
		CreateAt: post.CreateAt,
		UpdateAt: post.UpdateAt,
	})
	return
}

func GetPost(c *gin.Context) {
	param := c.Param("post_id")
	var id int
	var err error
	if id, err = strconv.Atoi(param); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	if post := models.FindPostById(id); post == nil {
		c.Status(http.StatusNotFound)
		return
	} else {
		user := models.FindUserById(post.Writer)
		c.JSON(http.StatusOK, responses.PostResponse{
			Id:      post.Id,
			Title:   post.Title,
			Content: post.Content,
			Writer: responses.UserResponse{
				Username: user.Username,
				Nickname: user.Nickname,
				Email:    user.Email,
			},
			CreateAt: post.CreateAt,
			UpdateAt: post.UpdateAt,
		})
	}
}

func UpdatePost(c *gin.Context) {
	body, _ := c.Get("body")
	req := body.(*requests.CreatePostRequest)
	u, _ := c.Get("user")
	user := u.(*models.User)

	param := c.Param("post_id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if len(req.Title) <= 0 || len(req.Content) <= 0 {
		c.Status(http.StatusBadRequest)
		return
	}

	if post := models.FindPostById(id); post == nil {
		c.Status(http.StatusNotFound)
		return
	} else {
		writer := models.FindUserById(post.Writer)
		if !user.Equals(writer) {
			c.Status(http.StatusForbidden)
			return
		}

		post.Title = req.Title
		post.Content = req.Content
		post.UpdateAt = time.Now()
		if err = post.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.JSON(http.StatusOK, responses.PostResponse{
			Id:      post.Id,
			Title:   post.Title,
			Content: post.Content,
			Writer: responses.UserResponse{
				Username: writer.Username,
				Nickname: writer.Nickname,
				Email:    writer.Email,
			},
			CreateAt: post.CreateAt,
			UpdateAt: post.UpdateAt,
		})
		return
	}
}

func DeletePost(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)

	param := c.Param("post_id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	if post := models.FindPostById(id); post == nil {
		c.Status(http.StatusNotFound)
		return
	} else {
		writer := models.FindUserById(post.Writer)
		if !user.Equals(writer) {
			c.Status(http.StatusForbidden)
			return
		}
		if err := models.DeletePostById(post.Id); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		c.Status(http.StatusOK)
		return
	}
}