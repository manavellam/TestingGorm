package models

import "github.com/jinzhu/gorm"

//Editorial has info about different editorials
type Editorial struct {
	gorm.Model
	Name  string `gorm:"NOT NULL"`
	Books []*Book
}
