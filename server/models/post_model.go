package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	Id       int       `json:"id" bson:"_id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Writer   bson.ObjectId   `json:"writer"`
	CreateAt time.Time `json:"create_at"`
	UpdateAt time.Time `json:"update_at"`
}

func (p *Post) Save() error {
	 _, err := posts.Upsert(bson.M{"_id": p.Id}, p)
	 return err
}

func NewPost(title, content string, writer *User) *Post {
	return &Post{
		Id: GetNextId("posts"),
		Title:    title,
		Content:  content,
		Writer:   writer.Id,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

func FindPostById(id int) *Post {
	var post *Post
	if err := posts.FindId(id).One(&post); err != nil {
		return nil
	}
	if user := FindUserById(post.Writer); user == nil {
		_ = DeletePostById(post.Id)
		return nil
	}
	return post
}

func DeletePostById(id int) error {
	return posts.RemoveId(id)
}

func ExistsPostById(id int) bool {
	return FindPostById(id) != nil
}