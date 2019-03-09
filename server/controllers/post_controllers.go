package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/responses"
	"net/http"
	"strconv"
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
	println(id)
	if post := models.FindPostById(id); post == nil {
		c.Status(http.StatusNotFound)
		return
	} else {
		if user := models.FindUserById(post.Writer); user == nil {
			if err := models.DeletePostById(post.Id); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}
			c.Status(http.StatusGone)
			return
		} else {
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
}
