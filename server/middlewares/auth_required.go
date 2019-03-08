package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hwangseonu/gin-backend-example/server/models"
	"github.com/hwangseonu/gin-backend-example/server/security"
	"net/http"
	"strings"
)

func AuthRequired(sub string, roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
		if tokenString == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		var token *jwt.Token
		var err error
		if token, err = jwt.ParseWithClaims(tokenString, &security.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		}); err != nil {
			c.AbortWithStatus(http.StatusUnprocessableEntity)
		}
		claims := token.Claims.(*security.CustomClaims)
		if claims.Subject != sub {
			c.AbortWithStatusJSON(422, gin.H{"message": "jwt subject must " + sub})
			return
		}
		u := models.FindByUsername(claims.Identity)
		if u == nil {
			c.AbortWithStatus(http.StatusUnprocessableEntity)
		} else if !containsAll(u.Roles, roles) {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Set("user", u)
	}
}

func containsAll(s []string, e[] string) bool {
	cnt := len(e)
	for _, a := range s {
		for _, b := range e {
			if a == b {
				cnt--
			}
		}
	}
	return cnt == 0
}
