package models

import "github.com/jinzhu/gorm"

//User structurates user info to write and read from users table
type User struct {
	gorm.Model
	Name     string `gorm:"NOT NULL"`
	Password string `gorm:"NOT NULL"`
}
