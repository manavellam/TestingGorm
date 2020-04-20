package models

import "github.com/jinzhu/gorm"

//BookCategorie has different categories and the details for each one
type BookCategorie struct {
	gorm.Model
	Name        string `gorm:"NOT NULL"`
	Description string `gorm:"NOT NULL"`
	Books       []Book
}
