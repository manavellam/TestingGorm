package models

import "github.com/dgrijalva/jwt-go"

//MyCustomClaims is mode for claims
type MyCustomClaims struct {
	//possible to add fields here
	jwt.StandardClaims
}
