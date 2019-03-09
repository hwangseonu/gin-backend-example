package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var db *mgo.Database
var users *mgo.Collection
var posts *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	db = s.DB("backend")
	users = db.C("users")
	posts = db.C("posts")
}

func GetNextId(doc string) int {
	var counter map[string]interface{}

	if err := db.C("auto_increment").Find(bson.M{"document": doc}).One(&counter); err != nil || counter == nil {
		_ = db.C("auto_increment").Insert(map[string]interface{}{
			"count":    1,
			"document": doc,
		})
		return 0
	} else {
		id := counter["count"].(int)
		counter["count"] = id + 1
		_ = db.C("auto_increment").Update(bson.M{"document": doc}, counter)
		return id
	}
}

