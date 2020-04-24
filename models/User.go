package models

import "github.com/jinzhu/gorm"

//User structurates user info to write and read from users table
type User struct {
	gorm.Model
	//EACH USER BELONGS TO ONE MEMBERSHIP
	MembershipID int `gorm:"NOT NULL"`
	Membership   Membership
	//EACH USER BELONGS TO ONE ACCES LEVEL
	AccessLevelsID int `gorm:"NOT NULL"`
	AccessLevels   AccessLevels
	//EACH USER HAS MANY CREDIT CARDS
	Creditcards []Creditcard
	//-----------------------------------
	Name     string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}
