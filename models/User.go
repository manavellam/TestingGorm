package models

import "github.com/jinzhu/gorm"

//User structurates user info to write and read from users table
type User struct {
	gorm.Model
	Name     string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
	//Membership uint    `gorm:"NOT NULL"`                 //EACH USER HAS ONE MEMBERSHIP
	//Books      []*Book `gorm:"many2many:user_readbooks"` //EACH USER READ MANY BOOKS
	//Rented     []Book  //USER HAS MANY BOOKS RENTED
}
