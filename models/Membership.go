package models

import "github.com/jinzhu/gorm"

//Membership contains details about the rights of each user
type Membership struct {
	gorm.Model         //        uint   `gorm:"primary_key; NOT NULL; unique "`
	Name        string `gorm:"NOT NULL"`
	Description string `gorm:"NOT NULL"`
}
