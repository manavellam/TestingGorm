package models

import "github.com/jinzhu/gorm"

//Author has details about diff authors
type Author struct {
	gorm.Model
	Name   string `gorm:"NOT NULL"`
	BookID uint
	Books  []Book `gorm:"many2many:book_author"`
}
