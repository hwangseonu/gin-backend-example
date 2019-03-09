package tests

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/requests"
	"github.com/hwangseonu/gin-backend-example/server/security"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestSignIn_Success(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	user := &models.User{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
		Roles: []string{"ROLE_USER"},
	}
	req := &requests.SignInRequest{
		Username: name,
		Password: name,
	}
	_ = user.Save()
	res, err := DoPost("/auth", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Status)

	var resp map[string]string
	err = json.Unmarshal([]byte(res.Content), &resp)
	assert.Nil(t, err)

	log.Printf("access: %s\n", resp["access"])
	log.Printf("refresh: %s\n", resp["refresh"])

	_ = models.DeleteUserByUsername(name)
}

func TestSignIn_NotFound(t *testing.T) {
	name := "test1234"
	req := &requests.SignInRequest{
		Username: name,
		Password: name,
	}
	_ = models.DeleteUserByUsername(name)
	res, err := DoPost("/auth", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusNotFound, res.Status)
	_ = models.DeleteUserByUsername(name)
}

func TestSignIn_Unauthorized(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	user := &models.User{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
		Roles: []string{"ROLE_USER"},
	}
	req := &requests.SignInRequest{
		Username: name,
		Password: "not password",
	}
	_ = user.Save()
	res, err := DoPost("/auth", req)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusUnauthorized, res.Status)
	_ = models.DeleteUserByUsername(name)
}

func TestRefresh_Success_ReturnOnlyAccess(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	jwt, err := security.GenerateToken(security.REFRESH, name)
	assert.Nil(t, err)
	user := &models.User{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
		Roles: []string{"ROLE_USER"},
	}
	_ = user.Save()
	res, err := DoGetWithJwt("/auth/refresh", jwt)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Status)

	var resp map[string]string
	err = json.Unmarshal([]byte(res.Content), &resp)
	assert.Nil(t, err)

	log.Println("access: " + resp["access"])
	_, ok := resp["refresh"]
	assert.False(t, ok)

	_ = models.DeleteUserByUsername(name)
}

func TestRefresh_Success_ReturnAccessAndRefresh(t *testing.T) {
	name := "test1234"
	email := "test@email.com"
	jwt, err := GenerateTestToken(security.REFRESH, name, time.Now().AddDate(0, 0, 6).Unix())
	assert.Nil(t, err)
	user := &models.User{
		Username: name,
		Password: name,
		Nickname: name,
		Email: email,
		Roles: []string{"ROLE_USER"},
	}
	_ = user.Save()
	res, err := DoGetWithJwt("/auth/refresh", jwt)
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.Status)

	var resp map[string]string
	err = json.Unmarshal([]byte(res.Content), &resp)
	assert.Nil(t, err)

	log.Println("access: " + resp["access"])
	log.Println("refresh: ", resp["refresh"])

	_ = models.DeleteUserByUsername(name)
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