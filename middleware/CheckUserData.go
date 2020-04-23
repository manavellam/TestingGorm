package middleware

import (
	"fmt"
	"strings"

	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/TestingGorm/util"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

//CheckUserData checks if the information sent  to register a new user is complete and not repited
func CheckUserData() gin.HandlerFunc {
	return func(c *gin.Context) {
		var us []models.User
		c.ShouldBindBodyWith(&us, binding.JSON)
		for i := range us {
			if len(strings.TrimSpace(us[i].Name)) != 0 && len(strings.TrimSpace(us[i].Password)) != 0 {

				us[i].Password = util.HashString(us[i].Password)
				if services.ContainsUser(&us[i]) {
					AbortWithStatusAndMessage(c, fmt.Sprintf("User with Name: %v already exists.\n", us[i].Name), 403)
				}
			} else {
				AbortWithStatusAndMessage(c, fmt.Sprintf("Missing information in user number: %v.\n", i), 403)
			}
		}

		if len(us) > 0 {
			//Stores read data back in context as env variable
			c.Set("UserList", us)
		} else {
			AbortWithStatusAndMessage(c, fmt.Sprintf("No valid user(s) sended.\n"), 403)
		}
	}
}
