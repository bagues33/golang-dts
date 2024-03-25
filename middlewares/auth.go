package middlewares

import (
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var JwtKey = []byte("rahasia")

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		getHeader := c.GetHeader("Authorization")

		split := strings.Split(getHeader, "Bearer ")

		errInvalidToken := errors.New("invalid token")

		if len(split) != 2 {
			c.AbortWithStatusJSON(401, gin.H{
				"message": errInvalidToken.Error(),
			})
			return
		}
		getToken := split[1]
		log.Println("getToken", getToken)
		validated, err := jwt.Parse(getToken, func(token *jwt.Token) (interface{}, error) {

			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errInvalidToken
			}

			return []byte(JwtKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{
				"message": errInvalidToken.Error(),
			})
			return
		}

		if _, ok := validated.Claims.(jwt.MapClaims); !ok && !validated.Valid {
			c.AbortWithStatusJSON(401, gin.H{
				"message": errInvalidToken.Error(),
			})

			return
		}

		c.Set("user", validated.Claims.(jwt.MapClaims))

		log.Println("user", validated.Claims.(jwt.MapClaims))

		c.Next()
	}
}
