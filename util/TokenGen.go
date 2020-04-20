package util

import (
	"log"
	"time"

	"github.com/TestingGorm/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var hmacSampleSecret []byte = []byte("Supersecretkey2")

//TokenGen generates a token for a valid user
func TokenGen(t *string, u *models.User) {

	claims := models.MyCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 1).Unix(),
			Issuer:    "test",
			Subject:   HashString(u.Name),
		},
	}

	log.Print("Hashed user: ", HashString(u.Name))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var err error
	*t, err = token.SignedString(hmacSampleSecret)
	if err != nil {
		log.Print("Error generating token")
	}
}
