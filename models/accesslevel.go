package models

import "github.com/jinzhu/gorm"

//AccessLevels gives each user certain permissions
type AccessLevels struct {
	gorm.Model
	Role string
}
