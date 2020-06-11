package services

import "github.com/TestingGorm/models"

//GetProducts gets the products to be purchased
func GetProducts(list []int, products *[]models.Product) {
	database.db.Where("id IN (?)", list).Find(products)
}
