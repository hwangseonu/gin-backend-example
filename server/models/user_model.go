package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Username string        `json:"username"`
	Password string        `json:"password"`
	Nickname string        `json:"nickname"`
	Email    string        `json:"email"`
	Roles    []string      `json:"roles"`
}

func NewUser(username, password, nickname, email string, roles ...string) *User {
	return &User{
		Id: bson.NewObjectId(),
		Username: username,
		Password: password,
		Nickname: nickname,
		Email:    email,
		Roles:    roles,
	}
}

func (u *User) Save() error {
	_, err := users.Upsert(bson.M{"username": u.Username}, u)
	return err
}

func (u *User) Equals(t *User) bool {
	return u.Id.String() == t.Id.String()
}

func FindUserByUsername(username string) *User {
	var user *User
	if err := users.Find(bson.M{"username": username}).One(&user); err != nil {
		return nil
	}
	return user
}

func FindUserById(id bson.ObjectId) *User {
	var user *User
	if err := users.FindId(id).One(&user); err != nil {
		return nil
	}
	return user
}

func ExistsUserByUsername(username string) bool {
	var user *User
	if err := users.Find(bson.M{username: username}).One(&user); err != nil {
		return false
	}
	return user != nil
}

func ExistsUserByUsernameOrNicknameOrEmail(username, nickname, email string) bool {
	var user *User
	if err := users.Find(bson.M{"$or": []bson.M{ {"username": username}, {"nickname": nickname}, {"email": email} }}).One(&user); err != nil {
		return false
	}
	return user != nil
}

func DeleteUserByUsername(username string) error {
	return users.Remove(bson.M{"username": username})
}