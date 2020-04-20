package controllers

import (
	"log"

	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"

	"github.com/gin-gonic/gin"
)

//UserRegister responds to a POST request at /register, and adds an user to the database if not exists.
func UserRegister(c *gin.Context) {
	loaduser := c.MustGet("UserList").([]models.User)
	//Podria implementarse goroutine
	for _, u := range loaduser {
		log.Print(u)
		services.Insert("users", &u)
	}
	c.Writer.WriteHeader(200)
	c.Writer.WriteString("User(s) succesfuly registered into DB.")
}
