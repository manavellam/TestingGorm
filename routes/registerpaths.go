package routes

import (
	"github.com/TestingGorm/controllers"
	"github.com/TestingGorm/middleware"
	"github.com/gin-gonic/gin"
)

func registerpaths(r *gin.Engine) {
	u := r.Group("/register")
	{
		u.POST("/", middleware.CheckUserData(), controllers.UserRegister)
	}
}
