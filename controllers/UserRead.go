package controllers

import (
	"log"
	"strconv"

	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/TestingGorm/util"
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
		c.JSON(200, u)
	} else {
		util.AbortWithStatusAndMessage(c, "User does not exist", 403)
	}

}
