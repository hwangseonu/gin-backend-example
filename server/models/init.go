package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

var db *mgo.Database
var users *mgo.Collection

func init() {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	db = s.DB("backend")
	users = db.C("users")
}
