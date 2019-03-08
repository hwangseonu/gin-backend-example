package tests

import (
	s "github.com/hwangseonu/gin-backend-example/server"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server = httptest.NewServer(s.GenerateApp())

func TestSignUp_Success(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	req := &requests.SignUpRequest{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
	}
	_ = models.DeleteByUsername(name)
	res, err := DoPost(server.URL + "/users", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.Status)
	log.Println(res.Content)
	_ = models.DeleteByUsername(name)
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
	_ = models.DeleteByUsername(name)
	res, err := DoPost(server.URL + "/users", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.Status)
	assert.Nil(t, models.DeleteByUsername(name))
}