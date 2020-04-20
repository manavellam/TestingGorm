package routes

import (
	"github.com/TestingGorm/controllers"
	"github.com/TestingGorm/middleware"

	"github.com/gin-gonic/gin"
)

func loginpaths(r *gin.Engine) {
	l := r.Group("/login")
	{
		l.POST("/auth", middleware.CheckUserData(), controllers.LoginUser)
	}
}
