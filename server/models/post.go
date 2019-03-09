package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	Id       int           `json:"id" bson:"_id"`
	Title    string        `json:"title"`
	Content  string        `json:"content"`
	Writer   bson.ObjectId `json:"writer"`
	CreateAt time.Time     `json:"create_at"`
	UpdateAt time.Time     `json:"update_at"`
}

func NewPost(title, content string, writer *User, createAt, updateAt time.Time) *Post {
	return &Post{
		Id: GetNextId("posts"),
		Title:    title,
		Content:  content,
		Writer:   writer.Id,
		CreateAt: createAt,
		UpdateAt: updateAt,
	}
}
