package tests

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	s "github.com/hwangseonu/gin-backend-example/server"
	"github.com/hwangseonu/gin-backend-example/server/security"
	"gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"
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

func DoGetWithJwt(url, jwt string) (*Response, error) {
	var req *http.Request
	var err error

	if req, err = http.NewRequest(http.MethodGet, server.URL + url, nil); err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + jwt)

	return DoRequest(req)
}

func DoPostWithJwt(url, name string, body interface{}) (*Response, error) {
	var b []byte
	var err error
	var jwt string
	var req *http.Request

	if b, err = json.MarshalIndent(body, "", "  "); err != nil {
		return nil, err
	}

	if jwt, err = GenerateTestToken(security.ACCESS, name, time.Now().AddDate(0, 0, 1).Unix()); err != nil {
		return nil, err
	}
	if req, err = http.NewRequest(http.MethodPost, server.URL + url, strings.NewReader(string(b))); err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer " + jwt)
	return DoRequest(req)

}

func GenerateTestToken(t, username string, exp int64) (string, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "jwt-secret"
	}
	claims := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: exp,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "",
		Subject: t,
		NotBefore: time.Now().Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS512, security.CustomClaims{StandardClaims: claims, Identity: username}).SignedString([]byte(secret))
}