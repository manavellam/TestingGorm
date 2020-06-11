package models

import "github.com/jinzhu/gorm"

//Creditcard is the model for creditcards
type Creditcard struct {
	gorm.Model
	Number string
	UserID int
}
