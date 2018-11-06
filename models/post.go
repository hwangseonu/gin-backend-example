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

func (p *Post) AddComment(author *User, content string) error {
	p.Comments = append(p.Comments, Comment{
		ID:       uint32(GetNextId("comment")),
		Content:  content,
		Author:   author.ID,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	return DB.C("posts").Update(bson.M{"_id": p.ID}, p)
}

func (p *Post) RemoveComment(id uint32) {
	for i, v := range p.Comments {
		if v.ID == id {
			p.Comments = append(p.Comments[:i], p.Comments[i+1:]...)
		}
	}
	DB.C("posts").Update(bson.M{"_id": p.ID}, p)
}

func (p *Post) GetComment(id uint32) *Comment {
	for _, v := range p.Comments {
		if v.ID == id {
			return &v
		}
	}
	return nil
}

func CreatePost(author *User, title, content string) (*Post, error) {
	post := &Post{uint32(GetNextId("post")), title, content, author.ID, time.Now(), time.Now(), []Comment{}}
	if err := DB.C("posts").Insert(post); err != nil {
		return nil, err
	}
	return post, nil
}

func RemovePost(id uint32) error {
	return DB.C("posts").RemoveId(id)
}

func GetPosts() ([]*Post, error) {
	var posts []*Post
	if err := DB.C("posts").Find(bson.M{}).All(&posts); err != nil {
		return nil, err
	}
	return posts, nil
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
