package routes

import (
	"github.com/TestingGorm/controllers"
	"github.com/TestingGorm/middleware"

	"github.com/gin-gonic/gin"
)

func userpaths(r *gin.Engine) {
	u := r.Group("/user", middleware.Authorization())
	{
		u.GET("/readall", controllers.UserReadAll)
		u.GET("/read/:id", controllers.UserRead)
	}

}
