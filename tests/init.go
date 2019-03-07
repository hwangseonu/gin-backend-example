package tests

import (
	"gopkg.in/mgo.v2"
	"log"
)

var session *mgo.Session

func init() {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	session = s
}
