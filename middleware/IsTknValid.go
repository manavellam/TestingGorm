package middleware

import (
	"github/TestingGorm/models"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//IsTknValid checks whether there is a token or not, and  if it is valid or not
func IsTknValid(t *string) bool {
	token, err := jwt.ParseWithClaims(*t, &models.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return hmacSampleSecret, nil
	})

	if claims, ok := token.Claims.(*models.MyCustomClaims); ok && token.Valid {
		if claims.ExpiresAt > time.Now().Unix() {
			return true
		}
	}
	log.Println(err)
	return false
}
