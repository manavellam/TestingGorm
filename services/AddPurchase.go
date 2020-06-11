package services

import "github.com/TestingGorm/models"

//AddPurchase appends one or several products to user
func AddPurchase(u *models.User, p *[]models.Product) {
	database.db.Model(&u).Association("Products").Append(p)
}
