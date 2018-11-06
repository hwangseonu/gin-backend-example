package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var DB *mgo.Database

func init() {
	if s, err := mgo.Dial("mongodb://localhost:27017/backend"); err != nil {
		panic(err)
	} else {
		DB = s.DB("backend")
	}
	initCounters()
}

func initCounters() {
	var idCounter []map[string]interface{}

	if err := DB.C("identitycounters").Find(bson.M{}).All(&idCounter); err != nil || len(idCounter) < 2 {
		DB.C("identitycounters").Insert(map[string]interface{}{
			"count": 0,
			"document": "post",
		})
		DB.C("identitycounters").Insert(map[string]interface{}{
			"count": 0,
			"document": "comment",
		})
	}
}
