package routes

import (
	"github.com/TestingGorm/controllers"

	"github.com/gin-gonic/gin"
)

func userpaths(r *gin.Engine) {
	u := r.Group("/user")
	{
		u.GET("/read", controllers.UserReadAll)
		u.POST("/add", controllers.UserAdd)
	}

}
