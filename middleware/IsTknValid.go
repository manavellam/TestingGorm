package middleware

import (
	"log"
	"time"

	"github.com/TestingGorm/models"

	jwt "github.com/dgrijalva/jwt-go"
)

//IsTknValid checks whether there is a token or not, and  if it is valid or not
func IsTknValid(t *string) bool {
	if *t != "" {
		token, err := jwt.ParseWithClaims(*t, &models.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(*models.MyCustomClaims); ok && token.Valid {
			log.Print(claims.ExpiresAt, " ", time.Now().Unix())
			if claims.ExpiresAt > time.Now().Unix() {
				return true
			}
		}
		log.Println(err)
	}
	return false
}
