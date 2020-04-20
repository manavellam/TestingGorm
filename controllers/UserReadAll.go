package controllers

import (
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"

	"github.com/gin-gonic/gin"

	//Mysqldriver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//UserReadAll responds to a GET request to /user/read and returns all the user info from the database
func UserReadAll(c *gin.Context) {
	var usersList []models.User
	services.Query("users", &usersList)
	c.JSON(200, usersList)
}
