package models

import "github.com/jinzhu/gorm"

//User structurates user info to write and read from users table
type User struct {
	gorm.Model
	//EACH USER BELONGS TO ONE MEMBERSHIP
	MembershipID string `gorm:"NOT NULL"`
	Membership   Membership
	//EACH USER BELONGS TO ONE ACCES LEVEL
	AccessLevelID string `gorm:"NOT NULL"`
	AccessLevels  AccessLevels
	//-----------------------------------
	Name     string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}
