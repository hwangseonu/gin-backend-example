package tests

import (
	"encoding/json"
	s "github.com/hwangseonu/gin-backend-example/server"
	"github.com/hwangseonu/gin-backend-example/server/security"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
)

var session *mgo.Session
var server = httptest.NewServer(s.GenerateApp())

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
	if res, err = http.Post(server.URL + url, "application/json", strings.NewReader(string(b))); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if b, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	return &Response{Content:string(b), Status:res.StatusCode}, nil
}

func DoGet(url string) (*Response, error) {
	var b []byte
	var res *http.Response
	var err error

	if res, err = http.Get(server.URL + url); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if b, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	return &Response{Content:string(b), Status:res.StatusCode}, nil
}

func DoRequest(req *http.Request) (*Response, error) {
	var b []byte
	var res *http.Response
	var err error
	if res, err = server.Client().Do(req); err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if b, err = ioutil.ReadAll(res.Body); err != nil {
		return nil, err
	}
	return &Response{Content:string(b), Status:res.StatusCode}, nil
}

func DoGetWithJwt(url, username string) (*Response, error) {
	var req *http.Request
	var err error
	var access string

	if req, err = http.NewRequest(http.MethodGet, server.URL + "/users", nil); err != nil {
		return nil, err
	}
	if access, err = security.GenerateToken(security.ACCESS, username); err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + access)

	return DoRequest(req)
}