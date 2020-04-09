package services

import (
	"github/TestingGorm/models"
)

//ContainsUser check if an entry exists
func ContainsUser(u *models.User) bool {
	var result models.User
	database.db.Where("name = ? AND password= ?", u.Name, u.Password).Find(&result)

	if result.ID != 0 {
		return true
	}

	return false
}
