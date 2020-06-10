package routes

import (
	"github.com/gin-gonic/gin"
)

//Routes handles routing for every endpoint
func Routes() *gin.Engine {
	router := gin.Default()
	userpaths(router)
	loginpaths(router)
	registerpaths(router)
	adminpaths(router)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "It works!",
		})
	})
	return router
}
