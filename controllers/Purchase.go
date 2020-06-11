package controllers

import (
	"log"

	"github.com/TestingGorm/models"
	"github.com/TestingGorm/services"
	"github.com/gin-gonic/gin"
)

//Purchase list the purchased items of a given user
func Purchase(c *gin.Context) {
	var (
		u           models.User
		Products    []models.Product
		ProductsIDs []int
	)
	c.Bind(&ProductsIDs)
	log.Print(ProductsIDs)

	services.GetProducts(ProductsIDs, &Products)
	u.Name = c.MustGet("username").(string)
	services.GetUser(&u)

	services.AddPurchase(&u, &Products)
}
