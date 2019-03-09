package tests

import (
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/security"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"testing"
)

func TestSignUp_Success(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	req := &requests.SignUpRequest{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
	}
	_ = models.DeleteUserByUsername(name)
	res, err := DoPost("/users", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.Status)
	log.Println(res.Content)
	_ = models.DeleteUserByUsername(name)
}

func TestSignUp_BadRequest(t *testing.T) {
	name := "tes"
	email := "test"
	req := &requests.SignUpRequest{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
	}
	_ = models.DeleteUserByUsername(name)
	res, err := DoPost("/users", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.Status)
	_ = models.DeleteUserByUsername(name)
}

func TestSignUp_Conflict(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	user := &models.User{Username: name, Password: name, Nickname: name, Email: email, Roles: []string{"ROLE_USER"}}
	req := &requests.SignUpRequest{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
	}
	assert.Nil(t, user.Save())
	res, err := DoPost("/users", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusConflict, res.Status)
	_ = models.DeleteUserByUsername(name)
}

func TestGetUser_Success(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	jwt, err := security.GenerateToken(security.ACCESS, name)
	assert.Nil(t, err)
	u := &models.User{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
		Roles: []string{"ROLE_USER"},
	}
	err = u.Save()
	assert.Nil(t, err)
	res, err := DoGetWithJwt("/users", jwt)
	assert.Nil(t, err)
	log.Println(res.Content)
	assert.Equal(t, http.StatusOK, res.Status)
	_ = models.DeleteUserByUsername(name)
}

func TestGetUser_UnprocessableEntity(t *testing.T) {
	name := "test1234"
	jwt, err := security.GenerateToken(security.ACCESS, name)
	assert.Nil(t, err)
	_ = models.DeleteUserByUsername(name)
	res, err := DoGetWithJwt("/users", jwt)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnprocessableEntity, res.Status)
}