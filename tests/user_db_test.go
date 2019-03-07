package tests

import (
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestSave_Success(t *testing.T) {
	name := "test"
	email := "test@email.com"
	u := models.NewUser(name, name, name, email, "ROLE_USER")
	err := u.Save()
	assert.Nil(t, err)
	err = session.DB("backend").C("users").Find(bson.M{"username": name}).One(&u);
	assert.Nil(t, err)
	assert.Equal(t, name, u.Username)
	assert.Equal(t, name, u.Nickname)
	assert.Equal(t, email, u.Email)
	assert.Len(t, u.Roles, 1)
	_ = models.DeleteByUsername("test")
}