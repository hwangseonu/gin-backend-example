package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/responses"
	"net/http"
)

func SignUp(c *gin.Context) {
	body, _ := c.Get("body")
	req := body.(*requests.SignUpRequest)

	if models.ExistsByUsernameOrNicknameOrEmail(req.Username, req.Nickname, req.Email) {
		c.JSON(http.StatusConflict, gin.H{"message": "user already exists"})
	}

	u := models.NewUser(req.Username, req.Password, req.Nickname, req.Email, "ROLE_USER")
	if err := u.Save(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, responses.UserResponse{
		Username: u.Username,
		Nickname: u.Nickname,
		Email:    u.Email,
	})
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