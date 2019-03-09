package tests

import (
	"encoding/json"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/responses"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"testing"
)

const name = "test1234"

func Before() {
	user := models.User{
		Id: bson.NewObjectId(),
		Username: name,
		Password: name,
		Nickname: name,
		Email: name + "@email.com",
		Roles: []string{"ROLE_USER"},
	}
	_ = user.Save()
}

func After() {
	_ = models.DeleteUserByUsername(name)
}

func TestCreatPost_Success(t *testing.T) {
	Before()
	title := "test_title"
	content := " test_content"
	req := &requests.CreatePostRequest{
		Title: title,
		Content: content,
	}
	res, err := DoPostWithJwt("/posts", name, req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.Status)
	log.Println(res.Content)
	var resp responses.PostResponse
	err = json.Unmarshal([]byte(res.Content), &resp)
	assert.Nil(t, err)
	assert.Equal(t, title, resp.Title)
	assert.Equal(t, content, resp.Content)
	assert.Equal(t, name, resp.Writer.Username)
	After()
}

func TestCreatPost_BadRequest(t *testing.T) {
	Before()
	title := ""
	content := ""
	req := &requests.CreatePostRequest{
		Title: title,
		Content: content,
	}
	res, err := DoPostWithJwt("/posts", name, req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.Status)
	After()
}
