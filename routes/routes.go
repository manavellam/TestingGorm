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
		c.Data(200, "text/html; charset=utf-8", []byte("<h1>Your go-app is working!</h1>"))
	})
	return router
}
