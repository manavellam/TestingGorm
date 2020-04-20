package models

import "github.com/jinzhu/gorm"

//Membership contains details about the rights of each user
type Membership struct {
	gorm.Model
	Name        string `gorm:"NOT NULL"`
	Description string `gorm:"NOT NULL"`
	UserID      []User `gorm:"NOT NULL"` //EACH MEMBERSHIP MAY CONTAIN ZERO TO MANY USERS
}
