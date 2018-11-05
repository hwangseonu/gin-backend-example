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

func CreateUser(username, password, nickname, email string) (*User, error) {
	if u, _ := GetUser(username, nickname, email); u != nil {
		return nil, errors.New("user already exists")
	}

	u := &User{username, password, nickname, email, "ROLE_USER"}

	if err := DB.C("users").Insert(u); err != nil {
		return nil, err
	}
	return u, nil
}

func GetUser(username, nickname, email string) (*User, error) {
	var u *User
	q := []bson.M{{"username": username}, {"nickname": nickname}, {"email": email}}
	if err := DB.C("users").Find(bson.M{"$or": q}).One(&u); err != nil {
		return nil, err
	}
	if u == nil {
		return nil, errors.New("could not find user")
	}
	return u, nil
}