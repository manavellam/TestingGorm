package controllers

import (
	"log"
	"strconv"

	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/gin-gonic/gin"
)

//UserRead returns with the id specified in path user/read/:id
func UserRead(c *gin.Context) {
	var u models.User
	a, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		log.Print(err)
	}

	u.ID = uint(a)
	if services.ContainsUser(&u) {
		services.RelateUser(&u)
		c.JSON(200, u)
	} else {
		c.Writer.WriteHeader(403)
		c.Writer.Write([]byte("User does not exist"))
	}

}
