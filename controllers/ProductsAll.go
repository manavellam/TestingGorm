package controllers

import (
	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/gin-gonic/gin"
)

//ProductsAll returns all products
func ProductsAll(c *gin.Context) {
	var products []models.Product
	services.GetAllProducts(&products)
	c.JSON(200, products)

}
