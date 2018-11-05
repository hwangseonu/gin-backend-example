package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type Post struct {
	ID       uint32        `json:"id" bson:"_id"`
	Title    string        `json:"title" bson:"title"`
	Content  string        `json:"content" bson:"content"`
	Author   bson.ObjectId `json:"author" bson:"author"`
	CreateAt time.Time     `json:"create_at" bson:"create_at"`
	UpdateAt time.Time     `json:"update_at" bson:"update_at"`
	Comments []Comment     `json:"comments" bson:"comments"`
}

type Comment struct {
	ID       uint32        `json:"id" bson:"_id"`
	Content  string        `json:"content" bson:"content"`
	Author   bson.ObjectId `json:"author" bson:"author"`
	CreateAt time.Time     `json:"create_at" bson:"create_at"`
	UpdateAt time.Time     `json:"update_at" bson:"update_at"`
}

func CreatePost(author *User, title, content string) (*Post, error) {
	var posts []*Post
	var id uint32
	if err := DB.C("posts").Find(bson.M{}).All(&posts); err != nil || len(posts) == 0 {
		id = 0
	} else {
		id = posts[len(posts) - 1].ID + 1
	}
	post := &Post{id, title, content, author.ID, time.Now(), time.Now(), []Comment{}}
	if err := DB.C("posts").Insert(post); err != nil {
		return nil, err
	}
	return post, nil
}