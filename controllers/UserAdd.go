package controllers

import (
	"log"

	"github.com/TestingGorm/middleware"
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"

	"github.com/gin-gonic/gin"
)

//UserAdd responds to a POST request at /user/add, and adds an user to the database given a balid token.
func UserAdd(c *gin.Context) {
	var (
		loaduser []models.User
	)

	token := c.GetHeader("Authorization")

	if middleware.IsTknValid(&token) {
		c.Bind(&loaduser)
		//Podria implementarse goroutine
		for _, u := range loaduser {
			services.Insert("users", &u)
		}
	} else {
		//TOKEN NO VALIDO. REDIRECCIONAR.
		log.Print("Session time expired")
		c.Status(401)
	}
}
