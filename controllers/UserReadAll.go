package controllers

import (
	"github/TestingGorm/middleware"
	"github/TestingGorm/models"
	"github/TestingGorm/services"
	"log"

	"github.com/gin-gonic/gin"

	//Mysqldriver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//UserReadAll responds to a GET request to /user/read and returns all the user info from the database
func UserReadAll(c *gin.Context) {
	var user1 []models.User
	token := c.GetHeader("Authorization")

	if middleware.IsTknValid(&token) {
		services.Query("users", &user1)
		c.JSON(200, user1)
	} else {
		//TOKEN NO VALIDO. REDIRECCIONAR.
		log.Print("Session time expired")
		c.Status(401)
	}
}
