package routes

import (
	"github/TestingGorm/controllers"

	"github.com/gin-gonic/gin"
)

func loginpaths(r *gin.Engine) {
	l := r.Group("/login")
	{
		l.POST("/auth", controllers.LoginUser)
	}
}
