package controllers

import (
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/TestingGorm/util"

	"github.com/gin-gonic/gin"
)

//LoginUser checks that the user is not logged in (Has no valid token) and, if not, check if the U+P pair is valid.
func LoginUser(c *gin.Context) {
	var (
		token string
	)

	loguser := c.MustGet("UserList").([]models.User)

	if services.ContainsUser(&loguser[0]) {
		//Generates Token if valid User info.
		util.TokenGen(&token, &loguser[0])
		c.Header("Authorization", token)
	} else {
		c.Writer.Write([]byte("User does not exist"))
		c.Status(401)
	}
}
