package services

import (
	"log"

	"github.com/TestingGorm/models"
)

//ContainsUser check if an entry exists
func ContainsUser(u *models.User) bool {
	var result models.User
	database.db.Where("name = ? AND password= ?", u.Name, u.Password).Find(&result)
	log.Print("USER ID: ", result.ID)
	if result.ID != 0 {
		return true
	}
	return false
}
