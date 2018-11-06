package models

import (
	"errors"
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
	post := &Post{uint32(getNextId("post")), title, content, author.ID, time.Now(), time.Now(), []Comment{}}
	if err := DB.C("posts").Insert(post); err != nil {
		return nil, err
	}
	return post, nil
}

func GetPostById(id uint32) (*Post, error) {
	var post *Post
	if err := DB.C("posts").FindId(id).One(&post); err != nil {
		return nil, err
	} else if post == nil {
		return nil, errors.New("could not find post by id")
	}
	return post, nil
}

func getNextId(doc string) int {
	var counter map[string]interface{}
	DB.C("identitycounters").Find(bson.M{"document": doc}).One(&counter)
	id := counter["count"].(int)
	counter["count"] = id + 1
	DB.C("identitycounters").Update(bson.M{"document": doc}, counter)
	return id
}