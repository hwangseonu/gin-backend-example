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

func (p *Post) Save() error {
	 _, err := posts.Upsert(bson.M{"_id": p.Id}, p)
	 return err
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

func FindPostById(id int) *Post {
	var post *Post
	if err := posts.Find(bson.M{"_id": id}).One(&post); err != nil {
		return nil
	}
	return post
}

func DeletePostById(id int) error {
	return users.Remove(bson.M{"_id": id})
}