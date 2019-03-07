package tests

import (
	"github.com/hwangseonu/gin-backend-example/server/models"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestSave_Success(t *testing.T) {
	name := "test"
	email := "test@email.com"
	u := models.NewUser(name, name, name, email, "ROLE_USER")
	err := u.Save()
	if err != nil {
		t.Error(err)
	}
	if err = session.DB("backend").C("users").Find(bson.M{"username": name}).One(&u); err != nil {
		t.Error(err)
	} else {
		if u.Username != name || u.Nickname != name || u.Email != email || len(u.Roles) <= 0{
			t.Fail()
		}
	}
	_ = models.DeleteByUsername("test")
}