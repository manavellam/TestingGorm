package models

import "github.com/jinzhu/gorm"

//Book stores data about the different books
type Book struct {
	gorm.Model
	Name        string `gorm:"NOT NULL"`
	Author      string `gorm:"NOT NULL"`
	CategorieID uint
	Collection  uint
	Rented      bool
	Authors     []*Author    `gorm:"many2many:book_author"`
	Readby      []*User      `gorm:"many2many:book_readby"`
	Editorials  []*Editorial `gorm:"many2many:book_editorial"`
}
