package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username" bson:"username"`
	Password string        `json:"password" bson:"password"`
	Nickname string        `json:"nickname" bson:"nickname"`
	Email    string        `json:"email" bson:"email"`
	Role     string        `json:"role" bson:"role"`
}

func (u *User) Verify(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) == nil
}

func CreateUser(username, password, nickname, email string) (*User, error) {
	if ExistsUser(username, nickname, email) {
		return nil, errors.New("user already exists")
	}
	encrypt, err := bcrypt.GenerateFromPassword([]byte(password), 0)

	if err != nil {
		return nil, err
	}

	u := &User{bson.NewObjectId(), username, string(encrypt), nickname, email, "ROLE_USER"}

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

func GetUserById(id bson.ObjectId) (*User, error) {
	var u *User
	if err := DB.C("users").FindId(id).One(&u); err != nil {
		return nil, err
	} else if u == nil {
		return nil, errors.New("could not find user by id")
	}
	return u, nil
}