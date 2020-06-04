package models

import "github.com/jinzhu/gorm"

//Product list products to be sold
type Product struct {
	gorm.Model
	Name  string
	Users []*User `gorm:"many2many:user_products;"`
}
