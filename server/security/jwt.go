package security

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"time"
)

type CustomClaims struct {
	jwt.StandardClaims
	Identity string `json:"identity"`
}

func (c CustomClaims) Valid() error {
	if err := c.StandardClaims.Valid(); err != nil {
		return err
	}
	if models.ExistsByUsername(c.Identity){
		return errors.New("cannot find user by username")
	}
	return nil
}

func GenerateToken(t, username string) (string, error) {
	var expire int64

	if t == "access" {
		expire = time.Now().Add(time.Hour).Unix()
	} else {
		expire = time.Now().AddDate(0, 1, 0).Unix()
	}

	claims := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: expire,
		IssuedAt:  time.Now().Unix(),
		Issuer:    "",
		Subject: t,
		NotBefore: time.Now().Unix(),
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS512, CustomClaims{claims, username}).SignedString([]byte(secret))
}
