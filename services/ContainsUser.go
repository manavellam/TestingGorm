package services

import (
	"github.com/TestingGorm/models"
)

//ContainsUser check if an entry exists
func ContainsUser(u *models.User) bool {
	var result models.User
	if u.Password != "" {
		database.db.Where("name = ? AND password= ?", u.Name, u.Password).Set("gorm:auto_preload", true).Find(&result)
	} else if u.ID != 0 {
		database.db.Where("ID = ?", u.ID).Set("gorm:auto_preload", true).Find(&result)
	}
	if result.ID != 0 {
		*u = result
		return true
	}

	return false
}
