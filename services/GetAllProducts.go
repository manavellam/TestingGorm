package services

import (
	"github.com/TestingGorm/models"
)

//GetAllProducts returns all products with associations
func GetAllProducts(p *[]models.Product) {
	database.db.Set("gorm:auto_preload", true).Find(p)
}
