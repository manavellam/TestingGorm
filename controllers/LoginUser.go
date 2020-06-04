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
		token   string
		loguser models.User
	)

	c.Bind(&loguser)
	loguser.Password = util.HashString(loguser.Password)
	if services.ContainsUser(&loguser) {
		//Generates Token if valid User info.
		util.TokenGen(&token, &loguser)
		c.Header("Authorization", token)
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Authorization")

	} else {
		c.Header("Authorization", "Not-Authorized ")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Expose-Headers", "Authorization")
		util.AbortWithStatusAndMessage(c, "Controller: User does not exist", 401)
	}
}
