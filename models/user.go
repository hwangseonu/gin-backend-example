package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

func (u *User) Verify(password string) bool {
	return u.Password == password
}

func CreateUser(username, password, nickname, email string) (*User, error) {
	if ExistsUser(username, nickname, email) {
		return nil, errors.New("user already exists")
	}

	u := &User{username, password, nickname, email, "ROLE_USER"}

	if err := DB.C("users").Insert(u); err != nil {
		return nil, err
	}
	return u, nil
}

func ExistsUser(username, nickname, email string) bool {
	var u *User
	q := []bson.M{{"username": username}, {"nickname": nickname}, {"email": email}}
	if err := DB.C("users").Find(bson.M{"$or": q}).One(&u); err != nil || u == nil {
		return false
	}
	return true
}

func GetUser(username string) (*User, error) {
	var u *User
	if err := DB.C("users").Find(bson.M{"username": username}).One(&u); err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("could not find user")
	}
	return u, nil
}