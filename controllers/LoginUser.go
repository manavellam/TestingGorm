package controllers

import (
	"log"

	"github.com/TestingGorm/middleware"
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
	log.Print("INFO:", loguser)
	if services.ContainsUser(&loguser) {

		//Generates Token if valid User info.
		util.TokenGen(&token, &loguser)
		log.Print()
		c.Header("Authorization", token)
	} else {
		middleware.AbortWithStatusAndMessage(c, "Controller: User does not exist", 401)
	}
}
