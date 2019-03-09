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
	"strconv"
	"testing"
)

const name = "test1234"
var user *models.User
const title = "test_title"
var post *models.Post

func Before() {
	user = &models.User{
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

func SavePost() error {
	post = models.NewPost(title, title, user)
	return post.Save()
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

func TestGetPost_Success(t *testing.T) {
	Before()
	assert.Nil(t, SavePost())
	res, err := DoGet("/posts/" + strconv.Itoa(post.Id))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Status)
	After()
}

func TestGetPost_NotFound(t *testing.T) {
	Before()
	assert.Nil(t, SavePost())
	res, err := DoGet("/posts/" + strconv.Itoa(post.Id + 1))
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.Status)
	After()
}

func TestGetPost_BadRequest(t *testing.T) {
	res, err := DoGet("/posts/abcde")
	assert.Nil(t, err)
	assert.Equal(t, http.StatusBadRequest, res.Status)
}