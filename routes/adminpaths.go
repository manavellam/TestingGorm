package routes

import (
	"github.com/TestingGorm/controllers"
	"github.com/TestingGorm/middleware"
	"github.com/gin-gonic/gin"
)

func adminpaths(r *gin.Engine) {
	u := r.Group("/admin", middleware.Authorization())
	{
		u.GET("/allProducts", controllers.ProductsAll)
	}

}
