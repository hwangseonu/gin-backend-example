package models

import "gopkg.in/mgo.v2"

var DB *mgo.Database

func init() {
	if s, err := mgo.Dial("mongodb://localhost:27017/backend"); err != nil {
		panic(err)
	} else {
		DB = s.DB("backend")
	}
}
