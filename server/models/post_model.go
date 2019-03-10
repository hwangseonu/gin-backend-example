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
	Comments []Comment     `json:"comments"`
	CreateAt time.Time     `json:"create_at"`
	UpdateAt time.Time     `json:"update_at"`
}

type Comment struct {
	Id       int           `json:"id" bson:"_id"`
	Content  string        `json:"content"`
	Writer   bson.ObjectId `json:"writer"`
	CreateAt time.Time     `json:"create_at"`
	UpdateAt time.Time     `json:"update_at"`
}

func (p *Post) Save() error {
	 _, err := posts.Upsert(bson.M{"_id": p.Id}, p)
	 return err
}

func (p *Post) AddComment(comment Comment) {
	p.Comments = append(p.Comments, comment)
}

func NewPost(title, content string, writer *User) *Post {
	return &Post{
		Id:       GetNextId("posts"),
		Title:    title,
		Content:  content,
		Writer:   writer.Id,
		Comments: make([]Comment, 0),
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}

func NewComment(content string, writer *User) Comment {
	return Comment{
		Id:       GetNextId("comments"),
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
	for i, c := range post.Comments {
		if user := FindUserById(c.Writer); user == nil {
			post.Comments = append(post.Comments[:i], post.Comments[i+1:]...)
			if err := post.Save(); err != nil {
				return nil
			}
		}
	}
	return post
}

func DeletePostById(id int) error {
	return posts.RemoveId(id)
}

func ExistsPostById(id int) bool {
	return FindPostById(id) != nil
}