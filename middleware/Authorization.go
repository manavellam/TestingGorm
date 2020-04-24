package middleware

import (
	"log"

	"github.com/TestingGorm/models"
	"github.com/gin-gonic/gin"

	jwt "github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret []byte = []byte("Supersecretkey2")

//Authorization checks whether there is a token or not, and  if it is valid or not
func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			token, err := jwt.ParseWithClaims(c.GetHeader("Authorization"), &models.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return hmacSampleSecret, nil
			})
			if err != nil {
				log.Print("AUTH ERR: ", err)
				AbortWithStatusAndMessage(c, "Token expired or invalid, please log in again.", 418)
				return
			}

			if claims, ok := token.Claims.(*models.MyCustomClaims); ok && token.Valid {
				log.Print("Read User From Token: ", claims.Subject)
			}
		} else {
			AbortWithStatusAndMessage(c, "No token found. Please log in again.", 401)
		}
	}
}
