package controllers

import (
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/TestingGorm/util"
	"github.com/gin-gonic/gin"
)

//Cont is used to store creditcard number from request's body
type Cont struct {
	Number string
}

//AddCreditcard adds creditcard into database.
func AddCreditcard(c *gin.Context) {
	var Number Cont
	c.Bind(&Number)
	var u models.User

	u.Name = c.MustGet("username").(string)
	services.GetUser(&u)

	if !services.ContainsCreditcard(&u, Number.Number) {
		services.AddCreditCard(&u, Number.Number)
		c.Writer.WriteHeader(200)
		c.Writer.WriteString("Creditcard added to database")
	} else {
		util.AbortWithStatusAndMessage(c, "The Credit card is already in the database", 400)
	}
}
