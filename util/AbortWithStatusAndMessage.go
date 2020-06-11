package util

import "github.com/gin-gonic/gin"

//AbortWithStatusAndMessage writes header and message before aborting request
func AbortWithStatusAndMessage(c *gin.Context, message string, status int) {
	c.Writer.WriteHeader(status)
	c.Writer.Write([]byte(message))
	c.Abort()
}