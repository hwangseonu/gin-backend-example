package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	custom "github.com/hwangseonu/gin-backend/jwt"
	"github.com/hwangseonu/gin-backend/models"
	"strings"
)

func AuthRequired(sub string) gin.HandlerFunc {
	return func (c *gin.Context) {
		tokenString := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &custom.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			c.AbortWithStatusJSON(422, gin.H{"message": err.Error()})
			return
		}
		claims := token.Claims.(*custom.CustomClaims)

		if claims.Subject != sub {
			c.AbortWithStatusJSON(422, gin.H{"message": "jwt subject must " + sub})
			return
		}

		user, _ := models.GetUser(claims.Identity)
		c.Set("user", user)
		c.Next()
	}
}
