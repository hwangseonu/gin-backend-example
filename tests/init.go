package tests

import (
	"encoding/json"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var session *mgo.Session

func init() {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
	}
	session = s
}

type Response struct {
	Content string
	Status  int
}

func DoPost(url string, body interface{}) (*Response, error) {
	var b []byte
	var res *http.Response
	var err error
	if b, err = json.MarshalIndent(body, "", "  "); err != nil {
		return nil, err
	}
	if res, err = http.Post(url, "application/json", strings.NewReader(string(b))); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if b, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	return &Response{Content:string(b), Status:res.StatusCode}, nil
}
