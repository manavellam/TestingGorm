package controllers

import (
	"github.com/TestingGorm/models"

	"github.com/TestingGorm/middleware"

	"github.com/gin-gonic/gin"
)

//LoginUser checks that the user is not logged in (Has no valid token) and, if not, check if the U+P pair is valid.
func LoginUser(c *gin.Context) {
	var (
		loguser models.User
	)
	token := c.GetHeader("Authorization")
	//Check if already logged in
	if !middleware.IsTknValid(&token) {
		c.Bind(&loguser)
		if middleware.IsUser(&loguser) {
			//Generates Token if valid User info.
			middleware.TokenGen(&token)
		} else {
			c.Writer.Write([]byte("User does not exist"))
			c.Status(401)
		}
	}
	c.Header("Authorization", token)
	//REDIRECCIONAR A HOMEPAGE SI TIENE TOKEN VALIDO
}
