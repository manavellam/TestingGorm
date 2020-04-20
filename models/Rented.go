package models

import "github.com/jinzhu/gorm"

//Rented details who has currently wich book, and when should the User give back each book
type Rented struct {
	gorm.Model	
}