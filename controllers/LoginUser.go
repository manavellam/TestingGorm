package controllers

import (
	"github/TestingGorm/middleware"
	"github/TestingGorm/models"

	"github.com/gin-gonic/gin"
)

//LoginUser checks that the user is not logged in (Has no valid token) and, if not, check if the U+P pair is valid.
func LoginUser(c *gin.Context) {
	var (
		loguser models.User
	)
	token := c.GetHeader("Authorization")
	//Check if already logged in.
	if !middleware.IsTknValid(&token) {

		c.Bind(&loguser)
		if middleware.IsUser(&loguser) {
			//Generates Token if valid User info.
			middleware.TokenGen(&token)
			c.Header("Authorization", token)
		} else {
			c.Writer.Write([]byte("User does not exist"))
			c.Status(401)
		}
	}
	//REDIRECCIONAR A HOMEPAGE SI TIENE TOKEN VALIDO
}
