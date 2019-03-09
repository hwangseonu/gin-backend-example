package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/responses"
	"net/http"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`)

func SignUp(c *gin.Context) {
	body, _ := c.Get("body")
	req := body.(*requests.SignUpRequest)

	if models.ExistsUserByUsernameOrNicknameOrEmail(req.Username, req.Nickname, req.Email) {
		c.JSON(http.StatusConflict, gin.H{"message": "user already exists"})
		return
	}

	if len(req.Username) < 4 || len(req.Password) < 8 || len(req.Nickname) <= 0 || !emailRegex.MatchString(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return
	}

	u := models.NewUser(req.Username, req.Password, req.Nickname, req.Email, "ROLE_USER")
	if err := u.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, responses.UserResponse{
		Username: u.Username,
		Nickname: u.Nickname,
		Email:    u.Email,
	})
	return
}

func GetUser(c *gin.Context) {
	u, _ := c.Get("user")
	user := u.(*models.User)

	c.JSON(http.StatusOK, responses.UserResponse{
		Username: user.Username,
		Nickname: user.Nickname,
		Email:    user.Email,
	})
}